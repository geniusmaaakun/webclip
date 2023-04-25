package controllers

import (
	"database/sql"
	"net/http"
	"webclip/src/server/controllers/handler"
	"webclip/src/server/models"

	"github.com/gorilla/mux"
)

type Server struct {
	Host string
	Port string

	MarkdownHandler *handler.MarkdownHandler
}

func NewServer(host, port string, db *sql.DB) *Server {
	markdownHandler := handler.NewMarkdownHandler(models.NewMarkdownRepo(db))
	return &Server{Host: host, Port: port, MarkdownHandler: markdownHandler}
}

//gorilla/muxを使ってルーティングを設定する
func (s *Server) Run() error {
	router := mux.NewRouter().StrictSlash(true) //末尾/を許可しない
	router.Use(s.enableCORS)
	//Reactのhtmlを返す
	//router.HandleFunc("/api/markdowns", s.MarkdownHandler.ListAll).Methods("GET")
	router.HandleFunc("/api/markdowns", s.MarkdownHandler.ListAll).Methods("GET")

	err := http.ListenAndServe(s.Host+":"+s.Port, router)
	return err
}
