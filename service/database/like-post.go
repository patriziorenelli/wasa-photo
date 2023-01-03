package database

import (
	"database/sql"
	"errors"
)

// VA BENE
func (db *appdbimpl) LikePost(userId int, photoId int) int {

	// Variabile di tipo User usato per il check
	var us User

	// Controllo che l'utente indicato esista
	row := db.c.QueryRow(`SELECT * from user where id = ?`, userId)
	err := row.Scan(&us.ID, &us.USERNAME)

	if errors.Is(err, sql.ErrNoRows) {
		return -1
	}

	// Variabile di tipo Post usata per il check
	var post Post
	row = db.c.QueryRow(`SELECT * from post where id = ?`, photoId)
	err = row.Scan(&post.ID, &post.USERID, &post.PHOTO)

	if errors.Is(err, sql.ErrNoRows) {
		return -2
	}

	// Variabile di tipo Ban usata per i check
	var ban Ban

	// Controllo che l'utente che si vuole seguire non abbia bloccato l'utente che lo vuole seguire
	row = db.c.QueryRow(`SELECT * from ban where uid = ? and uid2 = ?`, userId, post.USERID)
	err = row.Scan(&ban.UID1, &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows) {
		return -3
	}
	// Controllo che l'utente non abbia bloccato chi vuole seguire
	row = db.c.QueryRow(`SELECT * from ban where uid = ? and uid2 = ?`, post.USERID, userId)
	err = row.Scan(&ban.UID1, &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows) {
		return -4
	}

	// Aggiungo il like al post
	_, err = db.c.Exec(`INSERT INTO like VALUES (? , ?)`, photoId, userId)

	// Caso in cui ci sia già quel like
	if err != nil && (err.Error()) == "UNIQUE constraint failed: like.phid, like.uid" {
		return -5
	} else if err != nil {
		return -6
	}

	return 0

}
