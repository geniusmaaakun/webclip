package models

//DBを作成する sqlite3を使用
import (
	"database/sql"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

func GetDatabasePath() (string, error) {
	var folderPath string

	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	// 実行中のOSを判別
	switch runtime.GOOS {
	case "windows":
		// C:\Users\username\AppData\Local\AppName
		folderPath = filepath.Join(currentUser.HomeDir, "AppData", "Local", "webclip")
	case "darwin":
		// ユーザーのホームディレクトリ内のLibrary/Cachesフォルダに保存
		// /Users/aoyagimasanori/Library/Caches/webclip
		folderPath = filepath.Join(currentUser.HomeDir, "Library", "Caches", "webclip")
	case "linux":
		// /var/lib/AppName
		folderPath = filepath.Join("/var/lib", "webclip")
	default:
		return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
	fmt.Println(folderPath)
	// ディレクトリが存在しない場合は作成
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.MkdirAll(folderPath, 0755)
		if err != nil {
			return "", err
		}
	}
	return folderPath, nil
}

func NewDB(dbPath string) (*sql.DB, error) {
	//1
	//DBを開く  なければ作成される
	//dbの場所はぜったいパス
	db, err := sql.Open("sqlite3", dbPath)
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
