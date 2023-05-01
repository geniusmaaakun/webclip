package actions_test

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"webclip/src/actions"
	"webclip/src/server/models"
	"webclip/src/server/models/rdb"
	"webclip/src/server/usecases"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

//exampletestでもいい？
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

/*
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


		// コマンドライン引数なしのテスト
		// err := app.Run([]string{"", "greet"})
		// assert.NoError(t, err)
		// assert.Equal(t, "Hello, world!\n", buf.String())
		// t.Log(errbuf.String())
		// t.Log(buf.String())


}
*/

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

	//ここら辺はmockにする
	//dbの作成
	db, err := models.NewDB("webclip.sql")
	if err != nil {
		t.Fatalf("database error: %v\n", err)
	}
	txRepo := rdb.NewTransactionManager(db)
	markdownRepo := rdb.NewMarkdownRepo()
	markdownUsecase := usecases.NewMarkdownInteractor(txRepo, markdownRepo)

	type args struct {
		title  string
		path   string
		srcUrl string
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	//testcaseの作成
	tests := []struct {
		id        int
		name      string
		args      args
		wantTitle string //インスタンスを返す様にcreateを変更する？
	}{}

	for i := 0; i < 100; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("test%d", i))
		s := fmt.Sprintf("%x", h.Sum(nil))
		//t.Log(s)
		tests = append(tests, struct {
			id        int
			name      string
			args      args
			wantTitle string
		}{i, fmt.Sprintf("normal %d", i), args{fmt.Sprintf("test%s", string(s)), fmt.Sprintf("testdata/%s/README.md", string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s))}, fmt.Sprintf("test%s", string(s))})
		err := markdownUsecase.Create(fmt.Sprintf("test%s", string(s)), fmt.Sprintf("testdata/%s/README.md", string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s)))
		if err != nil {
			t.Errorf("MarkdownInteractor.Delete() error = got %v, s value = %s\n", err, string(s))
		}

	}

	app := &cli.Context{
		App: actions.NewWebClip("webclip.sql"),
	}

	args2 := []string{"", "clean"}

	//clean後に確認
	got := extractStdout(t, app, args2)

	t.Log(got)

	mds, err := markdownUsecase.FindAll()
	if err != nil {
		t.Errorf("MarkdownInteractor.FindAll() error = %v\n", err)
	}

	if (len(mds)) != 0 {
		t.Errorf("MarkdownInteractor.FindAll() error = %v\n", err)
	}

}

func TestSearch(t *testing.T) {
	//bodyとtitleで検索
	//先にデータベースを作成しておく　リスト

	//検索し、結果を表示する
}

func TestZip(t *testing.T) {

}

func TestServer(t *testing.T) {

}
