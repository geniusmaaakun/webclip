package actions

import (
	"fmt"
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
		fmt.Printf("delete?: yes/no: ")
		var f string
		fmt.Scanf("%s", &f)
		if f == "yes" {
			err = os.RemoveAll(folderPath)
			if err != nil {
				log.Fatalf("main: %v\n", err)
			}
			fmt.Printf("deletedb: %s\n", folderPath)
		}
		return nil
	}
}
