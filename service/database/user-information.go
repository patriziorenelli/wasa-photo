package database

// VA BENE
func (db *appdbimpl) GetUserFollowers(userId int, userId2 int) (int, []UserId) {

	// Controllo che l'utente di cui si vogliono sapere i followers esista
	if db.UserExist(userId) == -1 {
		return -1, nil
	}

	errore := db.UserExist(userId2)

	// Controllo che l'utente che vuole sapere i followers esista
	if errore == -1 {
		return -2, nil
	}

	// Controllo che l'utente di cui si vogliono sapere i followers non abbia bannato l'utente
	if db.CheckBan(userId2, userId) == 0 {
		return -3, nil
	}

	// Controllo che l'utente non abbia bloccato l'utente di cui vuole sapere i followers
	if db.CheckBan(userId, userId2) == 0 {
		return -4, nil
	}

	var follow Follow

	// Prendo i followers
	row, err := db.c.Query(`SELECT * FROM follow WHERE uid2 = ?`, userId2)

	// Errore nella query
	if err != nil {
		return -5, nil
	}

	// Creo l'array che conterrà gli id dei follower
	var followers []UserId
	// Riempo l'array con gli id dei follower
	for row.Next() {
		err = row.Scan(&follow.UID1, &follow.UID2)
		if err == nil {
			followers = append(followers, UserId{int(follow.UID1)})
		}
	}

	return 0, followers

}

// VA BENE
func (db *appdbimpl) GetUserFollowing(userId int, userId2 int) (int, []UserId) {

	// Controllo che l'utente di cui si vogliono sapere i following esista
	if db.UserExist(userId) == -1 {
		return -1, nil
	}

	errore := db.UserExist(userId2)

	// Controllo che l'utente che vuole sapere i following esista
	if errore == -1 {
		return -2, nil
	}

	// Controllo che l'utente di cui si vogliono sapere i following non abbia bannato l'utente
	if db.CheckBan(userId2, userId) == 0 {
		return -3, nil
	}

	// Controllo che l'utente non abbia bloccato l'utente di cui vuole sapere i following
	if db.CheckBan(userId, userId2) == 0 {
		return -4, nil
	}

	var follow Follow

	// Prendo le tuple (oggetti di tipo follow)
	row, err := db.c.Query(`SELECT * FROM follow WHERE uid = ?`, userId2)

	// Errore nella query
	if err != nil {
		return -5, nil
	}

	// Creo l'array che conterrà gli id dei following
	var followers []UserId
	// Riempo l'array con gli id dei following
	for row.Next() {
		err = row.Scan(&follow.UID1, &follow.UID2)
		if err == nil {
			followers = append(followers, UserId{int(follow.UID2)})
		}
	}

	return 0, followers

}
