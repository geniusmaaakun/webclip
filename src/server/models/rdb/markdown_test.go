package rdb_test

import (
	"os"
	"strconv"
	"testing"
	"webclip/src/server/models"
	"webclip/src/server/models/rdb"
)

/*
データベースに対するテストケースを実施する際には、様々なシナリオを考慮することが重要です。以下に、一般的なテストケースをいくつか挙げます：

1 正常系（成功）:

データベースへのレコードの正常な挿入、更新、削除、取得など、基本的なCRUD操作が正常に動作することを確認するテストケース。

2 異常系（失敗）:

入力データが不正な場合の挙動（例：NULLや空文字列、制約違反、型の不整合など）。
SQLクエリエラー（例：構文エラーやテーブル名・カラム名の間違い）。
データベース接続エラー（例：接続が失敗した場合やタイムアウト）。
一意性制約違反（例：一意であるべきカラムに重複したデータを挿入しようとする場合）。
参照制約違反（例：親テーブルのデータを参照している子テーブルのデータを削除しようとする場合）。

3 境界値テスト:

最大・最小のデータ長や範囲を超えるデータを入力し、エラーが適切に発生することを確認するテストケース。

4 同時実行（並行性）テスト:

複数のクライアントが同時にデータベースにアクセスした場合の挙動を確認するテストケース。例えば、同じレコードに対する同時更新や、トランザクションの競合など。

5 トランザクションテスト:

トランザクションが正しく動作し、適切なロールバックやコミットが行われることを確認するテストケース。

6 パフォーマンステスト:

データベースの応答時間や負荷状況を評価するテストケース。これには、大量のデータを処理する場合や高負荷状況下でのデータベースの挙動を確認することが含まれます。


これらのテストケースを検討し、アプリケーションの要件に応じて適切なテストシナリオを実装することが、データベースの堅牢性と正確性を保証するために重要です。ただし、すべてのケースをカバーする必要はありません。代わりに、アプリケーションの要件、データベースの設計、およびリスクの高い部分に焦点を当ててテストを計画し、実装してください。

また、データベースのテストを行う際には、モックやスタブなどのテストダブルを使用して、実際のデータベースへの依存を排除し、テストの速度と信頼性を向上させることが推奨されます。ただし、一部のケースでは、統合テストを行い、実際のデータベースとの連携が正しく機能していることを確認することも重要です。

データベーステストの実施により、バグや不具合を早期に発見し、アプリケーションの信頼性を向上させることができます。適切なテスト戦略を立てて実行することで、データベースの品質を維持し、アプリケーション全体の安定性を確保できます。
*/

