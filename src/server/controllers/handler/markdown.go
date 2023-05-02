package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
	"webclip/src/server/models"
	"webclip/src/server/usecases"

	"github.com/gorilla/mux"
)

type MarkdownMemo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Path      string    `json:"path"`
	SrcUrl    string    `json:"src_url"`
	CreatedAt time.Time `json:"created_at"`
}

type MarkdownMemos []*MarkdownMemo

func (m *MarkdownMemo) ConvertTo() models.MarkdownMemo {
	return models.MarkdownMemo{Id: m.Id, Title: m.Title, Path: m.Path, SrcUrl: m.SrcUrl, CreatedAt: m.CreatedAt}
}

func (m *MarkdownMemo) ConvertFrom(md *models.MarkdownMemo) error {
	//中身を取得して格納
	file, err := os.Open(md.Path)
	if err != nil {
		//独自エラーを返す
		return err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		//独自エラーを返す
		return err
	}

	m.Id = md.Id
	m.Title = md.Title
	m.Content = string(content)
	m.Path = md.Path
	m.SrcUrl = md.SrcUrl
	m.CreatedAt = md.CreatedAt

	return nil
}

type MarkdownHandler struct {
	MarkdownInteractor usecases.MarkdownUsecase
}

func NewMarkdownHandler(i usecases.MarkdownUsecase) *MarkdownHandler {
	return &MarkdownHandler{MarkdownInteractor: i}
}

func (h *MarkdownHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	markdowns, err := h.MarkdownInteractor.FindAll()
	if err != nil {
		//独自エラーを返す
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//jsonで返す
	jsonData, err := json.Marshal(markdowns)
	if err != nil {
		//独自エラーを返す
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		//独自エラーを返す
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

//Idで取得？
/*
IDをキーとしてデータベースからデータを取得すること自体は危険ではありません。IDは一般的に一意であり、データベース内の特定のレコードを効率的に検索・取得するために使用されます。

ただし、IDに関連するセキュリティリスクが存在します。以下は一般的なリスクと対策です。

1. SQLインジェクション: ユーザー入力に基づいてSQLクエリを生成する際に、悪意のあるコードが埋め込まれることがあります。対策として、クエリのパラメータ化やプリペアドステートメントの使用を検討してください。

2. 不適切な公開: セキュアでない状況でIDが公開されると、不正アクセスや情報漏えいのリスクが高まります。対策として、IDの取り扱いに関するポリシーを適切に設定し、必要最低限の範囲でIDを公開してください。

3. ID列挙攻撃: 連続したIDが予測可能である場合、攻撃者が他のIDを推測してアクセスを試みる可能性があります。対策として、ランダムな文字列やUUIDなど、予測しにくいIDを使用してください。

4. アクセス制御: データへのアクセス権限が適切に設定されていない場合、認証されていないユーザーや不正なアクセスによりデータが漏洩するリスクがあります。対策として、アクセス制御リストやロールベースのアクセス制御を適切に設定してください。

これらの対策を適切に実施すれば、IDをキーとしてデータベースからデータを取得することは安全です。
*/
func (h *MarkdownHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	md, err := h.MarkdownInteractor.FindById(idStr)
	if err != nil {
		//独自エラーを返す
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	markdown := &MarkdownMemo{}
	err = markdown.ConvertFrom(md)
	if err != nil {
		//独自エラーを返す
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//fmt.Println(markdown)
	//fmt.Println(md)

	//jsonで返す
	jsonData, err := json.Marshal(markdown)
	if err != nil {
		//独自エラーを返す
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		//独自エラーを返す
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *MarkdownHandler) List(w http.ResponseWriter, r *http.Request) {

}

func (h *MarkdownHandler) ListByTitle(w http.ResponseWriter, r *http.Request) {

}
