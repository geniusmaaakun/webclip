package main

import (
	"log"
	"os"
	"webclip/src/actions"
	"webclip/src/server/models"

	"github.com/urfave/cli/v2"
)

//このコードでは、`urfave/cli/v2`ライブラリを使用してコマンドラインオプションを定義し、`Action`関数で処理を実行しています。このライブラリを使用することで、コマンドラインオプションの定義が簡潔になり、エラーチェックやヘルプメッセージの生成などが自動化されます。
//上記のコードでは、`HTMLToMarkdownConverter`と`wcdownloader`という構造体を使用して、処理を行っています。これにより、コードがより構造化され、プロのエンジニアが書いたようなスタイルになっています。

func main() {
	_, err := models.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	app := &cli.App{
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
				Name:    "db",
				Aliases: []string{"save"},
				Usage:   "Save to DB",
			},
		},
		//WebClip
		Action: actions.Download,
	}

	app.Commands = []*cli.Command{
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

		//すべてのファイルを削除する　危険
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