func TestNewDB(t *testing.T) {
	_, err := models.NewDB()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateMd(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}

	txManager := rdb.NewTransactionManager(db)
	MarkdownRepo := rdb.NewMarkdownRepo()

	type args struct {
		title  string
		path   string
		srcURL string
	}
	//testcase
	tests := []struct {
		name string
		args args
		want *models.MarkdownMemo
	}{
		{name: "name", args: args{"test", "test", "test"}, want: &models.MarkdownMemo{Title: "test", Path: "test", SrcUrl: "test"}},
		{name: "name2", args: args{"test2", "test2", "test2"}, want: &models.MarkdownMemo{Title: "test2", Path: "test2", SrcUrl: "test2"}},

		//書くパラメーターが空の場合　別のテストでやる usecaseでバリデーションをかける
		//{name: "null", args: args{"", "test3", "test3"}, want: nil},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tx, err := txManager.NewTransaction(false)
			if err != nil {
				t.Fatal(err)
			}
			md := models.NewMarkdownMemo(tt.args.title, tt.args.path, tt.args.srcURL)
			err = MarkdownRepo.Create(tx, md)
			if err != nil {
				t.Fatal(err)
			}
			got, err := MarkdownRepo.FindByTitleLastOne(tx, md.Title)
			if err != nil {
				t.Fatal(err)
			}
			if got.Title != tt.want.Title {
				t.Errorf("CreateMd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateMdUnique(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}

	txManager := rdb.NewTransactionManager(db)
	MarkdownRepo := rdb.NewMarkdownRepo()

	type args struct {
		title  string
		path   string
		srcURL string
	}
	//testcase
	tests := []struct {
		name  string
		args1 args
		args2 args
	}{
		{name: "name", args1: args{"test", "test", "test"}, args2: args{"test", "test", "test"}},
		{name: "name2", args1: args{"test2", "test2", "test2"}, args2: args{"test2", "test2", "test2"}},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tx, err := txManager.NewTransaction(false)
			if err != nil {
				t.Fatal(err)
			}
			md1 := models.NewMarkdownMemo(tt.args1.title, tt.args1.path, tt.args1.srcURL)
			err = MarkdownRepo.Create(tx, md1)
			if err != nil {
				t.Fatal(err)
			}
			md2 := models.NewMarkdownMemo(tt.args2.title, tt.args2.path, tt.args2.srcURL)
			err = MarkdownRepo.Create(tx, md2)
			if err == nil {
				t.Fatal(err)
			}
		})
	}
}

//sqlite3では最大文字数が不明なため、テストできない。
/*
func TestCreateMdMax(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}

	txManager := rdb.NewTransactionManager(db)
	MarkdownRepo := rdb.NewMarkdownRepo()

	type args struct {
		title  string
		path   string
		srcURL string
	}
	//testcase
	tests := []struct {
		name string
		args args
		want bool
	}{
		//タイトルが違うのでUNIQUEに引っかかりテストできない。
		{name: "title max over", args: args{strings.Repeat("a", math.MaxInt16*30500), "test1", "test1"}, want: false},
		//{name: "titile max - 1", args: args{strings.Repeat("a", 65535), "test1", "test1"}, want: true},
		//{name: "path max over", args: args{"test2", strings.Repeat("a", 655360), "test2"}, want: false},
		//{name: "path max -1", args: args{"test2", strings.Repeat("a", 65535), "test2"}, want: true},
		//{name: "src max over", args: args{"test3", "test3", strings.Repeat("a", 655360)}, want: false},
		//{name: "src max -1", args: args{"test3", "test3", strings.Repeat("a", 65535)}, want: true},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()

			tx, err := txManager.NewTransaction(false)
			if err != nil {
				t.Fatal(err)
			}
			md := models.NewMarkdownMemo(tt.args.title, tt.args.path, tt.args.srcURL)
			err = MarkdownRepo.Create(tx, md)
			if err != nil {
				t.Fatal(err)
			}
				if tt.want {
					if err != nil {
						t.Fatal(err)
					}
				} else {
					if err == nil {
						t.Fatalf("CreateMdMax() want error: %v\n", tt.name)
					}
				}
			_, err = MarkdownRepo.FindByTitleLastOne(tx, md.Title)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(math.MaxInt16 * 30500)
			//t.Log(len(got.Title))
			//t.Log(got)
			err = MarkdownRepo.DeleteByTitle(tx, md)
			if err != nil {
				t.Fatal(err)
			}

			//got, err := MarkdownRepo.FindByTitleLastOne(tx, md.Title)
			//t.Log(len(got.Title))
		})
	}
}
*/

func TestCreateMdInjection(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}

	txManager := rdb.NewTransactionManager(db)
	MarkdownRepo := rdb.NewMarkdownRepo()

	type args struct {
		title  string
		path   string
		srcURL string
	}
	//testcase
	tests := []struct {
		name string
		args args
	}{
		//insert文への追加insert文
		{name: "injection 1", args: args{"test1", "test1", "'test1', 2023-04-28 11:47:16.789825 +0900 +0900); INSERT INTO markdown_memo (title, path, src_url) VALUES ('test2', 'test2, 'test2'"}},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tx, err := txManager.NewTransaction(false)
			if err != nil {
				t.Fatalf("CreateMdinjection(): %s\n", err)
			}
			md := models.NewMarkdownMemo(tt.args.title, tt.args.path, tt.args.srcURL)
			err = MarkdownRepo.Create(tx, md)
			if err != nil {
				t.Fatalf("CreateMdinjection(): %s\n", err)

			}
			mds, err := MarkdownRepo.FindAll(tx)
			if err != nil {
				t.Fatalf("CreateMdinjection(): %s\n", err)
			}
			if (len(mds)) != 1 {
				t.Fatalf("CreateMdinjection(): injecstion succes!!")
			}
		})
	}
}

