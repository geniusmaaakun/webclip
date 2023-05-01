package actions

import (
	"log"
	"path/filepath"
	"webclip/src/server/models"
	"webclip/src/server/models/rdb"
	"webclip/src/server/usecases"

	"github.com/urfave/cli/v2"
)

func Zip(dbPath string) func(*cli.Context) error {
	return func(c *cli.Context) error {
		db, err := models.NewDB(dbPath)
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
}

func zip(c *cli.Context) error {
	folderPath, err := models.GetDatabasePath()
	if err != nil {
		log.Fatalf("SaveDatabase: %v\n", err)
	}
	db, err := models.NewDB(filepath.Join(folderPath, "webclip.sql"))
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
