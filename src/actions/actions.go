package actions

import "github.com/urfave/cli/v2"

func NewWebClip(dbPath string) *cli.App {
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
				Name:    "save",
				Aliases: []string{"sv"},
				Usage:   "Save to DB",
			},
		},
		//WebClip
		Action: Download(dbPath),

		Commands: []*cli.Command{
			//sub command : webclip server
			{
				Name:   "server",
				Usage:  "Start Web Server",
				Action: Server(dbPath),
			},
			//sub command : webclip clean
			{
				Name:   "clean",
				Usage:  "clean database if file is not exist",
				Action: Clean(dbPath),
			},
			//sub command : webclip search
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
				Action: Search(dbPath),
			},
			//zip化する
			//sub command : webclip zip
			{
				Name:  "zip",
				Usage: "zip markdown file",
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
				Action: Zip(dbPath),
			},
			{
				Name:   "info",
				Usage:  "webclip information",
				Action: Info(),
			},
			{
				Name:   "resetdb",
				Usage:  "delete db file",
				Action: ResetDb(),
			},
		},
	}

	return app
}
