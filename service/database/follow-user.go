package database

func (db *appdbimpl) FollowUser(userId int, followId int) (int, Username) {

	var username Username

	// Controllo che l'utente che sta richiedendo il follow esista
	if db.UserExist(userId) == -1 {
		return -1, username
	}

	errore, user := db.FindUsername(followId)

	// Controllo che l'utente che si vuole seuire esista e prende il suo username
	if errore == -1 {
		return -2, username
	} else {
		username.USERNAME = user
	}

	// Controllo che l'utente che si vuole seguire non abbia bannato l'utente che lo vuole seguire
	if db.CheckBan(followId, userId) == 0 {
		return -3, username
	}

	// Controllo che l'utente non abbia bloccato chi vuole seguire
	if db.CheckBan(userId, followId) == 0 {
		return -4, username
	}

	// Aggiungo il follow nel database
	_, err := db.c.Exec(`INSERT INTO follow VALUES (? , ?)`, userId, followId)

	// Caso in cui ci sia già quel follow
	if err != nil && (err.Error()) == "UNIQUE constraint failed: follow.uid, follow.uid2" {
		return -5, username
	} else if err != nil {
		return -6, username
	}

	return 0, username

}
