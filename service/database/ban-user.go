package database

// Va bene
func (db *appdbimpl) BanUser(userId int, banId int) (int, Username) {

	var username Username

	// Controllo che l'utente indicato esista
	if db.UserExist(userId) == -1 {
		return -1, username
	}

	// Controllo che l'utente che si vuole bannare esista e prende il suo username
	errore, user := db.FindUsername(banId)
	if errore == -1 {
		return -2, username
	} else {
		username.USERNAME = user
	}

	// Controllo che l'utente che si vuole bannare non abbia bloccato l'utente che lo vuole bloccare
	if db.CheckBan(banId, userId) == 0 {
		return -3, username
	}

	// Controllo che l'utente non abbia già bloccato chi vuole bloccare
	if db.CheckBan(userId, banId) == 0 {
		return -4, username
	}

	// BAN DELL'UTENTE  E SMETTERE DI SEGUIRE L'UTENTE banId E VERIFICARNE IL BUON ESITO
	_, err := db.c.Exec(`INSERT INTO ban VALUES (? , ?); DELETE FROM follow WHERE uid = ? AND uid2 = ?`, userId, banId, userId, banId)

	// Caso in cui già si ha bannato l'user
	if err != nil && (err.Error()) == "UNIQUE constraint failed: ban.uid, ban.uid2" {
		return -4, username
	} else if err != nil {
		return -6, username
	}

	return 0, username

}
