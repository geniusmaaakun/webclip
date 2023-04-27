package rdb

import (
	"database/sql"
	"fmt"
	"webclip/src/server/usecases"
)

//unused

type TransactionManager struct {
	db *sql.DB
}

func NewTransactionManager(db *sql.DB) *TransactionManager {
	return &TransactionManager{db}
}

func (tx *TransactionManager) NewTransaction(on bool) (usecases.Transaction, error) {
	if !on {
		return tx.db, nil
	}
	return tx.db.Begin()
}

//トランザクションの処理
func (*TransactionManager) Transaction(CTx usecases.Transaction, txFunc func(*sql.Tx) (interface{}, error)) (data interface{}, err error) {
	tx, isTx := CTx.(*sql.Tx)
	if !isTx {
		return nil, fmt.Errorf("not tx")
	}

	defer func() {
		//panic時
		//p panicの中身のエラーが返る
		if p := recover(); p != nil {
			tx.Rollback()
			err = fmt.Errorf("original=%v", p)

			//panicを起こさせる
			//panic(p)

			//return の前に実行
			//panicではなく、error時
		} else if err != nil {
			rerr := tx.Rollback()
			err = fmt.Errorf("original=%v, rerr=%v", err, rerr)
		} else {
			if rerr := tx.Commit(); rerr != nil {
				err = fmt.Errorf("original=%v, rerr=%v", err, rerr)
			}
		}
	}()

	data, err = txFunc(tx)
	return data, err
}
