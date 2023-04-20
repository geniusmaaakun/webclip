package controllers

import (
	"database/sql"
	"webclip/src/server/models"
)

type Server struct {
	Host         string
	Port         string
	MarkdownRepo *models.MarkdownRepo
	Memos        []models.MarkdownMemo
}

func NewServer(host, port string, db *sql.DB) *Server {
	return &Server{Host: host, Port: port, MarkdownRepo: models.NewMarkdownRepo(db)}
}

//gorilla/muxを使ってルーティングを設定する
func (s *Server) Run() error {

	return nil
}
