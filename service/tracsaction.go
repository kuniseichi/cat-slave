package service

import "cat-slave/model"

func Transact(txFunc func() error) (err error) {
	tx := model.DB.Mysql.Begin()

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	return txFunc()
}


