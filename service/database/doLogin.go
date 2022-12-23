package database

import (
	"database/sql"
)

func (db *appdbimpl) DoLogin(name Username) (User, error) {

	const query = ` SELECT * from user where username = ?`
	//row := db.c.Query(query, name.USERNAME)
	var us User

	row := db.c.QueryRow(query, name.USERNAME)
	switch err := row.Scan(&us.ID, &us.USERNAME); err {

	case sql.ErrNoRows:
		us.USERNAME = name.USERNAME
		us.ID = -2
		return us, nil
	case nil:
		return us, nil
	default:
		us.ID = -1
		return us, err

	}

}
