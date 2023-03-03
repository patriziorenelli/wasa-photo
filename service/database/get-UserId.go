package database

import (
	"database/sql"
	"errors"
)

// Va bene
func (db *appdbimpl) GetUsreId(userId int, name Username) (UserId, int) {

	var us UserId

	if db.UserExist(userId) == -1 {
		return us, -2
	}

	row := db.c.QueryRow(`SELECT id from user where username = ?`, name.USERNAME)
	err := row.Scan(&us.USERID)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return us, -1
	} else if err != nil {
		return us, -3
	} else {
		return us, 0
	}
}
