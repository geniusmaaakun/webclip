package actions_test

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"webclip/src/actions"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

// Stdoutに書き込まれた文字列を抽出する関数
// (Stderrも同じ要領で出力先を変更できます)
func extractStdout(t *testing.T, c *cli.Context, args []string) string {
	t.Helper()

	// 既存のStdoutを退避する
	orgStdout := os.Stdout
	defer func() {
		// 出力先を元に戻す
		os.Stdout = orgStdout
	}()
	// パイプの作成(r: Reader, w: Writer)
	r, w, _ := os.Pipe()
	// Stdoutの出力先をパイプのwriterに変更する
	os.Stdout = w
	// テスト対象の関数を実行する
	err := c.App.Run(args)
	if err != nil {
		t.Fatalf("failed to read buf: %v", err)
	}
	// Writerをクローズする
	// Writerオブジェクトはクローズするまで処理をブロックするので注意
	w.Close()
	// Bufferに書き込こまれた内容を読み出す
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("failed to read buf: %v", err)
	}
	// 文字列を取得する
	return strings.TrimRight(buf.String(), "\n")
}

func greet(c *cli.Context) error {
	name := "world"

	g := c.String("greet")
	fmt.Println(g)
	fmt.Printf("Hello, %s!\n", name)
	return nil
}

func TestGreet(t *testing.T) {
	//buf := new(bytes.Buffer)
	//errbuf := new(bytes.Buffer)

	c := &cli.Context{
		App: &cli.App{
			Name:  "sample",
			Usage: "A simple CLI application",
			// Writer:    buf,
			// ErrWriter: errbuf,
			Commands: []*cli.Command{
				{
					Name:    "greet",
					Aliases: []string{"g"},
					Usage:   "Prints a greeting message",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "greet",
							Value: "Good morning",
							Usage: "Set greeting message",
						},
					},

					Action: greet,
				},
			},
		},
	}

	got := extractStdout(t, c, []string{"", "greet", "--greet", "test"})

	assert.Equal(t, "test\nHello, world!", got)

	/*
		// コマンドライン引数なしのテスト
		err := app.Run([]string{"", "greet"})
		assert.NoError(t, err)
		assert.Equal(t, "Hello, world!\n", buf.String())
		t.Log(errbuf.String())
		t.Log(buf.String())

	*/

	// バッファをリセット
	/*
		buf.Reset()
		errbuf.Reset()

		// 名前引数付きのテスト
		err = app.Run([]string{"", "greet", "Alice"})
		assert.NoError(t, err)
		assert.Equal(t, "Hello, Alice!\n", buf.String())
	*/
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./testdata/test.html"))
	t.Execute(w, nil)
}

func TestDownload(t *testing.T) {
	//テストサーバー
	sv := httptest.NewServer(http.HandlerFunc(ImageHandler))
	defer sv.Close()

	c := &cli.Context{
		App: actions.NewWebClip("webclip.sql"),
	}

	//テスト用のコマンドライン引数
	//args := []string{"", "-u", "http://localhost:8090", "-o", "./testdata", "--download", "--save"}
	args := []string{"", "-u", sv.URL, "-o", "./testdata", "--download"}

	//テスト実行
	got := extractStdout(t, c, args)

	//出力が一致しているか
	assert.Equal(t, fmt.Sprintf("Target: %s", sv.URL), got)

	//ファイルが一致しているか？
	files, err := filepath.Glob("./testdata/*")
	if err != nil {
		t.Fatal(err)
	}

	expectedFiles := []string{"testdata/README.md", "testdata/test.html", "testdata/testdata-1.jpe", "testdata/testdata-10.png", "testdata/testdata-11.png", "testdata/testdata-12.png", "testdata/testdata-13.png", "testdata/testdata-14.png", "testdata/testdata-15.png", "testdata/testdata-16.png", "testdata/testdata-17.png", "testdata/testdata-18.png", "testdata/testdata-19.png", "testdata/testdata-2.png", "testdata/testdata-20.png", "testdata/testdata-21.png", "testdata/testdata-22.png", "testdata/testdata-23.png", "testdata/testdata-24.png", "testdata/testdata-25.png", "testdata/testdata-26.png", "testdata/testdata-27.png", "testdata/testdata-28.png", "testdata/testdata-29.png", "testdata/testdata-3.svg", "testdata/testdata-30.jpe", "testdata/testdata-31.png", "testdata/testdata-32.jpe", "testdata/testdata-33.jpe", "testdata/testdata-34.svg", "testdata/testdata-4.png", "testdata/testdata-5.png", "testdata/testdata-6.png", "testdata/testdata-7.png", "testdata/testdata-8.png", "testdata/testdata-9.png"}

	for i := range files {
		if files[i] != expectedFiles[i] {
			t.Errorf("file not match: %s", files[i])
		}
	}
}

func TestClean(t *testing.T) {

	/*
		app := &cli.Context{
			App: actions.NewWebClip("webclip.sql"),
		}

		args := []string{"", "clean"}

		//testdataDBの作成

		//clean後に確認
	*/

}

func TestSearch(t *testing.T) {

}

func TestZip(t *testing.T) {

}

func TestServer(t *testing.T) {

}
