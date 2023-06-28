package actions

import (
	"log"
	"os"
	"webclip/src/server/models"

	"github.com/urfave/cli/v2"
)

func ResetDb() func(*cli.Context) error {
	return func(c *cli.Context) error {
		//db path delete
		folderPath, err := models.GetDatabasePath()
		if err != nil {
			log.Fatalf("main: %v\n", err)
		}
		err = os.RemoveAll(folderPath)
		if err != nil {
			log.Fatalf("main: %v\n", err)
		}
		return nil
	}
}
