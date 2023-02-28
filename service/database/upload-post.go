package database

import (
	"time"
)

// VA BENE
func (db *appdbimpl) UploadPhoto(userId int) (int, int) {

	// Controllo che l'utente esista
	if db.UserExist(userId) == -1 {
		return -1, -1
	}

	// Data di quando la foto viene postata
	data := time.Now().Format(time.RFC3339)

	res, err := db.c.Exec(`INSERT INTO post (uid, date) VALUES (?,?)`, userId, data)
	// In caso di errore nella creazione dell'utente
	if err != nil {
		return -2, -1
	}

	// Prendo l'id del nuovo utente
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return -2, -1
	}

	return 0, int(lastInsertID)
}
