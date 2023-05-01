package actions_test

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
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
	t := template.Must(template.ParseFiles("./testdata/test/test.html"))
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
	args := []string{"", "-u", sv.URL, "-o", "./testdata/test", "--download"}

	//テスト実行
	got := extractStdout(t, c, args)

	//出力が一致しているか
	assert.Equal(t, fmt.Sprintf("Target: %s", sv.URL), got)

	//ファイルが一致しているか？
	files, err := filepath.Glob("./testdata/test/*")
	if err != nil {
		t.Fatal(err)
	}

	expectedFiles := []string{"testdata/test/README.md", "testdata/test/test.html", "testdata/test/testdata-1.jpe", "testdata/test/testdata-10.png", "testdata/test/testdata-11.png", "testdata/test/testdata-12.png", "testdata/test/testdata-13.png", "testdata/test/testdata-14.png", "testdata/test/testdata-15.png", "testdata/test/testdata-16.png", "testdata/test/testdata-17.png", "testdata/test/testdata-18.png", "testdata/test/testdata-19.png", "testdata/test/testdata-2.png", "testdata/test/testdata-20.png", "testdata/test/testdata-21.png", "testdata/test/testdata-22.png", "testdata/test/testdata-23.png", "testdata/test/testdata-24.png", "testdata/test/testdata-25.png", "testdata/test/testdata-26.png", "testdata/test/testdata-27.png", "testdata/test/testdata-28.png", "testdata/test/testdata-29.png", "testdata/test/testdata-3.svg", "testdata/test/testdata-30.jpe", "testdata/test/testdata-31.png", "testdata/test/testdata-32.jpe", "testdata/test/testdata-33.jpe", "testdata/test/testdata-34.svg", "testdata/test/testdata-4.png", "testdata/test/testdata-5.png", "testdata/test/testdata-6.png", "testdata/test/testdata-7.png", "testdata/test/testdata-8.png", "testdata/test/testdata-9.png"}
	expectedFilesMap := make(map[string]bool)

	for _, f := range expectedFiles {
		expectedFilesMap[f] = true
	}

	for _, v := range files {
		expectedFilesMap[v] = false
	}

	for k, v := range expectedFilesMap {
		if v {
			t.Errorf("file not found: %s", k)
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
		//存在するパス
		//err := markdownUsecase.Create(fmt.Sprintf("test%s", string(s)), fmt.Sprintf("testdata/%s/README.md", string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s)))
		//存在しないパス
		err := markdownUsecase.Create(fmt.Sprintf("test%s", string(s)), fmt.Sprintf("testdata/%sfalse/README.md", string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s)))
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
		t.Errorf("MarkdownInteractor.FindAll() error = %v\n", len(mds))
	}

}

func TestSearch(t *testing.T) {
	//bodyとtitleで検索
	//先にデータベースを作成しておく　リスト

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

	//検索し、結果を表示する
	args2 := []string{"", "search", "-b", "test"}
	got := extractStdout(t, app, args2)
	t.Log(got)
	//出力が一致するかを確認する
}

func TestZip(t *testing.T) {
	//zip化して、展開 usecasesとほぼ同じ

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
		os.Remove("webclip.zip")
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

	//args2 := []string{"", "zip", "-t", "test"}
	args2 := []string{"", "zip", "-b", "test"}

	//got := extractStdout(t, app, args2)
	//t.Log(got)

	_ = extractStdout(t, app, args2)

	zr, err := zip.OpenReader("webclip.zip")
	if err != nil {
		t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
	}

	defer zr.Close()

	for _, zfile := range zr.File {
		reader, err := zfile.Open()
		if err != nil {
			t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
		}

		got, err := ioutil.ReadAll(reader)
		if err != nil {
			t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
		}

		//bodyの中身がtestdataの中身と一致するかどうか
		want, err := ioutil.ReadFile(filepath.Join("testdata", filepath.Dir(zfile.Name), filepath.Base(zfile.Name)))

		if err != nil {
			t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
		}

		if !bytes.Equal(got, want) {
			t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
		}

		defer reader.Close()
	}

}

func TestServer(t *testing.T) {
	//server起動してhttpアクセス reactでやる

}
