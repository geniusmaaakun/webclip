package main

import (
	"log"
	"os"
	"webclip/src/actions"
)

//このコードでは、`urfave/cli/v2`ライブラリを使用してコマンドラインオプションを定義し、`Action`関数で処理を実行しています。このライブラリを使用することで、コマンドラインオプションの定義が簡潔になり、エラーチェックやヘルプメッセージの生成などが自動化されます。
//上記のコードでは、`HTMLToMarkdownConverter`と`wcdownloader`という構造体を使用して、処理を行っています。これにより、コードがより構造化され、プロのエンジニアが書いたようなスタイルになっています。

func main() {
	//production
	// folderPath, err := models.GetDatabasePath()
	// if err != nil {
	// 	log.Fatalf("main: %v\n", err)
	// }
	// fmt.Println(folderPath)

	// app := actions.NewWebClip(filepath.Join(folderPath, "webclip.sql"))

	//development
	app := actions.NewWebClip("webclip.sql")

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
