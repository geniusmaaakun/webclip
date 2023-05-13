package controllers

import (
	"database/sql"
	"net/http"
	"webclip/src/server/controllers/handler"
	"webclip/src/server/models/rdb"
	"webclip/src/server/usecases"

	"github.com/gorilla/mux"
)

type Server struct {
	Host string
	Port string

	MarkdownHandler *handler.MarkdownHandler
}

func NewServer(host, port string, db *sql.DB) *Server {
	txRepo := rdb.NewTransactionManager(db)
	markdownRepo := rdb.NewMarkdownRepo()
	markdownUsecase := usecases.NewMarkdownInteractor(txRepo, markdownRepo)
	markdownHandler := handler.NewMarkdownHandler(markdownUsecase)
	return &Server{Host: host, Port: port, MarkdownHandler: markdownHandler}
}

//gorilla/muxを使ってルーティングを設定する
func (s *Server) Run() error {
	router := mux.NewRouter().StrictSlash(true) //末尾/を許可しない

	//サブルーターを切る
	ApiRouter := router.PathPrefix("/api").Subrouter()

	//CORSを許可する
	//react側のAPIアクセスをポート指定不要
	ApiRouter.Use(s.enableCORS)

	//router.HandleFunc("/api/markdowns", s.MarkdownHandler.ListAll).Methods("GET")
	ApiRouter.HandleFunc("/markdowns", s.MarkdownHandler.ListAll).Methods("GET")
	ApiRouter.HandleFunc("/markdowns/{id}", s.MarkdownHandler.GetById).Methods("GET")

	//Reactのhtmlを返す
	fs := http.FileServer(http.Dir("./frontend/build"))
	router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	err := http.ListenAndServe(s.Host+":"+s.Port, router)
	return err
}
