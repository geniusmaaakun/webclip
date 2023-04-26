package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"webclip/src/server/controllers"
	"webclip/src/server/models"
	"webclip/src/server/usecases"
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
			&cli.BoolFlag{
				Name:    "db",
				Aliases: []string{"save"},
				Usage:   "Save to DB",
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

			//download
			downloader := wcdownloader.NewDownloader(url, outdir, imageDownloadFlag)
			doc, err := downloader.HtmlDownloader.CreateDocument()
			if err != nil {
				log.Fatalf("create Document(): %s", err.Error())
			}
			doc, err = downloader.ImageDownloader.ReplaceImageSrcToImageFile(doc)
			if err != nil {
				log.Fatalf("replace Document(): %s", err.Error())
			}

			//convert
			converter := wcconverter.NewConverter(outdir, "README.md", nil)
			markdown, err := converter.Convert(doc.Selection)
			if err != nil {
				log.Fatalf("Markdown Conversion Error: %s", err.Error())
			}
			markdown = converter.AddSrcUrlToMarkdown(url, markdown)
			err = converter.SaveToFile(markdown)
			if err != nil {
				log.Fatalf("Markdown Conversion Error: %s", err.Error())
			}

			//save to DB
			if c.Bool("save") {
				db, err := models.NewDB()
				if err != nil {
					log.Fatalf("SaveDatabase: %v\n", err)
				}
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
			}
			return nil

		},
	}

	app.Commands = []*cli.Command{
		//sub command : webclip server
		{
			Name:  "server",
			Usage: "Start Web Server",
			Action: func(c *cli.Context) error {
				//コマンド入力待機
				//search 特定の文字列を検索
				//clear ファイルパスが存在しない場合削除
				//list ファイルパスを表示
				fmt.Println("Start Web Server")
				//create db
				db, err := models.NewDB()
				if err != nil {
					log.Fatalf("SaveDatabase: %v\n", err)
				}
				//create usecase
				//create handler
				srv := controllers.NewServer("localhost", "8080", db)
				srv.Run()
				return nil
			},
		},
		//sub command : webclip clean
		{
			Name:  "clean",
			Usage: "clean database if file is not exist",
			Action: func(c *cli.Context) error {
				db, err := models.NewDB()
				if err != nil {
					log.Fatalf("CleanDatabase: %v\n", err)
				}
				markdownRepo := models.NewMarkdownRepo(db)
				markdownUsecase := usecases.NewMarkdownInteractor(markdownRepo)
				err = markdownUsecase.DeleteIfNotExistsByPath()
				if err != nil {
					log.Fatalf("CleanDatabase: %v\n", err)
				}

				return nil
			},
		},

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
			Action: func(c *cli.Context) error {
				title := c.String("title")
				body := c.String("body")
				if title == "" && body == "" {
					log.Fatal("title or body is required")
					return nil
				}

				db, err := models.NewDB()
				if err != nil {
					log.Fatalf("SearchDatabase: %v\n", err)
				}
				markdownRepo := models.NewMarkdownRepo(db)
				markdownUsecase := usecases.NewMarkdownInteractor(markdownRepo)

				if title != "" && body != "" {
					//find . | xargs grep -n hogehoge
					//bodyの行数も取得？

					markdownsByTitle, err := markdownUsecase.SearchByTitle(title)
					if err != nil {
						log.Fatalf("SearchDatabase: %v\n", err)
					}
					markdownsByBody, resultBodyMap, err := markdownUsecase.SearchByBody(body)
					if err != nil {
						log.Fatalf("SearchDatabase: %v\n", err)
					}

					mdPathMap := map[string]*models.MarkdownMemo{}

					for _, m := range markdownsByTitle {
						mdPathMap[m.Path] = m
					}
					for _, m := range markdownsByBody {
						mdPathMap[m.Path] = m
					}

					fmt.Println("Search Result")
					for _, m := range mdPathMap {
						fmt.Printf("id: %d, title: %s, path: %s, url: %s\n", m.Id, m.Title, m.Path, m.SrcUrl)
						for _, resultBody := range resultBodyMap[m.Path] {
							fmt.Printf("  %s\n", resultBody)
						}
					}

				} else if title != "" {
					markdowns, err := markdownUsecase.SearchByTitle(title)
					if err != nil {
						log.Fatalf("SearchDatabase: %v\n", err)
					}
					fmt.Println("Search Result")
					for _, m := range markdowns {
						fmt.Printf("id: %d, title: %s, path: %s, url: %s\n", m.Id, m.Title, m.Path, m.SrcUrl)
					}
				} else if body != "" {
					markdowns, resultBodyMap, err := markdownUsecase.SearchByBody(body)
					if err != nil {
						log.Fatalf("SearchDatabase: %v\n", err)
					}
					fmt.Println("Search Result")
					for _, m := range markdowns {
						fmt.Printf("id: %d, title: %s, path: %s, url: %s\n", m.Id, m.Title, m.Path, m.SrcUrl)
						//fmt.Println(resultBodyMap[m.Path])
						for _, resultBody := range resultBodyMap[m.Path] {
							fmt.Printf("  %s\n", resultBody)
						}
					}
				}

				return nil
			},
		},
		//zip化する
		{
			Name:  "zip",
			Usage: "zip markdown file",
			Action: func(c *cli.Context) error {
				db, err := models.NewDB()
				if err != nil {
					log.Fatalf("ZipDatabase: %v\n", err)
				}
				markdownRepo := models.NewMarkdownRepo(db)
				markdownUsecase := usecases.NewMarkdownInteractor(markdownRepo)
				// err = markdownUsecase.Zip()
				// if err != nil {
				// 	log.Fatalf("ZipDatabase: %v\n", err)
				// }
				return nil
			},
		},

		//すべてのファイルを削除する　危険
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
