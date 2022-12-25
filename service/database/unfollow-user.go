package database

import (
	"database/sql"
	"errors"
	
)

func (db *appdbimpl) UnfollowUser(userId int, unfollowId int) (int, Username) {

	var username Username
	var us User

	// Controllo che l'utente indicato esista
	row := db.c.QueryRow(`SELECT * from user where id = ?`, userId)
	err := row.Scan(&us.ID, &us.USERNAME)

	if errors.Is(err, sql.ErrNoRows) {
		return -1, username
	}

	// Controllo che l'utente che si vuole seguire esista
	row = db.c.QueryRow(`SELECT * from user where id = ?`, unfollowId)
	err = row.Scan(&us.ID, &us.USERNAME)

	if errors.Is(err, sql.ErrNoRows) {
		return -2, username
	}

	username.USERNAME = us.USERNAME

	// Variabile di tipo Ban usata per i check
	var ban Ban
	// Controllo che l'utente che si vuole seguire non abbia bloccato l'utente che lo vuole seguire
	row = db.c.QueryRow(`SELECT * from ban where uid = ? and uid2 = ?`, unfollowId, userId)
	err = row.Scan(&ban.UID1, &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows) {
		return -3, username
	}

	// Controllo che l'utente non abbia bloccato chi vuole seguire
	row = db.c.QueryRow(`SELECT uid from ban where uid = ? and uid2 = ?`, userId, unfollowId)
	err = row.Scan(&ban.UID1, &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows) {
		return -4, username
	}

	
	var follow Follow

	// -------------------------------------------------


	// Verifico se userId segue
	row = db.c.QueryRow(`SELECT * FROM follow WHERE uid = ? AND uid2 = ?`, userId, unfollowId)
	err = row.Scan(&follow.UID1, &follow.UID2)

	if err != nil && errors.Is(err, sql.ErrNoRows){
		return -5, username
	}else if  err != nil {
		return -6, username
	}


	_, err = db.c.Exec(`DELETE FROM follow WHERE uid = ? AND uid2 = ?`, userId, unfollowId)
	if err != nil {
		return -6, username
	}

	// --------------------------------------------




	return 0, username

}
