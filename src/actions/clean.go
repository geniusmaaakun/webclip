package actions

import (
	"log"
	"webclip/src/server/models"
	"webclip/src/server/models/rdb"
	"webclip/src/server/usecases"

	"github.com/urfave/cli/v2"
)

func Clean(c *cli.Context) error {
	db, err := models.NewDB()
	if err != nil {
		log.Fatalf("CleanDatabase: %v\n", err)
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
