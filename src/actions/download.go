package actions

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"webclip/src/server/models"
	"webclip/src/server/models/rdb"
	"webclip/src/wcconverter"
	"webclip/src/wcdownloader"

	"github.com/urfave/cli/v2"
)

func Download(c *cli.Context) error {
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
		//usecases
		txRepo := rdb.NewTransactionManager(db)
		markdownRepo := rdb.NewMarkdownRepo()
		absPath, err := filepath.Abs(filepath.Join(outdir, "README.md"))
		if err != nil {
			log.Fatalf("SaveDatabase: %v\n", err)
		}
		mdData := models.NewMarkdownMemo(outdir, absPath, url)
		tx, err := txRepo.NewTransaction(false)
		err = markdownRepo.Create(tx, mdData)
		if err != nil {
			existMd, err := markdownRepo.FindBySrcUrl(tx, url)
			if err != nil {
				log.Fatalf("SaveDatabase: %v\n", err)
			}
			if existMd != nil {
				log.Printf("already exist: %s", existMd.Path)
				return nil
			}
		}
	}
	return nil
}
