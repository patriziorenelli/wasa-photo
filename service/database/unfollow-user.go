package database

// Va bene
func (db *appdbimpl) UnfollowUser(userId int, unfollowId int) (int, Username) {

	var username Username

	// Controllo che l'utente indicato (colui che smette di seguire esista)
	if db.UserExist(userId) == -1 {
		return -1, username
	}

	errore, user := db.FindUsername(unfollowId)

	// Controllo che l'utente che si vuole smettere seuire esista e prende il suo username
	if errore == -1 {
		return -2, username
	} else {
		username.USERNAME = user
	}

	// Controllo che l'utente che si vuole smettere di seguire non abbia bloccato l'utente
	if db.CheckBan(unfollowId, userId) == 0 {
		return -3, username
	}

	// Controllo che l'utente non abbia bloccato chi vuole smettere di seguire
	if db.CheckBan(userId, unfollowId) == 0 {
		return -4, username
	}

	// Verifico se userId segue l'utente da smettere di seguire
	if db.CheckFollow(userId, unfollowId) == -1 {
		return -5, username
	}

	// Se l'utente segue l'user indicato smette di seguirlo
	_, err := db.c.Exec(`DELETE FROM follow WHERE uid = ? AND uid2 = ?`, userId, unfollowId)
	if err != nil {
		return -6, username
	}

	return 0, username
}
