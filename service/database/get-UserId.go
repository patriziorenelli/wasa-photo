package database

import (
	"database/sql"
	"errors"
)

// DA FARE
func (db *appdbimpl) GetUsreId(userId int, name Username) (UserId, int) {

	var us UserId
	// username associato a qualcuno?
	// controllo parsing risultato

	if db.UserExist(userId) == -1 {
		return us, -2
	}

	row := db.c.QueryRow(`SELECT id from user where username = ?`, name.USERNAME)
	err := row.Scan(&us.USERID)

	if errors.Is(err, sql.ErrNoRows) {
		return us, -1
	} else if err != nil {
		return us, -3
	} else {
		return us, 0
	}
}
