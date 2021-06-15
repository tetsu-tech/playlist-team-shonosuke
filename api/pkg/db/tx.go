package db

import (
	"database/sql"
)

func Tx(db *sql.DB, f func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			// nolint:errcheck
			tx.Rollback()
			panic(p)
		} else if err != nil {
			// nolint:errcheck
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = f(tx)
	return err
}
