package usecases_test

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
	"webclip/src/server/models"
	"webclip/src/server/models/rdb"
	"webclip/src/server/usecases"
)

//バリデーションとか。
//DB処理はmockにする 後ほど
func TestCreate(t *testing.T) {
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

	//testcaseの作成
	tests := []struct {
		name string
		args args
		want bool //インスタンスを返す様にcreateを変更する？
	}{
		{"normal 1", args{"test1", "/test1/test1", "http://test1/test1"}, true},
		{"fail SrcURL = isNotURL", args{"test2", "/test2/test2", "test1/test1"}, false},
		{"fail title = isEmpty", args{"", "/test2/test2", "test1/test1"}, false},
		{"fail path = isEmpty", args{"test3", "", "test1/test1"}, false},
		{"fail srcUrl = isEmpty", args{"test4", "test4", ""}, false},
		{"fail allField = isEmpty", args{"", "", ""}, false},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	//usecaseの実行
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := markdownUsecase.Create(tt.args.title, tt.args.path, tt.args.srcUrl)
			if tt.want {
				if err != nil {
					t.Errorf("MarkdownInteractor.Create() error = got %v\n", err)
				}
			} else {
				if err == nil {
					t.Errorf("MarkdownInteractor.Create() want error = got %v\n", err)
				}
			}
		})
	}
}

//unused
//タイトルで削除するのは危険なので、パスで削除するように変更する
//作成して、削除、確認
func TestDelete(t *testing.T) {
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

	//testcaseの作成
	tests := []struct {
		name string
		args args
		want bool //インスタンスを返す様にcreateを変更する？
	}{
		{"normal 1", args{"test1", "/test1/test1", "http://test1/test1"}, true},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	//usecaseの実行
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := markdownUsecase.Create(tt.args.title, tt.args.path, tt.args.srcUrl)
			if err != nil {
				t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
			}
			err = markdownUsecase.Delete(tt.args.title)
			if err != nil {
				t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
			}
			got, err := markdownUsecase.FindByTitle(tt.args.title)
			if err != nil {
				t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
			}
			if len(got) != 0 {
				t.Errorf("MarkdownInteractor.Delete() error = got %v\n", got)
			}
		})
	}
}

//作成して、削除、確認
func TestDeleteByPath(t *testing.T) {
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

	//testcaseの作成
	tests := []struct {
		name string
		args args
		want bool //インスタンスを返す様にcreateを変更する？
	}{
		{"normal 1", args{"test1", "/test1/test1", "http://test1/test1"}, true},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	//usecaseの実行
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := markdownUsecase.Create(tt.args.title, tt.args.path, tt.args.srcUrl)
			if err != nil {
				t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
			}
			err = markdownUsecase.DeleteByPath(tt.args.path)
			if err != nil {
				t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
			}
			got, err := markdownUsecase.FindByPath(tt.args.path)
			if err != nil {
				t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
			}
			if len(got) != 0 {
				t.Errorf("MarkdownInteractor.Delete() error = got %v\n", got)
			}
		})
	}
}

//先に追加しておき、その後に取得して一致するかどうか
func TestFindAll(t *testing.T) {
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

	//testcaseの作成 ランダムな文字列　これをいろんなテストケースに使う
	h := md5.New()
	for i := 0; i < 100; i++ {
		io.WriteString(h, fmt.Sprintf("test%d", i))
		s := fmt.Sprintf("%x", h.Sum(nil))
		t.Log(s)
		err := markdownUsecase.Create(fmt.Sprintf("test%s", string(s)), fmt.Sprintf("/test%s/test%s", string(s), string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s)))
		if err != nil {
			t.Errorf("MarkdownInteractor.Delete() error = got %v, s value = %s\n", err, string(s))
		}
	}
	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	//usecaseの実行
	mds, err := markdownUsecase.FindAll()
	if err != nil {
		t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
	}
	if len(mds) != 100 {
		t.Errorf("MarkdownInteractor.Delete() error = got %v\n", len(mds))
	}
}

