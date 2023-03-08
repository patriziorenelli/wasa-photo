package database

import (
	"time"
)

// VA BENE
func (db *appdbimpl) CommentPhoto(userId int, photoId int, text string) int {

	// Variabile di tipo Post usata per il check
	var post Post

	// Controllo che l'user indicato esista
	if db.UserExist(userId) == -1 {
		return -1
	}

	// Controllo che il post indicato esista
	post, exist := db.GetPhoto(photoId)
	// Controllo che il post indicato esista e raccolgo i dati del post
	if exist == -1 {
		return -2
	}

	// Controllo che l'utente che  non abbia bannato l'utente a cui si vuole aggiungere un commento
	if db.CheckBan(userId, post.USERID) == 0 {
		return -3
	}

	// Controllo che l'utente che ha pubblicato il post a cui si vuole aggiungere il commento non abbia bannato l'user
	if db.CheckBan(post.USERID, userId) == 0 {
		return -4
	}

	// Data di quando il commento viene postato
	date := time.Now().Format(time.RFC3339)

	// Aggiungo il commento al post
	_, err := db.c.Exec(`PRAGMA foreign_keys = ON; INSERT INTO comment (uid, phid, text, date) VALUES (? , ?, ?, ?)`, userId, photoId, text, date)

	// Errore durante il salvataggio del commento
	if err != nil {
		return -5
	}

	return 0

}
