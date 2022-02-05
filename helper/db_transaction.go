package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		err2 := tx.Rollback()
		PanicHandler(err2)
		panic(err)
	} else {
		err2 := tx.Commit()
		PanicHandler(err2)
	}
}
