package database

// Va bene
func (db *appdbimpl) CheckUserBan(userId int, banId int) (int, UserId) {

	var usId UserId

	// Controllo che l'utente indicato esista
	if db.UserExist(userId) == -1 {
		return -1, usId
	}

	// Controllo che l'utente che si vuole bannare esista
	if db.UserExist(banId) == -1 {
		return -2, usId
	}

	// Controllo che l'utente che si vuole bannare non abbia bloccato l'utente che lo vuole bloccare
	if db.CheckBan(banId, userId) == 0 {
		return -3, usId
	}

	// Controllo se l'utente abbia già bloccato chi vuole bloccare
	if db.CheckBan(userId, banId) == 0 {
		usId.USERID = banId
		return 0, usId
	} else {
		return -4, usId
	}

}
