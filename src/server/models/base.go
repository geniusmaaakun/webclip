package models

//DBを作成する sqlite3を使用
import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB() (*sql.DB, error) {
	//1
	//DBを開く  なければ作成される
	//dbの場所はぜったいパス
	db, err := sql.Open("sqlite3", "webclip.sql")
	//エラーハンドリング
	if err != nil {
		return nil, err
	}

	//DB作成 SQLコマンド
	query := `CREATE TABLE IF NOT EXISTS markdown_memo (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title String NOT NULL,
		path String UNIQUE NOT NULL,
		src_url String NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP)`

	//実行 結果は返ってこない為、_にする
	_, err = db.Exec(query)

	//エラーハンドリング
	if err != nil {
		return nil, err
	}

	//defer Db.Close()

	//疎通確認
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
