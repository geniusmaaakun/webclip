package handler

import (
	"encoding/json"
	"net/http"
	"webclip/src/server/models"
)

type MarkdownHandler struct {
	MarkdownRepo *models.MarkdownRepo
}

func NewMarkdownHandler(repo *models.MarkdownRepo) *MarkdownHandler {
	return &MarkdownHandler{MarkdownRepo: repo}
}

func (h *MarkdownHandler) ListAll(w http.ResponseWriter, r *http.Request) {
	markdowns, err := h.MarkdownRepo.FindAll()
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
	w.Write(jsonData)

}

func (h *MarkdownHandler) List(w http.ResponseWriter, r *http.Request) {

}

func (h *MarkdownHandler) ListByTitle(w http.ResponseWriter, r *http.Request) {

}
