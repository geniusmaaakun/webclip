package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
	"webclip/src/downloader"
	"webclip/src/server/models"
	"webclip/src/wcconverter"

	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/html/charset"
)

//このコードでは、`urfave/cli/v2`ライブラリを使用してコマンドラインオプションを定義し、`Action`関数で処理を実行しています。このライブラリを使用することで、コマンドラインオプションの定義が簡潔になり、エラーチェックやヘルプメッセージの生成などが自動化されます。
//上記のコードでは、`HTMLToMarkdownConverter`と`ImageDownloader`という構造体を使用して、処理を行っています。これにより、コードがより構造化され、プロのエンジニアが書いたようなスタイルになっています。

type App struct {
}

func main() {
	_, err := models.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	app := &cli.App{
		Name:  "HTML to Markdown converter",
		Usage: "Convert HTML files to Markdown with optional image downloading",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Usage:    "Target URL",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "outdir",
				Aliases:  []string{"o"},
				Usage:    "Output directory",
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "download",
				Aliases: []string{"d"},
				Usage:   "Download images",
			},
		},
		//WebClip
		Action: func(c *cli.Context) error {
			//1 setup main
			//コマンドラインオプションの値を取得
			url := c.String("url")
			outdir := c.String("outdir")
			imageDownloadFlag := c.Bool("download")

			err := os.MkdirAll(outdir, 0755)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Target: %s\n", url)

			//大元にdownloaderを作成 子にimageDownloaderとhtmldownloaderを作成
			//2 HTMLを取得 html downloader
			resp, err := http.Get(url)
			if err != nil {
				log.Fatalf("http error: %s", err.Error())
			}
			defer resp.Body.Close()

			utf8Reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
			if err != nil {
				log.Fatalf("encoding error: %s", err.Error())
			}

			//utf-8以外も文字化けしない様にする
			doc, err := goquery.NewDocumentFromReader(utf8Reader)
			//doc, err := goquery.NewDocument(url)
			if err != nil {
				log.Fatalf("goquery error: %s", err.Error())
			}

			//option?
			imageDownloader := &downloader.ImageDownloader{
				OutputDir: outdir,
				Client: &http.Client{
					Timeout: 10 * time.Second,
				},
			}

			//3 HTMLをMarkdownに変換
			if imageDownloadFlag {
				doc.Find("img").Each(func(i int, s *goquery.Selection) {
					imgURL, exists := s.Attr("src")

					if exists && imgURL != "" {
						filename, err := imageDownloader.Download(imgURL)
						if err != nil {
							fmt.Println("failed to download image: ", err)
							imgURL, exists = s.Attr("data-src")
							if exists && imgURL != "" {
								filename, err = imageDownloader.Download(imgURL)
								if err != nil {
									fmt.Println("failed to download image: ", err)
								} else {
									//画像ファイル名が日本語であっても、エンコードされたファイル名をMarkdownファイル内の画像リンクに使用するため、正しく表示されるようになります。ただし、この変更ではダウンロードされた画像ファイル自体のファイル名はエンコードされません。ダウンロードされた画像ファイルのファイル名もエンコードする場合は、ImageDownloaderのDownload関数内でos.Createを呼び出す際に、エンコードされたファイル名を使用するように変更してください。
									encodedFilename := encodeFilename(filename)
									s.SetAttr("src", encodedFilename)
								}
							}
						} else {
							encodedFilename := encodeFilename(filename)
							s.SetAttr("src", encodedFilename)
						}
					}
				})
			}

			converter := &wcconverter.HTMLToMarkdownConverter{}
			markdown, err := converter.Convert(doc.Selection)
			if err != nil {
				log.Fatalf("Markdown Conversion Error: %s", err.Error())
			}

			//4 Markdownをファイルに保存
			file, err := os.Create(filepath.Join(outdir, "README.md"))
			if err != nil {
				log.Fatalf("File Create Error: %s\n", err.Error())
			}
			defer file.Close()
			_, err = file.WriteString(markdown)
			if err != nil {
				log.Fatalf("File Write Error: %s\n", err.Error())
			}

			return nil
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

/*
func imgFilename(imgURL string) string {
	imgURL = strings.Split(imgURL, "?")[0]
	arr := strings.Split(imgURL, "/")
	fmt.Println(arr[len(arr)-1])
	return arr[len(arr)-1]
}
*/

func imgFilename(imgURL, ext string) string {
	imgURL = strings.Split(imgURL, "?")[0]
	arr := strings.Split(imgURL, "/")
	baseName := arr[len(arr)-1]
	// 拡張子がURLに含まれている場合は、それを削除します
	baseName = strings.TrimSuffix(baseName, filepath.Ext(baseName))
	return baseName + ext
}

//画像ファイル名をエンコードする関数
//画像ファイル名に日本語が含まれている場合に、URLエンコードされたファイル名に変換します。
// 変換されたファイル名を、Markdownファイル内の画像リンクに使用します。
func encodeFilename(filename string) string {
	encodedName := url.PathEscape(filename)
	return encodedName
}
