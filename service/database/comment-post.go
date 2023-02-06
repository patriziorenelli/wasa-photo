package database

// VA BENE
func (db *appdbimpl) CommentPhoto(userId int, photoId int, text string) int {

	// Variabile di tipo Post usata per il check
	var post Post

	// Controllo che l'user indicato esista
	if db.UserExist(userId) == -1 {
		return -1
	}

	// Controllo che il post indicato esista
	if db.PhotoExist(photoId) == -1 {
		return -2
	} else {
		post = db.GetPhoto(photoId)
	}

	// Controllo che l'utente che  non abbia bannato l'utente a cui si vuole aggiungere un commento
	if db.CheckBan(userId, post.USERID) == 0 {
		return -3
	}

	// Controllo che l'utente che ha pubblicato il post a cui si vuole aggiungere il commento non abbia bannato l'user
	if db.CheckBan(post.USERID, userId) == 0 {
		return -4
	}

	// Aggiungo il commento al post
	_, err := db.c.Exec(`INSERT INTO comment (uid, phid, text) VALUES (? , ?, ?)`, userId, photoId, text)

	// Errore durante il salvataggio del commento
	if err != nil {
		return -5
	}

	return 0

}
