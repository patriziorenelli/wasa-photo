package database

// Va bene
func (db *appdbimpl) CheckUserBan(userId int, banId int) (int, UserId) {

	var usId UserId

	// Controllo che l'utente indicato esista
	if db.UserExist(userId) == -1 {
		return -1, usId
	}

	// Controllo che l'utente di cui si vuole controllare il ban esista
	if db.UserExist(banId) == -1 {
		return -2, usId
	}

	// Controllo che l'utente di cui si vuole sapere il ban non sia bannato
	if db.CheckBan(banId, userId) == 0 {
		return -3, usId
	}

	// Controllo se l'utente abbia bloccato l'altro utente
	if db.CheckBan(userId, banId) == 0 {
		usId.USERID = banId
		return 0, usId
	} else {
		return -4, usId
	}

}
