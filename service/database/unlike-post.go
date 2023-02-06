package database

import (
	"database/sql"
	"errors"
)

// VA BENE
func (db *appdbimpl) UnlikePhoto(userId int, photoId int) int {

	// Variabile di tipo Post usata per il check
	var post Post

	if db.UserExist(userId) == -1 {
		return -1
	}

	// Controllo che il post indicato esista
	if db.PhotoExist(photoId) == -1 {
		return -2
	} else {
		post = db.GetPhoto(photoId)
	}

	// Controllo che l'utente che ha postato il post a cui si vuole togliere il like non abbia bloccato l'utente
	if db.CheckBan(userId, post.USERID) == 0 {
		return -3
	}

	// Controllo che l'utente non abbia bloccato l'utente a cui si vuole togliere il mi piace
	if db.CheckBan(post.USERID, userId) == 0 {
		return -4
	}

	// Controllo che l'utente abbia messo mi piace al post
	var like Like

	row := db.c.QueryRow(`SELECT * FROM like WHERE phid = ? AND uid = ?`, photoId, userId)
	err := row.Scan(&like.PHID, &like.UID)

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
