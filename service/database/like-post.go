package database

// VA BENE
func (db *appdbimpl) LikePhoto(userId int, photoId int) int {

	// Variabile di tipo post per raccogliere le informazioni di un post
	var post Post

	// Controllo che l'utente indicato esista
	if db.UserExist(userId) == -1 {
		return -1
	}

	// Controllo che il post indicato esista
	post, exist := db.GetPhoto(photoId)
	// Controllo che il post indicato esista e raccolgo i dati del post
	if exist == -1 {
		return -2
	}

	// Controllo che l'utente non abbia bannato l'utente a cui si vuole mettere mi piace
	if db.CheckBan(userId, post.USERID) == 0 {
		return -3
	}

	// Controllo che l'utente che ha pubblicato il post a cui si vuole mettere mi piace non abbia bannato l'utente
	if db.CheckBan(post.USERID, userId) == 0 {
		return -4
	}

	// Aggiungo il like al post
	_, err := db.c.Exec(`PRAGMA foreign_keys = ON; INSERT INTO like VALUES (? , ?)`, photoId, userId)

	// Caso in cui ci sia già quel like
	if err != nil && (err.Error()) == "UNIQUE constraint failed: like.phid, like.uid" {
		return -5
	} else if err != nil {
		return -6
	}

	return 0

}
