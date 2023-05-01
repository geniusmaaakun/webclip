package wcconverter_test

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"webclip/src/wcconverter"
	"webclip/src/wcdownloader"
)

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./testdata/test.html"))
	t.Execute(w, nil)
}

// htmlを変換

func TestConvertFromHTML(t *testing.T) {
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

	//convert
	converter := wcconverter.NewConverter("testdata", "README.md", nil)
	markdown, err := converter.Convert(doc.Selection)
	if err != nil {
		t.Fatalf("Markdown Conversion Error: %s", err.Error())
	}
	markdown = converter.AddSrcUrlToMarkdown("http://localhost:8090", markdown)
	err = converter.SaveToFile(markdown)
	if err != nil {
		t.Fatalf("Markdown Conversion Error: %s", err.Error())
	}

	got, err := os.Open("testdata/README.md")
	if err != nil {
		t.Fatalf("Markdown Conversion Error: %s", err.Error())
	}
	defer got.Close()
	want, err := os.Open("testdata/README_ok.md")
	if err != nil {
		t.Fatalf("Markdown Conversion Error: %s", err.Error())
	}
	defer want.Close()

	gotBody, err := io.ReadAll(got)
	if err != nil {
		t.Fatalf("Markdown Conversion Error: %s", err.Error())
	}
	wantBody, err := io.ReadAll(want)
	if err != nil {
		t.Fatalf("Markdown Conversion Error: %s", err.Error())
	}

	if bytes.Compare(gotBody, wantBody) != 0 {
		t.Fatalf("Markdown Conversion Error: got: %s, want %s\n", string(gotBody), string(wantBody))
	}

}
