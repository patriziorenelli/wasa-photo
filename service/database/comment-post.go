package database

import (
	"database/sql"
	"errors"
)

// VA BENE
func (db *appdbimpl) CommentPhoto(userId int, photoId int, text string) int {

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

	// Controllo che l'utente che a cui si vuole aggiungere un commento non abbia bloccato l'utente
	row = db.c.QueryRow(`SELECT * from ban where uid = ? and uid2 = ?`, userId, post.USERID)
	err = row.Scan(&ban.UID1, &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows) {
		return -3
	}

	// Controllo che l'utente non abbia bloccato l'utente che ha pubblicato il post a cui si vuole aggiungere il commento
	row = db.c.QueryRow(`SELECT * from ban where uid = ? and uid2 = ?`, post.USERID, userId)
	err = row.Scan(&ban.UID1, &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows) {
		return -4
	}

	// Aggiungo il commento al post
	_, err = db.c.Exec(`INSERT INTO comment (uid, phid, text) VALUES (? , ?, ?)`, userId, photoId, text)

	// Errore durante il salvataggio del commento
	if err != nil {
		return -5
	}

	return 0

}