//mockを使ったテスト
/*
import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"your_project_path/models"
	"your_project_path/usecases"
)

// トランザクションのモックを定義
type MockTransaction struct {
	mock.Mock
}

func (m *MockTransaction) Prepare(query string) (usecases.Stmt, error) {
	args := m.Called(query)
	return args.Get(0).(usecases.Stmt), args.Error(1)
}

// Stmtのモックを定義
type MockStmt struct {
	mock.Mock
}

func (m *MockStmt) Exec(args ...interface{}) (usecases.Result, error) {
	callArgs := m.Called(args...)
	return callArgs.Get(0).(usecases.Result), callArgs.Error(1)
}

func TestMarkdownManager_Create(t *testing.T) {
	// 1. テストケースを作成
	md := &models.MarkdownMemo{
		Title:     "Test Title",
		Path:      "/test/path",
		SrcUrl:    "https://example.com/test",
		CreatedAt: time.Now(),
	}

	t.Run("Create success", func(t *testing.T) {
		tx := new(MockTransaction)
		stmt := new(MockStmt)

		// 2. モックを設定
		tx.On("Prepare", "INSERT INTO markdown_memo (title, path, src_url, created_at) VALUES (?, ?, ?, ?)").Return(stmt, nil)
		stmt.On("Exec", md.Title, md.Path, md.SrcUrl, md.CreatedAt).Return(nil, nil)

		// 3. 関数を呼び出し
		m := &MarkdownManager{}
		err := m.Create(tx, md)

		// 4. エラーが発生しないことを確認
		assert.Nil(t, err)
		tx.AssertExpectations(t)
		stmt.AssertExpectations(t)
	})

	t.Run("Create failure", func(t *testing.T) {
		tx := new(MockTransaction)

		// 2. モックを設定
		tx.On("Prepare", "INSERT INTO markdown_memo (title, path, src_url, created_at) VALUES (?, ?, ?, ?)").Return(nil, errors.New("test error"))

		// 3. 関数を呼び出し
		m := &MarkdownManager{}
		err := m.Create(tx, md)

		// 4. エラーが発生することを確認
		assert.NotNil(t, err)
		assert.Equal(t, "test error", err.Error())
		tx.AssertExpectations(t)
	})
}

*/

func TestDeleteMd(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}
	txManager := rdb.NewTransactionManager(db)

	MarkdownRepo := rdb.NewMarkdownRepo()

	type args struct {
		title  string
		path   string
		srcURL string
	}
	//testcase
	tests := []struct {
		name string
		args args
		want *models.MarkdownMemo
	}{
		{name: "name", args: args{"test", "test", "test"}, want: &models.MarkdownMemo{Title: "test", Path: "test", SrcUrl: "test"}},
		{name: "name2", args: args{"test2", "test2", "test2"}, want: &models.MarkdownMemo{Title: "test2", Path: "test2", SrcUrl: "test2"}},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tx, err := txManager.NewTransaction(false)
			if err != nil {
				t.Fatal(err)
			}
			md := models.NewMarkdownMemo(tt.args.title, tt.args.path, tt.args.srcURL)
			err = MarkdownRepo.Create(tx, md)
			if err != nil {
				t.Fatal(err)
			}
			target, err := MarkdownRepo.FindByTitleLastOne(tx, md.Title)
			if err != nil {
				t.Fatal(err)
			}
			err = MarkdownRepo.Delete(tx, target)
			if err != nil {
				t.Fatal(err)
			}
			got, err := MarkdownRepo.FindByTitleLastOne(tx, md.Title)
			if err != nil {
				t.Fatal(err)
			}
			if got != nil {
				t.Errorf("DeleteMd() = %v, want %v", got, "NULL")
			}
		})
	}
}

