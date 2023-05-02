package models

import (
	"time"
)

type MarkdownMemo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Path      string    `json:"path"`
	SrcUrl    string    `json:"src_url"`
	CreatedAt time.Time `json:"created_at"`
}

type MarkdownMemos []*MarkdownMemo

func NewMarkdownMemo(title, path, srsUrl string) *MarkdownMemo {
	return &MarkdownMemo{Title: title, Path: path, SrcUrl: srsUrl, CreatedAt: time.Now()}
}
