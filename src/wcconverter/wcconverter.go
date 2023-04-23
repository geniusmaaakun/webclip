package wcconverter

import (
	"log"
	"os"
	"path/filepath"
	"webclip/src/server/models"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

type HTMLToMarkdownConverter struct {
	OutputDir string
	FileName  string
	Options   *md.Options
	repo      *models.MarkdownRepo
}

func NewConverter(outputDir string, filename string, op *md.Options) *HTMLToMarkdownConverter {
	db, err := models.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	return &HTMLToMarkdownConverter{
		OutputDir: outputDir,
		FileName:  filename,
		Options:   op,
		repo:      models.NewMarkdownRepo(db),
	}
}

func (c *HTMLToMarkdownConverter) Convert(selection *goquery.Selection) (string, error) {
	converter := md.NewConverter("", true, c.Options)
	markdown := converter.Convert(selection)
	return markdown, nil
}

func (c *HTMLToMarkdownConverter) SaveToFile(mdStr string) error {
	//4 Markdownをファイルに保存save
	file, err := os.Create(filepath.Join(c.OutputDir, c.FileName))
	if err != nil {
		//log.Fatalf("File Create Error: %s\n", err.Error())
		return err
	}
	defer file.Close()
	_, err = file.WriteString(mdStr)
	if err != nil {
		//log.Fatalf("File Write Error: %s\n", err.Error())
		return err
	}
	return nil
}