func TestDeleteByTitle(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}
	txManager := rdb.NewTransactionManager(db)

	MarkdownRepo := rdb.NewMarkdownRepo()

	type args struct {
		title  string
		path   string
		srcURL string
	}
	//testcase
	tests := []struct {
		name string
		args args
		want *models.MarkdownMemo
	}{
		{name: "name", args: args{"test", "test", "test"}, want: &models.MarkdownMemo{Title: "test", Path: "test", SrcUrl: "test"}},
		{name: "name2", args: args{"test2", "test2", "test2"}, want: &models.MarkdownMemo{Title: "test2", Path: "test2", SrcUrl: "test2"}},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tx, err := txManager.NewTransaction(false)
			if err != nil {
				t.Fatal(err)
			}
			md := models.NewMarkdownMemo(tt.args.title, tt.args.path, tt.args.srcURL)
			// err = MarkdownRepo.Create(tx, md)
			// if err != nil {
			// 	t.Fatal(err)
			// }

			//データがなくてもエラーは返さない
			err = MarkdownRepo.DeleteByTitle(tx, md)
			if err != nil {
				t.Fatal(err)
			}
			//データがなくてもエラーは返さない
			got, err := MarkdownRepo.FindByTitleLastOne(tx, md.Title)
			if err != nil {
				t.Fatal(err)
			}
			if got != nil {
				t.Errorf("DeleteMd() = %v, want %v", got, "NULL")
			}
		})
	}
}
func TestDeleteByTitleInNotExistData(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}
	txManager := rdb.NewTransactionManager(db)

	MarkdownRepo := rdb.NewMarkdownRepo()

	type args struct {
		title  string
		path   string
		srcURL string
	}
	//testcase
	tests := []struct {
		name string
		args args
		want *models.MarkdownMemo
	}{
		{name: "name", args: args{"test", "test", "test"}, want: &models.MarkdownMemo{Title: "test", Path: "test", SrcUrl: "test"}},
		{name: "name2", args: args{"test2", "test2", "test2"}, want: &models.MarkdownMemo{Title: "test2", Path: "test2", SrcUrl: "test2"}},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tx, err := txManager.NewTransaction(false)
			if err != nil {
				t.Fatal(err)
			}
			md := models.NewMarkdownMemo(tt.args.title, tt.args.path, tt.args.srcURL)
			err = MarkdownRepo.Create(tx, md)
			if err != nil {
				t.Fatal(err)
			}
			err = MarkdownRepo.DeleteByTitle(tx, md)
			if err != nil {
				t.Fatal(err)
			}
			got, err := MarkdownRepo.FindByTitleLastOne(tx, md.Title)
			if err != nil {
				t.Fatal(err)
			}
			if got != nil {
				t.Errorf("DeleteMd() = %v, want %v", got, "NULL")
			}
		})
	}
}

//unused
func TestUpdateMd(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}
	txManager := rdb.NewTransactionManager(db)
	MarkdownRepo := rdb.NewMarkdownRepo()

	type args struct {
		title  string
		path   string
		srcURL string
	}
	//testcase
	tests := []struct {
		name string
		args args
		want *models.MarkdownMemo
	}{
		{name: "name", args: args{"test", "test", "test"}, want: &models.MarkdownMemo{Title: "test", Path: "test", SrcUrl: "test"}},
		{name: "name2", args: args{"test2", "test2", "test2"}, want: &models.MarkdownMemo{Title: "test2", Path: "test2", SrcUrl: "test2"}},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tx, err := txManager.NewTransaction(false)
			md := models.NewMarkdownMemo(tt.args.title, tt.args.path, tt.args.srcURL)
			err = MarkdownRepo.Create(tx, md)
			if err != nil {
				t.Fatal(err)
			}

			updateMd, err := MarkdownRepo.FindByTitleLastOne(tx, md.Title)
			if err != nil {
				t.Fatal(err)
			}
			updateMd.Title = "testupdate"
			updateMd.SrcUrl = "srcupdate"
			updateMd.Path = "pathupdate"

			err = MarkdownRepo.Update(tx, updateMd)

			got, err := MarkdownRepo.FindByTitleLastOne(tx, updateMd.Title)
			if err != nil {
				t.Fatal(err)
			}

			if got.Title != updateMd.Title {
				t.Errorf("DeleteMd() = %v, want %v", got.Title, updateMd.Title)
			}

			if got.Path != updateMd.Path {
				t.Errorf("DeleteMd() = %v, want %v", got.Title, updateMd.Title)
			}

			if got.SrcUrl != updateMd.SrcUrl {
				t.Errorf("DeleteMd() = %v, want %v", got.Title, updateMd.Title)
			}
		})
	}
}

