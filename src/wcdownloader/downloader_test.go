package wcdownloader_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"text/template"
	"webclip/src/wcdownloader"
)

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello from handler!\n</h1>")
}

func TestDownloadImageFalse(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(SimpleHandler))

	//w := httptest.NewRequest("GET", "http://example.com", nil)
	downloader := wcdownloader.NewDownloader(sv.URL, "test", false)
	doc, err := downloader.HtmlDownloader.CreateDocument()
	if err != nil {
		t.Fatal(err)
	}
	doc, err = downloader.ImageDownloader.ReplaceImageSrcToImageFile(doc)
	if err != nil {
		t.Fatal(err)
	}

	if strings.Compare("Hello from handler!\n", doc.Selection.Find("h1").Text()) != 0 {
		t.Fatalf("invalid html: %s", doc.Selection.Find("h1").Text())
	}
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./testdata/test.html"))
	t.Execute(w, nil)
}

//imageタグ
func TestDownloadImageTrue(t *testing.T) {
	sv := httptest.NewServer(http.HandlerFunc(ImageHandler))

	//w := httptest.NewRequest("GET", "http://example.com", nil)

	defer t.Cleanup(func() {
		os.RemoveAll(t.TempDir())
	})
	downloader := wcdownloader.NewDownloader(sv.URL, t.TempDir(), true)
	doc, err := downloader.HtmlDownloader.CreateDocument()
	if err != nil {
		t.Fatal(err)
	}
	doc, err = downloader.ImageDownloader.ReplaceImageSrcToImageFile(doc)
	if err != nil {
		t.Fatal(err)
	}

	//testで画像がダウンロードできない？

	//fmt.Println(dirwalk(t.TempDir()))

	///body, err := io.ReadAll(res.Body)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//defer res.Body.Close()
	//fmt.Println(string(body))

	//t.Log(doc.Selection.Find("img").Attr("src"))

	/*
		if doc.Selection.Find("img").Length() != 1 {
			t.Fatal("invalid html")
		}
	*/
}

//mdファイル

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}