//先に追加しておき、その後に取得して一致するかどうか
func TestFindByTitle(t *testing.T) {
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

	//testcaseの作成
	tests := []struct {
		name      string
		args      args
		wantTitle string //インスタンスを返す様にcreateを変更する？
	}{}

	for i := 0; i < 1000; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("test%d", i))
		s := fmt.Sprintf("%x", h.Sum(nil))
		t.Log(s)
		tests = append(tests, struct {
			name      string
			args      args
			wantTitle string
		}{fmt.Sprintf("normal %d", i), args{fmt.Sprintf("test%s", string(s)), fmt.Sprintf("/test%s/test%s", string(s), string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s))}, fmt.Sprintf("test%s", string(s))})
		err := markdownUsecase.Create(fmt.Sprintf("test%s", string(s)), fmt.Sprintf("/test%s/test%s", string(s), string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s)))
		if err != nil {
			t.Errorf("MarkdownInteractor.Delete() error = got %v, s value = %s\n", err, string(s))
		}
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	//usecaseの実行
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mds, err := markdownUsecase.FindByTitle(tt.args.title)
			if err != nil {
				t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
			}
			/*
				if len(mds) != 1 {
					t.Errorf("MarkdownInteractor.Delete() error = got %v\n", len(mds))
				}
			*/
			if mds[0].Title != tt.wantTitle {
				t.Errorf("MarkdownInteractor.Delete() error = got %v, want %v\n", mds[0].Title, tt.wantTitle)
			}

		})
	}
}

//先に追加しておき、その後に取得して一致するかどうか
//testcaseと一致するか？
func TestFindById(t *testing.T) {
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

	//testcaseの作成
	tests := []struct {
		id        int
		name      string
		args      args
		wantTitle string //インスタンスを返す様にcreateを変更する？
	}{}

	for i := 0; i < 1000; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("test%d", i))
		s := fmt.Sprintf("%x", h.Sum(nil))
		t.Log(s)
		tests = append(tests, struct {
			id        int
			name      string
			args      args
			wantTitle string
		}{i + 1, fmt.Sprintf("normal %d", i), args{fmt.Sprintf("test%s", string(s)), fmt.Sprintf("/test%s/test%s", string(s), string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s))}, fmt.Sprintf("test%s", string(s))})
		err := markdownUsecase.Create(fmt.Sprintf("test%s", string(s)), fmt.Sprintf("/test%s/test%s", string(s), string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s)))
		if err != nil {
			t.Errorf("MarkdownInteractor.Delete() error = got %v, s value = %s\n", err, string(s))
		}
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	//usecaseの実行
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := markdownUsecase.FindById(fmt.Sprintf("%d", tt.id))
			if err != nil {
				t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
			}
			if got.Id != tt.id {
				t.Errorf("MarkdownInteractor.Delete() error = got %v, want %v\n", got.Id, tt.id)
			}

			if got.Title != tt.args.title {
				t.Errorf("MarkdownInteractor.Delete() error = got %v, want %v\n", got.Title, tt.args.title)
			}
		})
	}
}

//先に追加しておき、その後に取得して一致するかどうか
func TestFindByPath(t *testing.T) {
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

	//testcaseの作成
	tests := []struct {
		name     string
		args     args
		wantPath string //インスタンスを返す様にcreateを変更する？
	}{}

	for i := 0; i < 1000; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("test%d", i))
		s := fmt.Sprintf("%x", h.Sum(nil))
		t.Log(s)
		tests = append(tests, struct {
			name     string
			args     args
			wantPath string
		}{fmt.Sprintf("normal %d", i), args{fmt.Sprintf("test%s", string(s)), fmt.Sprintf("/test%s/test%s", string(s), string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s))}, fmt.Sprintf("/test%s/test%s", string(s), string(s))})
		err := markdownUsecase.Create(fmt.Sprintf("test%s", string(s)), fmt.Sprintf("/test%s/test%s", string(s), string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s)))
		if err != nil {
			t.Errorf("MarkdownInteractor.Delete() error = got %v, s value = %s\n", err, string(s))
		}
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	//usecaseの実行
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mds, err := markdownUsecase.FindByTitle(tt.args.title)
			if err != nil {
				t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
			}

			if len(mds) != 1 {
				t.Errorf("MarkdownInteractor.Delete() error = got %v\n", len(mds))
			}

			if mds[0].Path != tt.wantPath {
				t.Errorf("MarkdownInteractor.Delete() error = got %v, want %v\n", mds[0].Title, tt.wantPath)
			}

		})
	}
}

//go test src/server/usecases -update
//go test ./... -update では動かない
var update = flag.Bool("update", false, "update golden files")

