package models

import (
	"database/sql"
	"time"
)

type MarkdownRepo struct {
	db *sql.DB
}

func NewMarkdownRepo(db *sql.DB) *MarkdownRepo {
	return &MarkdownRepo{db: db}
}

type MarkdownMemo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Path      string    `json:"path"`
	SrcUrl    string    `json:"src_url"`
	CreatedAt time.Time `json:"created_at"`
}

func NewMarkdownMemo(title, path, srsUrl string) *MarkdownMemo {
	return &MarkdownMemo{Title: title, Path: path, SrcUrl: srsUrl, CreatedAt: time.Now()}
}

func (m *MarkdownRepo) Create(md *MarkdownMemo) error {
	_, err := m.db.Exec("INSERT INTO markdown_memo (title, path, src_url, created_at) VALUES (?, ?, ?, ?)", md.Title, md.Path, md.SrcUrl, md.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (m *MarkdownRepo) Delete(md *MarkdownMemo) error {
	_, err := m.db.Exec("DELETE FROM markdown_memo WHERE id = ?", md.Id)
	if err != nil {
		return err
	}
	return nil
}

func (m *MarkdownRepo) DeleteByTitle(md *MarkdownMemo) error {
	_, err := m.db.Exec("DELETE FROM markdown_memo WHERE title = ?", md.Title)
	if err != nil {
		return err
	}
	return nil
}

func (m *MarkdownRepo) Update(md *MarkdownMemo) error {
	_, err := m.db.Exec("UPDATE markdown_memo SET title = ?, path = ?, src_url = ? WHERE id = ?", md.Title, md.Path, md.SrcUrl, md.Id)
	if err != nil {
		return err
	}
	return nil
}

func (m *MarkdownRepo) FindById(id int) (*MarkdownMemo, error) {
	//prepareをすべてで使う
	stmt, err := m.db.Prepare("SELECT * FROM markdown_memo WHERE id = ?")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(id)
	md := &MarkdownMemo{}
	err = row.Scan(&md.Id, &md.Title, &md.Path, &md.SrcUrl, &md.CreatedAt)
	if err != nil {
		return nil, err
	}
	return md, nil
}

func (m *MarkdownRepo) FindAll() ([]*MarkdownMemo, error) {
	rows, err := m.db.Query("SELECT * FROM markdown_memo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	memos := []*MarkdownMemo{}
	for rows.Next() {
		memo := &MarkdownMemo{}
		err := rows.Scan(&memo.Id, &memo.Title, &memo.Path, &memo.SrcUrl, &memo.CreatedAt)
		if err != nil {
			return nil, err
		}
		memos = append(memos, memo)
	}
	return memos, nil
}

func (m *MarkdownRepo) FindByTitle(title string) ([]*MarkdownMemo, error) {
	rows, err := m.db.Query("SELECT * FROM markdown_memo WHERE title = ?", title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	memos := []*MarkdownMemo{}
	for rows.Next() {
		memo := &MarkdownMemo{}
		err := rows.Scan(&memo.Id, &memo.Title, &memo.Path, &memo.SrcUrl, &memo.CreatedAt)
		if err != nil {
			return nil, err
		}
		memos = append(memos, memo)
	}
	return memos, nil
}

//
func (m *MarkdownRepo) FindByPath(path string) ([]*MarkdownMemo, error) {
	rows, err := m.db.Query("SELECT * FROM markdown_memo WHERE path = ?", path)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	memos := []*MarkdownMemo{}
	for rows.Next() {
		memo := &MarkdownMemo{}
		err := rows.Scan(&memo.Id, &memo.Title, &memo.Path, &memo.SrcUrl, &memo.CreatedAt)
		if err != nil {
			return nil, err
		}
		memos = append(memos, memo)
	}
	return memos, nil
}

//test用
func (m MarkdownRepo) FindByTitleLastOne(title string) (*MarkdownMemo, error) {
	rows, err := m.db.Query("SELECT * FROM markdown_memo WHERE title = ? ORDER BY id DESC LIMIT 1", title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		md := &MarkdownMemo{}
		err := rows.Scan(&md.Id, &md.Title, &md.Path, &md.SrcUrl, &md.CreatedAt)
		if err != nil {
			return nil, err
		}
		return md, nil
	}
	return nil, nil
}
