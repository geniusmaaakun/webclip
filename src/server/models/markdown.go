package models

import "database/sql"

type MarkdownRepo struct {
	Db *sql.DB
}

func NewMarkdownRepo(db *sql.DB) *MarkdownRepo {
	return &MarkdownRepo{Db: db}
}

type MarkdownMemo struct {
	Id     int
	Title  string
	Path   string
	SrcUrl string
}

func NewMarkdownMemo(title, path, srsUrl string) *MarkdownMemo {
	return &MarkdownMemo{Title: path, Path: path, SrcUrl: srsUrl}
}

func (m *MarkdownRepo) Create(md *MarkdownMemo) error {
	_, err := m.Db.Exec("INSERT INTO markdown_memo (title, path, src_url) VALUES (?, ?, ?)", md.Title, md.Path, md.SrcUrl)
	if err != nil {
		return err
	}
	return nil
}

func (m *MarkdownRepo) Delete(md *MarkdownMemo) error {
	_, err := m.Db.Exec("DELETE FROM markdown_memo WHERE id = ?", md.Id)
	if err != nil {
		return err
	}
	return nil
}

func (m *MarkdownRepo) Update(md *MarkdownMemo) error {
	_, err := m.Db.Exec("UPDATE markdown_memo SET title = ?, path = ?, src_url = ? WHERE id = ?", md.Title, md.Path, md.SrcUrl, md.Id)
	if err != nil {
		return err
	}
	return nil
}

func (m *MarkdownRepo) FindById(id int) (*MarkdownMemo, error) {
	row := m.Db.QueryRow("SELECT * FROM markdown_memo WHERE id = ?", id)
	md := &MarkdownMemo{}
	err := row.Scan(&md.Id, &md.Title, &md.Path, &md.SrcUrl)
	if err != nil {
		return nil, err
	}
	return md, nil
}

func (m *MarkdownRepo) FindAll() ([]*MarkdownMemo, error) {
	rows, err := m.Db.Query("SELECT * FROM markdown_memo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	memos := []*MarkdownMemo{}
	for rows.Next() {
		memo := &MarkdownMemo{}
		err := rows.Scan(&memo.Id, &memo.Title, &memo.Path, &memo.SrcUrl)
		if err != nil {
			return nil, err
		}
		memos = append(memos, memo)
	}
	return memos, nil
}

func (m *MarkdownRepo) FindByTitle(title string) ([]*MarkdownMemo, error) {
	rows, err := m.Db.Query("SELECT * FROM markdown_memo WHERE title = ?", title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	memos := []*MarkdownMemo{}
	for rows.Next() {
		memo := &MarkdownMemo{}
		err := rows.Scan(&memo.Id, &memo.Title, &memo.Path, &memo.SrcUrl)
		if err != nil {
			return nil, err
		}
		memos = append(memos, memo)
	}
	return memos, nil
}

//
func (m *MarkdownRepo) FindByPath(path string) ([]*MarkdownMemo, error) {
	rows, err := m.Db.Query("SELECT * FROM markdown_memo WHERE path = ?", path)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	memos := []*MarkdownMemo{}
	for rows.Next() {
		memo := &MarkdownMemo{}
		err := rows.Scan(&memo.Id, &memo.Title, &memo.Path, &memo.SrcUrl)
		if err != nil {
			return nil, err
		}
		memos = append(memos, memo)
	}
	return memos, nil
}

//test用
func (m MarkdownRepo) FindByTitleLastOne(title string) (*MarkdownMemo, error) {
	rows, err := m.Db.Query("SELECT * FROM markdown_memo WHERE title = ? ORDER BY id DESC LIMIT 1", title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		md := &MarkdownMemo{}
		err := rows.Scan(&md.Id, &md.Title, &md.Path, &md.SrcUrl)
		if err != nil {
			return nil, err
		}
		return md, nil
	}
	return nil, nil
}
