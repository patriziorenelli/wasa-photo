package database

import (
	"database/sql"
	"errors"
)

// Va bene
func (db *appdbimpl) BanUser(userId int, banId int) (int, Username) {

	var username Username
	var us User

	// Controllo che l'utente indicato esista
	row := db.c.QueryRow(`SELECT * from user where id = ?`, userId)
	err := row.Scan(&us.ID, &us.USERNAME)

	if errors.Is(err, sql.ErrNoRows) {
		return -1, username
	}

	// Controllo che l'utente che si vuole seguire esista
	row = db.c.QueryRow(`SELECT * from user where id = ?`, banId)
	err = row.Scan(&us.ID, &us.USERNAME)

	if errors.Is(err, sql.ErrNoRows) {
		return -2, username
	}

	username.USERNAME = us.USERNAME

	// Variabile di tipo Ban usata per i check
	var ban Ban

	// Controllo che l'utente che si vuole bannare non abbia bloccato l'utente che lo vuole bloccare
	row = db.c.QueryRow(`SELECT * from ban where uid = ? and uid2 = ?`, banId, userId)
	err = row.Scan(&ban.UID1, &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows) {
		return -3, username
	}

	// Controllo che l'utente non abbia già bloccato chi vuole bloccare
	row = db.c.QueryRow(`SELECT uid from ban where uid = ? and uid2 = ?`, userId, banId)
	err = row.Scan(&ban.UID1, &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows) {
		return -4, username
	}

	// 	BISOGNA FARE IL BAN DELL'UTENTE  E SMETTERE DI SEGUIRE L'UTENTE banId E VERIFICARNE IL BUON ESITO

	_, err = db.c.Exec(`INSERT INTO ban VALUES (? , ?); DELETE FROM follow WHERE uid = ? AND uid2 = ?`, userId, banId, userId, banId)

	// Caso in cui già si ha bannato l'user
	if err != nil && (err.Error()) == "UNIQUE constraint failed: ban.uid, ban.uid2" {
		return -4, username
	} else if err != nil {
		return -6, username
	}

	return 0, username

}
