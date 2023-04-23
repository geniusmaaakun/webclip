package models_test

import (
	"os"
	"testing"
	"webclip/src/server/controllers"
	"webclip/src/server/models"
)

/*
//前処理、後処理
func TestMain(m *testing.M) {
	//前処理

	//テスト実行
	m.Run()
	//後処理
	teardown := func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	}
	defer teardown()
}
*/

func TestCreateMd(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}
	srv := controllers.NewServer("localhost", "8080", db)

	type args struct {
		title  string
		path   string
		srcURL string
	}
	//testcase
	tests := []struct {
		name string
		args args
		want *models.MarkdownMemo
	}{
		{name: "name", args: args{"test", "test", "test"}, want: &models.MarkdownMemo{Title: "test", Path: "test", SrcUrl: "test"}},
		{name: "name2", args: args{"test2", "test2", "test2"}, want: &models.MarkdownMemo{Title: "test2", Path: "test2", SrcUrl: "test2"}},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			md := models.NewMarkdownMemo(tt.args.title, tt.args.path, tt.args.srcURL)
			err = srv.MarkdownRepo.Create(md)
			if err != nil {
				t.Fatal(err)
			}
			got, err := srv.MarkdownRepo.FindByTitleLastOne(md.Title)
			if err != nil {
				t.Fatal(err)
			}
			if got.Title != tt.want.Title {
				t.Errorf("CreateMd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteMd(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}
	srv := controllers.NewServer("localhost", "8080", db)

	type args struct {
		title  string
		path   string
		srcURL string
	}
	//testcase
	tests := []struct {
		name string
		args args
		want *models.MarkdownMemo
	}{
		{name: "name", args: args{"test", "test", "test"}, want: &models.MarkdownMemo{Title: "test", Path: "test", SrcUrl: "test"}},
		{name: "name2", args: args{"test2", "test2", "test2"}, want: &models.MarkdownMemo{Title: "test2", Path: "test2", SrcUrl: "test2"}},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			md := models.NewMarkdownMemo(tt.args.title, tt.args.path, tt.args.srcURL)
			err = srv.MarkdownRepo.Create(md)
			if err != nil {
				t.Fatal(err)
			}
			err = srv.MarkdownRepo.DeleteByTitle(md)
			if err != nil {
				t.Fatal(err)
			}
			got, err := srv.MarkdownRepo.FindByTitleLastOne(md.Title)
			if err != nil {
				t.Fatal(err)
			}
			if got != nil {
				t.Errorf("DeleteMd() = %v, want %v", got, "NULL")
			}
		})
	}

}

func TestUpdateMd(t *testing.T) {
	db, err := models.NewDB()
	if err != nil {
		return
	}
	srv := controllers.NewServer("localhost", "8080", db)

	type args struct {
		title  string
		path   string
		srcURL string
	}
	//testcase
	tests := []struct {
		name string
		args args
		want *models.MarkdownMemo
	}{
		{name: "name", args: args{"test", "test", "test"}, want: &models.MarkdownMemo{Title: "test", Path: "test", SrcUrl: "test"}},
		{name: "name2", args: args{"test2", "test2", "test2"}, want: &models.MarkdownMemo{Title: "test2", Path: "test2", SrcUrl: "test2"}},
	}

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			md := models.NewMarkdownMemo(tt.args.title, tt.args.path, tt.args.srcURL)
			err = srv.MarkdownRepo.Create(md)
			if err != nil {
				t.Fatal(err)
			}

			updateMd, err := srv.MarkdownRepo.FindByTitleLastOne(md.Title)
			if err != nil {
				t.Fatal(err)
			}
			updateMd.Title = "testupdate"
			updateMd.SrcUrl = "srcupdate"
			updateMd.Path = "pathupdate"

			err = srv.MarkdownRepo.Update(updateMd)

			got, err := srv.MarkdownRepo.FindByTitleLastOne(updateMd.Title)
			if err != nil {
				t.Fatal(err)
			}

			if got.Title != updateMd.Title {
				t.Errorf("DeleteMd() = %v, want %v", got.Title, updateMd.Title)
			}

			if got.Path != updateMd.Path {
				t.Errorf("DeleteMd() = %v, want %v", got.Title, updateMd.Title)
			}

			if got.SrcUrl != updateMd.SrcUrl {
				t.Errorf("DeleteMd() = %v, want %v", got.Title, updateMd.Title)
			}
		})
	}
}

func TestFindById(t *testing.T) {

}

func TestFindAll(t *testing.T) {

}

func TestFindByTitle(t *testing.T) {

}
