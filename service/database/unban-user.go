package database


import (
	"database/sql"
	"errors"
)

// VA BENE 
func (db *appdbimpl) UnBanUser(userId int, unbanId int) (int, Username) {

	var username Username
	var us User

	// Controllo che l'utente indicato esista
	row := db.c.QueryRow(`SELECT * from user where id = ?`, userId)
	err := row.Scan(&us.ID, &us.USERNAME)

	if errors.Is(err, sql.ErrNoRows) {
		return -1, username
	}

	// Controllo che l'utente a cui si voglia togliere il ban esista
	row = db.c.QueryRow(`SELECT * from user where id = ?`, unbanId)
	err = row.Scan(&us.ID, &us.USERNAME)

	if errors.Is(err, sql.ErrNoRows) {
		return -2, username
	}

	username.USERNAME = us.USERNAME

	// Variabile di tipo Ban usata per i check
	var ban Ban
	// Controllo che l'utente a cui si voglia togliere il ban non abbia bloccato l'utente
	row = db.c.QueryRow(`SELECT * from ban where uid = ? and uid2 = ?`, unbanId, userId)
	err = row.Scan(&ban.UID1, &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows) {
		return -3, username
	}

	// Controllo che l'utente abbia veramente bannato l'altro user 
	row = db.c.QueryRow(`SELECT uid from ban where uid = ? and uid2 = ?`, userId, unbanId)
	err = row.Scan(&ban.UID1, &ban.UID2)

	if errors.Is(err, sql.ErrNoRows) {
		return -4, username
	}



	// Se l'utente aveva bannato l'altro user allora lo sblocco
	_, err = db.c.Exec(`DELETE FROM ban WHERE uid = ? AND uid2 = ?`, userId, unbanId)
	if err != nil {
		return -5, username
	}else{
		return 0, username
	}

}