//簡易的にDBを作って、パスが存在しない場合は削除 testdata
func TestDeleteIfNotExistsByPath(t *testing.T) {
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
		name     string
		args     args
		wantPath string //インスタンスを返す様にcreateを変更する？
	}{}

	if *update {
		os.RemoveAll("testdata")
	}

	for i := 0; i < 100; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("test%d", i))
		s := fmt.Sprintf("%x", h.Sum(nil))
		t.Log(s)
		tests = append(tests, struct {
			name     string
			args     args
			wantPath string
		}{fmt.Sprintf("normal %d", i), args{fmt.Sprintf("test%s", string(s)), fmt.Sprintf("test%s/test%s", string(s), string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s))}, fmt.Sprintf("/test%s/test%s", string(s), string(s))})
		err := markdownUsecase.Create(fmt.Sprintf("test%s", string(s)), fmt.Sprintf("test%s/test%s", string(s), string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s)))
		if err != nil {
			t.Errorf("MarkdownInteractor.Delete() error = got %v, s value = %s\n", err, string(s))
		}
		//go test src/server/usecases -update
		if *update {
			os.MkdirAll(filepath.Join("testdata", string(s)), 0777)
			file, err := os.Create(filepath.Join("testdata", string(s), "README.md"))
			if err != nil {
				t.Fatal(err)
			}
			_, err = io.WriteString(file, fmt.Sprintf("%stest%d%s", string(s), i, string(s)))
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	//usecaseの実行
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := markdownUsecase.DeleteIfNotExistsByPath()
			if err != nil {
				t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
			}
			mds, err := markdownUsecase.FindByPath(tt.args.path)
			if err != nil {
				t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
			}
			if (len(mds)) != 0 {
				t.Errorf("MarkdownInteractor.Delete() error = got %v\n", len(mds))
			}

		})
	}
	err = markdownUsecase.DeleteIfNotExistsByPath()
	if err != nil {
		t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
	}
	mds, err := markdownUsecase.FindAll()
	if err != nil {
		t.Errorf("MarkdownInteractor.Delete() error = got %v\n", err)
	}
	if (len(mds)) != 0 {
		t.Errorf("MarkdownInteractor.Delete() error = got %v\n", len(mds))
	}
}

/*
//先に追加しておき、その後に取得して一致するかどうか
func TestSearchByTitle(t *testing.T) {
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

	//testcaseの作成
	tests := []struct {
		name string
		args args
		want bool //インスタンスを返す様にcreateを変更する？
	}{
		{"normal 1", args{"test1", "/test1/test1", "http://test1/test1"}, true},
		{"fail SrcURL = isNotURL", args{"test2", "/test2/test2", "test1/test1"}, false},
		{"fail title = isEmpty", args{"", "/test2/test2", "test1/test1"}, false},
		{"fail path = isEmpty", args{"test3", "", "test1/test1"}, false},
		{"fail srcUrl = isEmpty", args{"test4", "test4", ""}, false},
		{"fail allField = isEmpty", args{"", "", ""}, false},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	//usecaseの実行
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

		})
	}
}
*/

/*
//先に追加しておき、その後に取得して一致するかどうか testdata
func TestSearchByBody(t *testing.T) {
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

	//testcaseの作成
	tests := []struct {
		name string
		args args
		want bool //インスタンスを返す様にcreateを変更する？
	}{
		{"normal 1", args{"test1", "/test1/test1", "http://test1/test1"}, true},
		{"fail SrcURL = isNotURL", args{"test2", "/test2/test2", "test1/test1"}, false},
		{"fail title = isEmpty", args{"", "/test2/test2", "test1/test1"}, false},
		{"fail path = isEmpty", args{"test3", "", "test1/test1"}, false},
		{"fail srcUrl = isEmpty", args{"test4", "test4", ""}, false},
		{"fail allField = isEmpty", args{"", "", ""}, false},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	//usecaseの実行
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

		})
	}
}

//testdataのファイルをzipにして、それを解凍して、その中身を取得して一致するかどうか
func TestCreateZipFile(t *testing.T) {
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

	//testcaseの作成
	tests := []struct {
		name string
		args args
		want bool //インスタンスを返す様にcreateを変更する？
	}{
		{"normal 1", args{"test1", "/test1/test1", "http://test1/test1"}, true},
		{"fail SrcURL = isNotURL", args{"test2", "/test2/test2", "test1/test1"}, false},
		{"fail title = isEmpty", args{"", "/test2/test2", "test1/test1"}, false},
		{"fail path = isEmpty", args{"test3", "", "test1/test1"}, false},
		{"fail srcUrl = isEmpty", args{"test4", "test4", ""}, false},
		{"fail allField = isEmpty", args{"", "", ""}, false},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	//usecaseの実行
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

		})
	}
}
*/
