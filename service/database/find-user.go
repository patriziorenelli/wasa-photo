package database

import "fmt"

func (db *appdbimpl) FindUsername(username string) (int, error) {

	// Ricerca un user in base al suo username
	const query = `SELECT id FROM user WHERE username = ?`

	rows, err := db.c.Query(query, username)
	if err != nil {
		return -1, err
	}

	fmt.Print(rows)

	var us User
	err = rows.Scan(&us.ID, us.USERNAME)
	if err != nil {
		return -1, err
	}

	return us.ID, nil

}
