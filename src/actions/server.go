package actions

import (
	"fmt"
	"log"
	"path/filepath"
	"webclip/src/server/controllers"
	"webclip/src/server/models"

	"github.com/urfave/cli/v2"
)

func Server(c *cli.Context) error {
	//コマンド入力待機
	//search 特定の文字列を検索
	//clear ファイルパスが存在しない場合削除
	//list ファイルパスを表示
	fmt.Println("Start Web Server")
	//create db
	folderPath, err := models.GetDatabasePath()
	if err != nil {
		log.Fatalf("SaveDatabase: %v\n", err)
	}
	db, err := models.NewDB(filepath.Join(folderPath, "webclip.sql"))
	//create usecase
	//create handler
	srv := controllers.NewServer("localhost", "8080", db)
	err = srv.Run()
	if err != nil {
		log.Fatalf("SaveDatabase: %v\n", err)
	}
	return nil
}
