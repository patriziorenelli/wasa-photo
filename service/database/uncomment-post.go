package database

import (
	"database/sql"
	"errors"
)

// VA BENE
func (db *appdbimpl) UncommentPhoto(userId int, photoId int, commentId int) int {

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

	// Controllo che l'utente che ha postato il post a cui si vuole togliere il commento non abbia bloccato l'utente
	row = db.c.QueryRow(`SELECT * from ban where uid = ? and uid2 = ?`, userId, post.USERID)
	err = row.Scan(&ban.UID1, &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows) {
		return -3
	}
	// Controllo che l'utente non abbia bloccato l'utente a cui si vuole togliere il commento
	row = db.c.QueryRow(`SELECT * from ban where uid = ? and uid2 = ?`, post.USERID, userId)
	err = row.Scan(&ban.UID1, &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows) {
		return -4
	}

	// Variabile di tipo Comment usata per il check
	var comment Comment
	// Seleziono il commento che si vuole eliminare
	row = db.c.QueryRow(`SELECT * FROM comment WHERE cid = ? `, commentId)
	err = row.Scan(&comment.CID, &comment.UID, &comment.PHID, &comment.TEXT)

	// Caso in cui il commento indicato non esiste
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return -5
	} else if err != nil {
		return -6
	}

	// Caso in cui l'utente non è autorizzato ad eliminare il commento
	if comment.UID != userId {
		return -6
	}

	// Caso in cui il commento non è associato alla foto indicata dall'utente
	if comment.PHID != photoId {
		return -7
	}

	// Se il commento esiste e l'utente può eliminarlo lo elimino
	_, err = db.c.Exec(`DELETE FROM comment WHERE cid = ?`, commentId)
	if err != nil {
		return -8
	}

	return 0
}
