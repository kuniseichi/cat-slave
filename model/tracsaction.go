package model

import (
	"cat-slave/model/passage"
)

func Transact(txFunc func() error) (err error) {
	tx := DB.Mysql.Begin()

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

	err = txFunc()
	return err
}


func DoSomething() error {
	return Transact(func () error {
		if _, err := passage.List(); err != nil {
			return err
		}
		if _, err := passage.Get(1); err != nil {
			return err
		}
		return nil
	})
}
