package main

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli/v2"
)

//このコードでは、`urfave/cli/v2`ライブラリを使用してコマンドラインオプションを定義し、`Action`関数で処理を実行しています。このライブラリを使用することで、コマンドラインオプションの定義が簡潔になり、エラーチェックやヘルプメッセージの生成などが自動化されます。
//上記のコードでは、`HTMLToMarkdownConverter`と`ImageDownloader`という構造体を使用して、処理を行っています。これにより、コードがより構造化され、プロのエンジニアが書いたようなスタイルになっています。

type HTMLToMarkdownConverter struct {
	Options *md.Options
}

func (c *HTMLToMarkdownConverter) Convert(selection *goquery.Selection) (string, error) {
	converter := md.NewConverter("", true, c.Options)
	markdown := converter.Convert(selection)
	return markdown, nil
}

type ImageDownloader struct {
	Client     *http.Client
	OutputDir  string
	imgCounter int
}

func (d *ImageDownloader) Download(imgURL string) (string, error) {
	/*
		resp, err := d.Client.Get(imgURL)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("failed to download image: status code %d", resp.StatusCode)
		}

		filename := imgFilename(imgURL)
		file, err := os.Create(filepath.Join(d.OutputDir, filename))
		if err != nil {
			return "", err
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return "", err
		}

		return filename, nil
	*/

	//画像ファイルの拡張子をより安全に取得
	resp, err := d.Client.Get(imgURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download image: status code %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")

	//mimeパッケージを使用してContent-Typeから拡張子を取得する方法を以下に示します。
	//mime.ExtensionsByType関数は、Content-Typeに対応する拡張子のスライスを返します。通常は最初の拡張子（ext[0]）を使用しますが、必要に応じて他の拡張子を選択することもできます。
	// これらの方法を使用することで、より正確な画像ファイルの拡張子を取得することができます。ただし、いずれの方法でもサーバーが正しいContent-Typeを返していることが前提となります。サーバーが間違ったContent-Typeを返す場合は、拡張子の判定が正しく行われないことがあります。
	// Determine the file extension from the content type
	ext, err := mime.ExtensionsByType(contentType)
	if err != nil {
		return "", fmt.Errorf("unsupported content type: %s", contentType)
	}

	//new imageFilename
	//filename := imgFilename(imgURL, ext[0])
	//old imageFIlename
	//filename := imgFilename(imgURL) + ext[0]

	d.imgCounter++
	filename := fmt.Sprintf("%s-%d%s", filepath.Base(d.OutputDir), d.imgCounter, ext[0])

	//画像ファイル名を [ディレクトリ名]-[番号].png の形式に変更
	file, err := os.Create(filepath.Join(d.OutputDir, filename))

	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func main() {
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
		Action: func(c *cli.Context) error {
			url := c.String("url")
			outdir := c.String("outdir")
			imageDownloadFlag := c.Bool("download")

			err := os.MkdirAll(outdir, 0755)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Target: %s\n", url)

			doc, err := goquery.NewDocument(url)
			if err != nil {
				log.Fatalf("goquery error: %s", err.Error())
			}

			imageDownloader := &ImageDownloader{
				OutputDir: outdir,
				Client: &http.Client{
					Timeout: 10 * time.Second,
				},
			}

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

			converter := &HTMLToMarkdownConverter{}
			markdown, err := converter.Convert(doc.Selection)
			if err != nil {
				log.Fatalf("Markdown Conversion Error: %s", err.Error())
			}

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

	err := app.Run(os.Args)
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
