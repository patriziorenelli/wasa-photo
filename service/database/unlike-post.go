package database

import (
	"database/sql"
	"errors"
)

// VA BENE
func (db *appdbimpl) UnlikePost(userId int, photoId int) int {

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

	// Controllo che l'utente che ha postato il post a cui si vuole togliere il like non abbia bloccato l'utente
	row = db.c.QueryRow(`SELECT * from ban where uid = ? and uid2 = ?`, userId, post.USERID)
	err = row.Scan(&ban.UID1, &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows) {
		return -3
	}
	// Controllo che l'utente non abbia bloccato l'utente a cui si vuole togliere il mi piace
	row = db.c.QueryRow(`SELECT * from ban where uid = ? and uid2 = ?`, post.USERID, userId)
	err = row.Scan(&ban.UID1, &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows) {
		return -4
	}

	// Controllo che l'utente abbia messo mi piace al post
	var like Like

	row = db.c.QueryRow(`SELECT * FROM like WHERE phid = ? AND uid = ?`, photoId, userId)
	err = row.Scan(&like.PHID, &like.UID)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return -5
	} else if err != nil {
		return -6
	}

	// Se l'utente ha già messo mi piace al post indicato da photoId lo elimino
	_, err = db.c.Exec(`DELETE FROM like WHERE phid = ? AND uid= ?`, photoId, userId)
	if err != nil {
		return -6
	}

	return 0
}
