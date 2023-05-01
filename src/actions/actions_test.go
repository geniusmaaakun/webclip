package actions_test

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
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
	go func() {
		err := http.ListenAndServe(":8090", http.HandlerFunc(ImageHandler))
		if err != nil {
			fmt.Println(err)
		}
	}()

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	c := &cli.Context{
		App: &cli.App{
			Name:  "HTML to Markdown converter",
			Usage: "Convert HTML files to Markdown with optional image downloading",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "url",
					Aliases: []string{"u"},
					Usage:   "Target URL",
					//Required: true,
				},
				&cli.StringFlag{
					Name:    "outdir",
					Aliases: []string{"o"},
					Usage:   "Output directory",
					//Required: true,
				},
				&cli.BoolFlag{
					Name:    "download",
					Aliases: []string{"d"},
					Usage:   "Download images",
				},
				&cli.BoolFlag{
					Name:    "save",
					Aliases: []string{"sv"},
					Usage:   "Save to DB",
				},
			},
			//WebClip
			Action: actions.Download,
		},
	}

	//すべてのファイルを削除する　危険

	c.App.Commands = []*cli.Command{
		//sub command : webclip server
		{
			Name:   "server",
			Usage:  "Start Web Server",
			Action: actions.Server,
		},
		//sub command : webclip clean
		{
			Name:   "clean",
			Usage:  "clean database if file is not exist",
			Action: actions.Clean,
		},
		//sub command : webclip search
		{
			Name:  "search",
			Usage: "search markdown file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "title",
					Aliases: []string{"t"},
					Usage:   "search title",
				},
				&cli.StringFlag{
					Name:    "body",
					Aliases: []string{"b"},
					Usage:   "search body",
				},
			},
			Action: actions.Search,
		},
		//zip化する
		//sub command : webclip zip
		{
			Name:  "zip",
			Usage: "zip markdown file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "title",
					Aliases: []string{"t"},
					Usage:   "search title",
				},
				&cli.StringFlag{
					Name:    "body",
					Aliases: []string{"b"},
					Usage:   "search body",
				},
			},
			Action: actions.Zip,
		},
	}

	/*
		defer func() {
			//すべてのファイルを削除する　危険
			os.RemoveAll("./testdata/")
		}()
	*/

	//テスト用のコマンドライン引数
	//args := []string{"", "-u", "http://localhost:8090", "-o", "./testdata", "--download", "--save"}
	args := []string{"", "-u", "http://localhost:8090", "-o", "./testdata", "--download"}

	//テスト実行
	err := c.App.Run(args)
	assert.NoError(t, err)

	//出力が一致しているか

	//ファイルが一致しているか？
}
