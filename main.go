package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"webclip/src/server/models"
	"webclip/src/wcconverter"
	"webclip/src/wcdownloader"

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
		},
		//WebClip
		Action: func(c *cli.Context) error {
			//1 setup main
			//コマンドラインオプションの値を取得
			url := c.String("url")
			outdir := c.String("outdir")
			imageDownloadFlag := c.Bool("download")

			if (outdir == "") || (url == "") {
				log.Println("Invalid arguments, usage: webclip -u <url> -o <outdir> [-d]")
				log.Fatal("more info: webclip -h")
			}

			err := os.MkdirAll(outdir, 0755)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Target: %s\n", url)

			downloader := wcdownloader.NewDownloader(url, outdir, imageDownloadFlag)
			doc, err := downloader.HtmlDownloader.CreateDocument()
			if err != nil {
				log.Fatalf("create Document(): %s", err.Error())
			}
			doc, err = downloader.ImageDownloader.ReplaceImageSrcToImageFile(doc)
			if err != nil {
				log.Fatalf("replace Document(): %s", err.Error())
			}

			converter := wcconverter.NewConverter(outdir, "README.md", nil)
			markdown, err := converter.Convert(doc.Selection)
			if err != nil {
				log.Fatalf("Markdown Conversion Error: %s", err.Error())
			}
			err = converter.SaveToFile(markdown)
			if err != nil {
				log.Fatalf("Markdown Conversion Error: %s", err.Error())
			}

			db, err := models.NewDB()
			repo := models.NewMarkdownRepo(db)
			absPath, err := filepath.Abs(filepath.Join(outdir, "README.md"))
			if err != nil {
				log.Fatalf("SaveDatabase: %v\n", err)
			}
			mdData := models.NewMarkdownMemo(outdir, absPath, url)
			err = repo.Create(mdData)
			if err != nil {
				log.Fatalf("SaveDatabase: %v\n", err)
			}

			return nil
		},
	}

	//sub command : webclip serverd
	app.Commands = []*cli.Command{
		{
			Name:  "server",
			Usage: "Start Web Server",
			Action: func(c *cli.Context) error {
				//コマンド入力待機
				//search 特定の文字列を検索
				//clear ファイルパスが存在しない場合削除
				//list ファイルパスを表示
				fmt.Println("Start Web Server")
				return nil
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
