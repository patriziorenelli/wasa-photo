package database

import (
	"database/sql"
	"errors"
)

// VA BENE
func (db *appdbimpl) UncommentPhoto(userId int, photoId int, commentId int) int {

	if db.UserExist(userId) == -1 {
		return -1
	}

	// Variabile di tipo Post usata per il check
	var post Post
	post, exist := db.GetPhoto(photoId)
	// Controllo che il post indicato esista e raccolgo i dati del post
	if exist == -1 {
		return -2
	}

	// Controllo che l'utente non abbia bannato l'utente che ha postato il post a cui si vuole togliere il commento
	if db.CheckBan(userId, post.USERID) == 0 {
		return -3
	}

	// Controllo che l'utente a cui si vuole togliere il commento non abbia bannato l'utente
	if db.CheckBan(post.USERID, userId) == 0 {
		return -4
	}

	// Variabile di tipo Comment usata per il check
	var comment Comment
	// Seleziono il commento che si vuole eliminare
	row := db.c.QueryRow(`SELECT * FROM comment WHERE cid = ? `, commentId)
	err := row.Scan(&comment.CID, &comment.UID, &comment.PHID, &comment.TEXT, &comment.DATE)

	// Caso in cui il commento indicato non esiste
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return -5
	} else if err != nil {
		return -8
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
