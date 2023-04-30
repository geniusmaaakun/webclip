package handler_test

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"webclip/src/server/controllers/handler"
	"webclip/src/server/models"
	"webclip/src/server/models/rdb"
	"webclip/src/server/usecases"

	"github.com/gorilla/mux"
)

//httptestを使って擬似的にhttpリクエストを送る

func TestListAll(t *testing.T) {
	db, err := models.NewDB("webclip.sql")
	if err != nil {
		t.Fatalf("database error: %v\n", err)
	}
	txRepo := rdb.NewTransactionManager(db)
	markdownRepo := rdb.NewMarkdownRepo()
	markdownUsecase := usecases.NewMarkdownInteractor(txRepo, markdownRepo)
	srv := handler.NewMarkdownHandler(markdownUsecase)

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	type args struct {
		title  string
		path   string
		srcUrl string
	}

	//testcaseの作成
	tests := []struct {
		id        int
		name      string
		args      args
		wantTitle string //インスタンスを返す様にcreateを変更する？
	}{}

	for i := 0; i < 100; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("test%d", i))
		s := fmt.Sprintf("%x", h.Sum(nil))
		//t.Log(s)
		tests = append(tests, struct {
			id        int
			name      string
			args      args
			wantTitle string
		}{i, fmt.Sprintf("normal %d", i), args{fmt.Sprintf("test%s", string(s)), fmt.Sprintf("testdata/%s/README.md", string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s))}, fmt.Sprintf("test%s", string(s))})
		err := markdownUsecase.Create(fmt.Sprintf("test%s", string(s)), fmt.Sprintf("testdata/%s/README.md", string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s)))
		if err != nil {
			t.Errorf("MarkdownInteractor.Delete() error = got %v, s value = %s\n", err, string(s))
		}

	}
	req := httptest.NewRequest("GET", "/api/markdowns", nil)
	got := httptest.NewRecorder()
	srv.ListAll(got, req)

	// Assertion
	// http.Clientなどで受け取ったhttp.Responseを検証するときとほぼ変わらない
	if got.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, got.Code)
	}

	mds, err := markdownUsecase.FindAll()
	if err != nil {
		t.Errorf("MarkdownInteractor.FindAll() error = %v\n", err)
	}
	want, err := json.Marshal(mds)
	if err != nil {
		t.Errorf("MarkdownInteractor.FindAll() error = %v\n", err)
	}

	// Bodyは*bytes.Buffer型なので文字列の比較は少しラク
	if got := got.Body.String(); strings.Compare(string(want), got) != 0 {
		t.Errorf("want %s, but %s", want, got)
	}

	/*
		// http.Responseオブジェクトとしても比較できる。
		if resp := got.Result().Cookies(); resp.ContentLength == 0 {
			t.Errorf("resp.ContentLength was 0")
		}
	*/
}

func TestFindById(t *testing.T) {
	db, err := models.NewDB("webclip.sql")
	if err != nil {
		t.Fatalf("database error: %v\n", err)
	}
	txRepo := rdb.NewTransactionManager(db)
	markdownRepo := rdb.NewMarkdownRepo()
	markdownUsecase := usecases.NewMarkdownInteractor(txRepo, markdownRepo)
	srv := handler.NewMarkdownHandler(markdownUsecase)

	defer t.Cleanup(func() {
		//テスト用のDBを削除
		os.Remove("webclip.sql")
	})

	type args struct {
		title  string
		path   string
		srcUrl string
	}

	//testcaseの作成
	tests := []struct {
		id        int
		name      string
		args      args
		wantTitle string //インスタンスを返す様にcreateを変更する？
	}{}

	for i := 0; i < 100; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("test%d", i))
		s := fmt.Sprintf("%x", h.Sum(nil))
		//t.Log(s)
		tests = append(tests, struct {
			id        int
			name      string
			args      args
			wantTitle string
		}{i, fmt.Sprintf("normal %d", i), args{fmt.Sprintf("test%s", string(s)), fmt.Sprintf("testdata/%s/README.md", string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s))}, fmt.Sprintf("test%s", string(s))})
		err := markdownUsecase.Create(fmt.Sprintf("test%s", string(s)), fmt.Sprintf("testdata/%s/README.md", string(s)), fmt.Sprintf("http://test%s/test%s", string(s), string(s)))
		if err != nil {
			t.Errorf("MarkdownInteractor.Delete() error = got %v, s value = %s\n", err, string(s))
		}

	}

	router := mux.NewRouter()
	router.HandleFunc("/api/markdowns/{id}", srv.GetById).Methods("GET")

	req, _ := http.NewRequest("GET", "http://localhost/api/markdowns/1", nil)
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	// Assertion
	// http.Clientなどで受け取ったhttp.Responseを検証するときとほぼ変わらない
	if res.Code != http.StatusOK {
		t.Errorf("want %d, but %d", http.StatusOK, res.Code)
	}

	want := `{"id":1,"title":"testf6f4061a1bddc1c04d8109b39f581270","content":"f6f4061a1bddc1c04d8109b39f581270test0f6f4061a1bddc1c04d8109b39f581270","path":"testdata/f6f4061a1bddc1c04d8109b39f581270/README.md","src_url":"http://testf6f4061a1bddc1c04d8109b39f581270/testf6f4061a1bddc1c04d8109b39f581270"}`

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("MarkdownInteractor.FindAll() error = %v\n", err)
	}

	// Bodyは*bytes.Buffer型なので文字列の比較は少しラク
	gotData := &handler.MarkdownMemo{}
	err = json.Unmarshal(body, &gotData)
	if err != nil {
		t.Errorf("MarkdownInteractor.FindAll() error = %v\n", err)
	}

	wantData := &handler.MarkdownMemo{}
	err = json.Unmarshal([]byte(want), &wantData)
	if err != nil {
		t.Errorf("MarkdownInteractor.FindAll() error = %v\n", err)
	}

	if gotData.Title != wantData.Title {
		t.Errorf("want %s, but %s", wantData.Title, gotData.Title)
	}

	if gotData.Content != wantData.Content {
		t.Errorf("want %s, but %s", wantData.Content, gotData.Content)
	}

	if gotData.Path != wantData.Path {
		t.Errorf("want %s, but %s", wantData.Path, gotData.Path)
	}

}
