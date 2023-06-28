package actions

import (
	"fmt"
	"log"
	"path/filepath"
	"webclip/src/server/models"
	"webclip/src/server/models/rdb"
	"webclip/src/server/usecases"

	"github.com/urfave/cli/v2"
)

func Search(dbPath string) func(*cli.Context) error {
	return func(c *cli.Context) error {
		title := c.String("title")
		body := c.String("body")
		if title == "" && body == "" {
			log.Fatal("title or body is required")
			return nil
		}

		db, err := models.NewDB(dbPath)
		if err != nil {
			log.Fatalf("SaveDatabase: %v\n", err)
		}
		markdownRepo := rdb.NewMarkdownRepo()
		txRepo := rdb.NewTransactionManager(db)
		markdownUsecase := usecases.NewMarkdownInteractor(txRepo, markdownRepo)

		//find . | xargs grep -n hogehoge
		//bodyの行数も取得？
		if title != "" && body != "" {
			markdownsByTitle, err := markdownUsecase.SearchByTitle(title)
			if err != nil {
				log.Fatalf("SearchDatabase: %v\n", err)
			}
			markdownsByBody, resultBodyMap, err := markdownUsecase.SearchByBody(body)
			if err != nil {
				log.Fatalf("SearchDatabase: %v\n", err)
			}

			mdPathMap := map[string]*models.MarkdownMemo{}

			for _, m := range markdownsByTitle {
				mdPathMap[m.Path] = m
			}
			for _, m := range markdownsByBody {
				mdPathMap[m.Path] = m
			}

			fmt.Println("Search Result")
			for _, m := range mdPathMap {
				fmt.Printf("id: %d, title: %s, path: %s, url: %s\n", m.Id, m.Title, m.Path, m.SrcUrl)
				for _, resultBody := range resultBodyMap[m.Path] {
					fmt.Printf("  %s\n", resultBody)
				}
			}

		} else if title != "" {
			markdowns, err := markdownUsecase.SearchByTitle(title)
			if err != nil {
				log.Fatalf("SearchDatabase: %v\n", err)
			}
			fmt.Println("Search Result")
			for _, m := range markdowns {
				fmt.Printf("id: %d, title: %s, path: %s, url: %s\n", m.Id, m.Title, m.Path, m.SrcUrl)
			}
		} else if body != "" {
			markdowns, resultBodyMap, err := markdownUsecase.SearchByBody(body)
			if err != nil {
				log.Fatalf("SearchDatabase: %v\n", err)
			}
			fmt.Println("Search Result")
			for _, m := range markdowns {
				fmt.Printf("id: %d, title: %s, path: %s, url: %s\n", m.Id, m.Title, m.Path, m.SrcUrl)
				//fmt.Println(resultBodyMap[m.Path])
				for _, resultBody := range resultBodyMap[m.Path] {
					fmt.Printf("  %s\n", resultBody)
				}
			}
		}
		// -a all flagを設定する?
		//else {
		// 	markdowns, err := markdownUsecase.FindAll()
		// 	if err != nil {
		// 		log.Fatalf("SearchDatabase: %v\n", err)
		// 	}
		// 	fmt.Println("Search Result")
		// 	for _, m := range markdowns {
		// 		fmt.Printf("id: %d, title: %s, path: %s, url: %s\n", m.Id, m.Title, m.Path, m.SrcUrl)
		//	}
		//}

		return nil
	}
}

func search(c *cli.Context) error {
	title := c.String("title")
	body := c.String("body")
	if title == "" && body == "" {
		log.Fatal("title or body is required")
		return nil
	}

	folderPath, err := models.GetDatabasePath()
	if err != nil {
		log.Fatalf("SaveDatabase: %v\n", err)
	}
	db, err := models.NewDB(filepath.Join(folderPath, "webclip.sql"))
	markdownRepo := rdb.NewMarkdownRepo()
	txRepo := rdb.NewTransactionManager(db)
	markdownUsecase := usecases.NewMarkdownInteractor(txRepo, markdownRepo)

	//find . | xargs grep -n hogehoge
	//bodyの行数も取得？
	if title != "" && body != "" {
		markdownsByTitle, err := markdownUsecase.SearchByTitle(title)
		if err != nil {
			log.Fatalf("SearchDatabase: %v\n", err)
		}
		markdownsByBody, resultBodyMap, err := markdownUsecase.SearchByBody(body)
		if err != nil {
			log.Fatalf("SearchDatabase: %v\n", err)
		}

		mdPathMap := map[string]*models.MarkdownMemo{}

		for _, m := range markdownsByTitle {
			mdPathMap[m.Path] = m
		}
		for _, m := range markdownsByBody {
			mdPathMap[m.Path] = m
		}

		fmt.Println("Search Result")
		for _, m := range mdPathMap {
			fmt.Printf("id: %d, title: %s, path: %s, url: %s\n", m.Id, m.Title, m.Path, m.SrcUrl)
			for _, resultBody := range resultBodyMap[m.Path] {
				fmt.Printf("  %s\n", resultBody)
			}
		}

	} else if title != "" {
		markdowns, err := markdownUsecase.SearchByTitle(title)
		if err != nil {
			log.Fatalf("SearchDatabase: %v\n", err)
		}
		fmt.Println("Search Result")
		for _, m := range markdowns {
			fmt.Printf("id: %d, title: %s, path: %s, url: %s\n", m.Id, m.Title, m.Path, m.SrcUrl)
		}
	} else if body != "" {
		markdowns, resultBodyMap, err := markdownUsecase.SearchByBody(body)
		if err != nil {
			log.Fatalf("SearchDatabase: %v\n", err)
		}
		fmt.Println("Search Result")
		for _, m := range markdowns {
			fmt.Printf("id: %d, title: %s, path: %s, url: %s\n", m.Id, m.Title, m.Path, m.SrcUrl)
			//fmt.Println(resultBodyMap[m.Path])
			for _, resultBody := range resultBodyMap[m.Path] {
				fmt.Printf("  %s\n", resultBody)
			}
		}
	}

	return nil
}
