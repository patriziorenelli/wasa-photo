package database

import (
	"database/sql"
	"errors"
)

// VA BENE
func (db *appdbimpl) SetMyUserName(userId int, newUsername string) int {

	var us User

	// Controllo che l'utente indicato esista
	row := db.c.QueryRow(`SELECT * from user where id = ?`, userId)
	err := row.Scan(&us.ID, &us.USERNAME)

	if errors.Is(err, sql.ErrNoRows) {
		return -1
	}

	// Controllo che l'username selezionato non sia già utilizzato
	row = db.c.QueryRow(`SELECT * from user where username = ?`, newUsername)
	err = row.Scan(&us.ID, &us.USERNAME)

	if !errors.Is(err, sql.ErrNoRows) {
		return -2
	}

	// Aggiorno il nickname
	_, err = db.c.Exec(`UPDATE user SET username = ? WHERE id = ?`, newUsername, userId)
	// In caso di errore nel cambio username
	if err != nil {
		return -3
	} else {
		return 0
	}

}