//存在しないデータの更新はエラーを返す

//injection系
func TestFindById(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}
	txManager := rdb.NewTransactionManager(db)
	MarkdownRepo := rdb.NewMarkdownRepo()

	type args struct {
		title  string
		path   string
		srcURL string
	}

	//testcase
	tests := []struct {
		name string
		args args
		id   int
		want *models.MarkdownMemo
	}{}

	for i := 1; i <= 100; i++ {
		tests = append(tests, struct {
			name string
			args args
			id   int
			want *models.MarkdownMemo
		}{name: "name", args: args{"test" + strconv.Itoa(i), "test" + strconv.Itoa(i), "test" + strconv.Itoa(i)}, id: i, want: &models.MarkdownMemo{Title: "test" + strconv.Itoa(i), Path: "test" + strconv.Itoa(i), SrcUrl: "test" + strconv.Itoa(i)}})
	}
	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			//IDが1から100までのデータを作成。順番が大事なので並列処理はしない
			//t.Parallel()
			tx, err := txManager.NewTransaction(false)
			if err != nil {
				t.Fatal(err)
			}
			md := models.NewMarkdownMemo(tt.args.title, tt.args.path, tt.args.srcURL)
			err = MarkdownRepo.Create(tx, md)
			if err != nil {
				t.Fatal(err)
			}
			got, err := MarkdownRepo.FindById(tx, tt.id)
			if err != nil {
				t.Fatal(err)
			}
			if got.Title != tt.want.Title {
				t.Errorf("FindByIdMd() = %v, want %v", got.Title, tt.want.Title)
			}
		})
	}
}

func TestFindAll(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}
	txManager := rdb.NewTransactionManager(db)
	MarkdownRepo := rdb.NewMarkdownRepo()

	type args struct {
		title  string
		path   string
		srcURL string
	}

	//testcase
	tests := []struct {
		name string
		args args
		id   int
		want int
	}{}

	for i := 1; i <= 100; i++ {
		tests = append(tests, struct {
			name string
			args args
			id   int
			want int
		}{name: "name", args: args{"test" + strconv.Itoa(i), "test" + strconv.Itoa(i), "test" + strconv.Itoa(i)}, id: i, want: i})
	}
	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			//IDが1から100までのデータを作成。順番が大事なので並列処理はしない
			//t.Parallel()
			tx, err := txManager.NewTransaction(false)
			if err != nil {
				t.Fatal(err)
			}
			md := models.NewMarkdownMemo(tt.args.title, tt.args.path, tt.args.srcURL)
			err = MarkdownRepo.Create(tx, md)
			if err != nil {
				t.Fatal(err)
			}
			got, err := MarkdownRepo.FindAll(tx)
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != tt.want {
				t.Errorf("FindByIdMd() = %v, want %v", len(got), tt.id)
			}
		})
	}
}

func TestFindByTitle(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}
	txManager := rdb.NewTransactionManager(db)
	MarkdownRepo := rdb.NewMarkdownRepo()

	type args struct {
		title  string
		path   string
		srcURL string
	}
	//testcase
	tests := []struct {
		name string
		args args
		want *models.MarkdownMemo
	}{
		{name: "name", args: args{"test", "test", "test"}, want: &models.MarkdownMemo{Title: "test", Path: "test", SrcUrl: "test"}},
		{name: "name2", args: args{"test2", "test2", "test2"}, want: &models.MarkdownMemo{Title: "test2", Path: "test2", SrcUrl: "test2"}},

		//空でも入ってしまう
		//{name: "empty title", args: args{"", "test3", "test3"}, want: nil},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tx, err := txManager.NewTransaction(false)
			if err != nil {
				t.Fatal(err)
			}
			md := models.NewMarkdownMemo(tt.args.title, tt.args.path, tt.args.srcURL)
			err = MarkdownRepo.Create(tx, md)
			if err != nil {
				t.Fatal(err)
			}
			got, err := MarkdownRepo.FindByTitle(tx, tt.args.title)
			if err != nil {
				t.Fatal(err)
			}
			if tt.want != nil && got != nil {
				if got[0].Title != tt.want.Title {
					t.Errorf("FindByIdMd() = got %v, want %v", got[0].Title, tt.want.Title)
				}
			} else {
				if got != nil {
					t.Errorf("FindByIdMd() = %v, want %v", got[0].Title, tt.want)
				}
			}
		})
	}
}

