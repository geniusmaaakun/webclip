package rdb

import (
	"webclip/src/server/models"
	"webclip/src/server/usecases"
)

type MarkdownManager struct {
}

func NewMarkdownRepo() *MarkdownManager {
	return &MarkdownManager{}
}

func (m *MarkdownManager) Create(tx usecases.Transaction, md *models.MarkdownMemo) error {
	//SQLインジェクションを回避される
	stmt, err := tx.Prepare("INSERT INTO markdown_memo (title, path, src_url, created_at) VALUES (?, ?, ?, ?)")
	_, err = stmt.Exec(md.Title, md.Path, md.SrcUrl, md.CreatedAt)
	//_, err := tx.Exec("INSERT INTO markdown_memo (title, path, src_url, created_at) VALUES (?, ?, ?, ?)", md.Title, md.Path, md.SrcUrl, md.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

//unused
func (m *MarkdownManager) Delete(tx usecases.Transaction, md *models.MarkdownMemo) error {
	stmt, err := tx.Prepare("DELETE FROM markdown_memo WHERE id = ?")
	_, err = stmt.Exec(md.Id)
	if err != nil {
		return err
	}
	return nil
}

func (m *MarkdownManager) DeleteByTitle(tx usecases.Transaction, title string) error {
	stmt, err := tx.Prepare("DELETE FROM markdown_memo WHERE title = ?")
	_, err = stmt.Exec(title)
	if err != nil {
		return err
	}
	return nil
}

//cleanコマンドで削除する
func (m *MarkdownManager) DeleteByPath(tx usecases.Transaction, path string) error {
	stmt, err := tx.Prepare("DELETE FROM markdown_memo WHERE path = ?")
	_, err = stmt.Exec(path)
	if err != nil {
		return err
	}
	return nil
}

//unused
func (m *MarkdownManager) Update(tx usecases.Transaction, md *models.MarkdownMemo) error {
	stmt, err := tx.Prepare("UPDATE markdown_memo SET title = ?, path = ?, src_url = ? WHERE id = ?")
	_, err = stmt.Exec(md.Title, md.Path, md.SrcUrl, md.Id)
	if err != nil {
		return err
	}
	return nil
}

func (m *MarkdownManager) FindById(tx usecases.Transaction, id int) (*models.MarkdownMemo, error) {
	stmt, err := tx.Prepare("SELECT * FROM markdown_memo WHERE id = ?")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(id)
	md := &models.MarkdownMemo{}
	err = row.Scan(&md.Id, &md.Title, &md.Path, &md.SrcUrl, &md.CreatedAt)
	if err != nil {
		return nil, err
	}
	return md, nil
}

func (m *MarkdownManager) FindAll(tx usecases.Transaction) ([]*models.MarkdownMemo, error) {
	stmt, err := tx.Prepare("SELECT * FROM markdown_memo")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	memos := []*models.MarkdownMemo{}
	for rows.Next() {
		memo := &models.MarkdownMemo{}
		err := rows.Scan(&memo.Id, &memo.Title, &memo.Path, &memo.SrcUrl, &memo.CreatedAt)
		if err != nil {
			return nil, err
		}
		memos = append(memos, memo)
	}
	return memos, nil
}

func (m *MarkdownManager) FindByTitle(tx usecases.Transaction, title string) ([]*models.MarkdownMemo, error) {
	// stmt, err := tx.Prepare("SELECT * FROM markdown_memo WHERE title = ?")
	// if err != nil {
	// 	return nil, err
	// }
	// rows, err := stmt.Query(title)
	// if err != nil {
	// 	return nil, err
	// }
	rows, err := tx.Query("SELECT * FROM markdown_memo WHERE title = ?", title)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	memos := []*models.MarkdownMemo{}
	for rows.Next() {
		memo := &models.MarkdownMemo{}
		err := rows.Scan(&memo.Id, &memo.Title, &memo.Path, &memo.SrcUrl, &memo.CreatedAt)
		if err != nil {
			return nil, err
		}
		memos = append(memos, memo)
	}
	return memos, nil
}

//
func (m *MarkdownManager) FindByPath(tx usecases.Transaction, path string) ([]*models.MarkdownMemo, error) {
	stmt, err := tx.Prepare("SELECT * FROM markdown_memo WHERE path = ?")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(path)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	memos := []*models.MarkdownMemo{}
	for rows.Next() {
		memo := &models.MarkdownMemo{}
		err := rows.Scan(&memo.Id, &memo.Title, &memo.Path, &memo.SrcUrl, &memo.CreatedAt)
		if err != nil {
			return nil, err
		}
		memos = append(memos, memo)
	}
	return memos, nil
}

//test用
func (m *MarkdownManager) FindByTitleLastOne(tx usecases.Transaction, title string) (*models.MarkdownMemo, error) {
	stmt, err := tx.Prepare("SELECT * FROM markdown_memo WHERE title = ? ORDER BY id DESC LIMIT 1")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		md := &models.MarkdownMemo{}
		err := rows.Scan(&md.Id, &md.Title, &md.Path, &md.SrcUrl, &md.CreatedAt)
		if err != nil {
			return nil, err
		}
		return md, nil
	}
	return nil, nil
}

//unused
//SrcUrlが一致するものを探す
func (m *MarkdownManager) FindBySrcUrl(tx usecases.Transaction, srcUrl string) (*models.MarkdownMemo, error) {
	stmt, err := tx.Prepare("SELECT * FROM markdown_memo WHERE src_url = ?")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(srcUrl)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		md := &models.MarkdownMemo{}
		err := rows.Scan(&md.Id, &md.Title, &md.Path, &md.SrcUrl, &md.CreatedAt)
		if err != nil {
			return nil, err
		}
		return md, nil
	}
	return nil, nil
}

func (m *MarkdownManager) SearchByTitle(tx usecases.Transaction, title string) ([]*models.MarkdownMemo, error) {
	stmt, err := tx.Prepare("SELECT * FROM markdown_memo WHERE title LIKE ?")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query("%" + title + "%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	memos := []*models.MarkdownMemo{}
	for rows.Next() {
		memo := &models.MarkdownMemo{}
		err := rows.Scan(&memo.Id, &memo.Title, &memo.Path, &memo.SrcUrl, &memo.CreatedAt)
		if err != nil {
			return nil, err
		}
		memos = append(memos, memo)
	}
	return memos, nil
}
