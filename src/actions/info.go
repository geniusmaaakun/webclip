package actions

import (
	"fmt"
	"log"
	"os"
	"webclip/src/server/models"

	"github.com/urfave/cli/v2"
)

func Info() func(*cli.Context) error {
	return func(c *cli.Context) error {
		//db path
		folderPath, err := models.GetDatabasePath()
		if err != nil {
			log.Fatalf("main: %v\n", err)
		}
		fmt.Printf("db path:	%s\n", folderPath)
		//db size
		info, err := os.Stat(folderPath)
		fmt.Printf("db size:	 %dbyte\n", info.Size())
		return nil
	}
}