func TestFindByTitleInjection(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}
	txManager := rdb.NewTransactionManager(db)
	MarkdownRepo := rdb.NewMarkdownRepo()

	type args struct {
		title  string
		path   string
		srcURL string
	}

	tx, err := txManager.NewTransaction(false)
	if err != nil {
		t.Fatal(err)
	}
	//testcase
	for i := 1; i <= 100; i++ {
		err = MarkdownRepo.Create(tx, &models.MarkdownMemo{Title: "test" + strconv.Itoa(i), Path: "test" + strconv.Itoa(i), SrcUrl: "test" + strconv.Itoa(i)})
		if err != nil {
			t.Fatal(err)
		}
	}
	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	//SQLインジェクションのテスト
	mds, err := MarkdownRepo.FindByTitle(tx, "test1' OR '1' = '1")
	if err != nil {
		t.Fatal(err)
	}
	if len(mds) != 0 {
		t.Errorf("FindByTitleInjection() = got %v, want %v", len(mds), 1)
	}

}

func TestSearchByTitle(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}
	txManager := rdb.NewTransactionManager(db)
	MarkdownRepo := rdb.NewMarkdownRepo()

	type args struct {
		title  string
		path   string
		srcURL string
	}

	testdata := []args{
		{"test", "test", "test"},
		{"test2", "test2", "test2"},
		{"test3", "test3", "test3"},
		{"test4", "test4", "test4"},
		{"test5", "test5", "test5"},
		{"test6", "test6", "test6"},
	}
	//testcase
	tests := []struct {
		name    string
		title   string
		wantLen int
		want    []*models.MarkdownMemo
	}{
		{name: "test length = 1", title: "test2", wantLen: 1, want: []*models.MarkdownMemo{{Title: "test2", Path: "test2", SrcUrl: "test2"}}},
		{name: "test length = 0", title: "a", wantLen: 0, want: nil},

		{name: "test length = 6", title: "t", wantLen: 6, want: []*models.MarkdownMemo{{Title: "test", Path: "test", SrcUrl: "test"},
			{Title: "test2", Path: "test2", SrcUrl: "test2"},
			{Title: "test3", Path: "test3", SrcUrl: "test3"},
			{Title: "test4", Path: "test4", SrcUrl: "test4"},
			{Title: "test5", Path: "test5", SrcUrl: "test5"},
			{Title: "test6", Path: "test6", SrcUrl: "test6"}}},

		{name: "test length = 6", title: "test", wantLen: 6,
			want: []*models.MarkdownMemo{{Title: "test", Path: "test", SrcUrl: "test"},
				{Title: "test2", Path: "test2", SrcUrl: "test2"},
				{Title: "test3", Path: "test3", SrcUrl: "test3"},
				{Title: "test4", Path: "test4", SrcUrl: "test4"},
				{Title: "test5", Path: "test5", SrcUrl: "test5"},
				{Title: "test6", Path: "test6", SrcUrl: "test6"}}},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	tx, err := txManager.NewTransaction(false)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range testdata {
		md := models.NewMarkdownMemo(v.title, v.path, v.srcURL)
		err = MarkdownRepo.Create(tx, md)
		if err != nil {
			t.Fatal(err)
		}
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			got, err := MarkdownRepo.SearchByTitle(tx, tt.title)
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != tt.wantLen {
				t.Fatalf("SearchByTitle() = got %v, want %v", len(got), len(testdata))
			}

			for i := range got {
				if got[i].Title != tt.want[i].Title {
					t.Errorf("SearchByTitle() = got %v, want %v", got[i].Title, tt.want[i].Title)
				}
			}
		})
	}
}
