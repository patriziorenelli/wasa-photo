package database

// VA BENE
func (db *appdbimpl) UnBanUser(userId int, unbanId int) (int, Username) {

	var username Username

	// Controllo che l'utente indicato esista
	if db.UserExist(userId) == -1 { return -1, username }

	// Controllo che l'utente a cui si voglia togliere il ban esista
	errore, user := db.FindUsername(unbanId)
	// Controllo che l'utente che si vuole seuire esista e prende il suo username 
	if errore == -1{ return -2, username } else { username.USERNAME =  user }


	// Controllo che l'utente a cui si voglia togliere il ban non abbia bloccato l'utente
	if db.CheckBan(unbanId, userId) == 0 { return -3, username }

	// Controllo che l'utente abbia veramente bannato l'altro user
	if db.CheckBan(userId, unbanId) == -1 { return -4, username }

	// Se l'utente aveva bannato l'altro user allora lo sblocco
	_, err := db.c.Exec(`DELETE FROM ban WHERE uid = ? AND uid2 = ?`, userId, unbanId)
	if err != nil {
		return -5, username
	}

	return 0, username

}
