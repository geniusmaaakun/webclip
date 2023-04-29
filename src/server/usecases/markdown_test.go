package usecases_test

import (
	"os"
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

//先に追加しておき、その後に取得して一致するかどうか
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

//簡易的にDBを作って、パスが存在しない場合は削除 testdata
func TestMarkdownInteractor_DeleteIfNotExistsByPath(t *testing.T) {
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
