package actions

import (
	"log"
	"webclip/src/server/models"
	"webclip/src/server/models/rdb"
	"webclip/src/server/usecases"

	"github.com/urfave/cli/v2"
)

func Zip(c *cli.Context) error {
	db, err := models.NewDB()
	if err != nil {
		log.Fatalf("CreateZip: %v\n", err)
	}
	txRepo := rdb.NewTransactionManager(db)
	markdownRepo := rdb.NewMarkdownRepo()
	markdownUsecase := usecases.NewMarkdownInteractor(txRepo, markdownRepo)

	title := c.String("title")
	body := c.String("body")

	files := []*models.MarkdownMemo{}

	if title != "" && body != "" {
		markdownsByTitle, err := markdownUsecase.SearchByTitle(title)
		if err != nil {
			log.Fatalf("CreateZip: %v\n", err)
		}
		markdownsByBody, _, err := markdownUsecase.SearchByBody(body)
		if err != nil {
			log.Fatalf("CreateZip: %v\n", err)
		}

		mdPathMap := map[string]*models.MarkdownMemo{}

		for _, m := range markdownsByTitle {
			mdPathMap[m.Path] = m
		}
		for _, m := range markdownsByBody {
			mdPathMap[m.Path] = m
		}

		for _, m := range mdPathMap {
			files = append(files, m)
		}
	} else if title != "" {
		markdowns, err := markdownUsecase.SearchByTitle(title)
		if err != nil {
			log.Fatalf("CreateZip: %v\n", err)
		}
		files = markdowns
	} else if body != "" {
		markdowns, _, err := markdownUsecase.SearchByBody(body)
		if err != nil {
			log.Fatalf("CreateZip: %v\n", err)
		}
		files = markdowns
	} else {
		markdowns, err := markdownUsecase.FindAll()
		if err != nil {
			log.Fatalf("CreateZip: %v\n", err)
		}
		files = markdowns
	}

	err = markdownUsecase.CreateZipFile(files)
	if err != nil {
		log.Fatalf("CreateZip: %v\n", err)
	}
	return nil
}