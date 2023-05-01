package actions

import (
	"log"
	"path/filepath"
	"webclip/src/server/models"
	"webclip/src/server/models/rdb"
	"webclip/src/server/usecases"

	"github.com/urfave/cli/v2"
)

func Clean(dbPath string) func(*cli.Context) error {
	return func(c *cli.Context) error {
		db, err := models.NewDB(dbPath)
		if err != nil {
			log.Fatalf("SaveDatabase: %v\n", err)
		}
		txRepo := rdb.NewTransactionManager(db)
		markdownRepo := rdb.NewMarkdownRepo()
		markdownUsecase := usecases.NewMarkdownInteractor(txRepo, markdownRepo)
		err = markdownUsecase.DeleteIfNotExistsByPath()
		if err != nil {
			log.Fatalf("CleanDatabase: %v\n", err)
		}
		return nil
	}
}

func clean(c *cli.Context) error {
	folderPath, err := models.GetDatabasePath()
	if err != nil {
		log.Fatalf("SaveDatabase: %v\n", err)
	}
	db, err := models.NewDB(filepath.Join(folderPath, "webclip.sql"))

	txRepo := rdb.NewTransactionManager(db)
	markdownRepo := rdb.NewMarkdownRepo()
	markdownUsecase := usecases.NewMarkdownInteractor(txRepo, markdownRepo)
	err = markdownUsecase.DeleteIfNotExistsByPath()
	if err != nil {
		log.Fatalf("CleanDatabase: %v\n", err)
	}
	return nil
}
