package wcdownloader

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
)

type Downloader struct {
	ImageDownloader *ImageDownloader
	HtmlDownloader  *HtmlDownloader
}

func NewDownloader(srcUrl string, outputDir string, imageDownloadFlag bool) *Downloader {
	imageDownloader := &ImageDownloader{OutputDir: outputDir, ImageDownloadFlag: imageDownloadFlag, Client: &http.Client{
		Timeout: 10 * time.Second,
	}}
	htmlDownloader := &HtmlDownloader{url: srcUrl}

	return &Downloader{imageDownloader, htmlDownloader}
}

type HtmlDownloader struct {
	url string
}

func (d *HtmlDownloader) CreateDocument() (*goquery.Document, error) {

	//大元にwcdownloaderを作成 子にwcdownloaderとhtmlwcdownloaderを作成
	//2 HTMLを取得 html wcdownloader createdocument
	resp, err := http.Get(d.url)
	if err != nil {
		//log.Fatalf("http error: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	utf8Reader, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		//log.Fatalf("encoding error: %s", err.Error())
		return nil, err
	}

	//utf-8以外も文字化けしない様にする
	doc, err := goquery.NewDocumentFromReader(utf8Reader)
	//doc, err := goquery.NewDocument(url)
	if err != nil {
		//log.Fatalf("goquery error: %s", err.Error())
		return nil, err
	}
	return doc, nil
}

type ImageDownloader struct {
	Client            *http.Client
	OutputDir         string
	ImageDownloadFlag bool
	imgCounter        int
}

func (d *ImageDownloader) ReplaceImageSrcToImageFile(doc *goquery.Document) (*goquery.Document, error) {
	if d.ImageDownloadFlag {
		doc.Find("img").Each(func(i int, s *goquery.Selection) {
			imgURL, exists := s.Attr("src")

			if exists && imgURL != "" {
				filename, err := d.Download(imgURL)
				if err != nil {
					fmt.Println("failed to download image: ", err)
					imgURL, exists = s.Attr("data-src")
					if exists && imgURL != "" {
						filename, err = d.Download(imgURL)
						if err != nil {
							fmt.Println("failed to download image: ", err)
						} else {
							//画像ファイル名が日本語であっても、エンコードされたファイル名をMarkdownファイル内の画像リンクに使用するため、正しく表示されるようになります。ただし、この変更ではダウンロードされた画像ファイル自体のファイル名はエンコードされません。ダウンロードされた画像ファイルのファイル名もエンコードする場合は、wcdownloaderのDownload関数内でos.Createを呼び出す際に、エンコードされたファイル名を使用するように変更してください。
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
	return doc, nil
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
