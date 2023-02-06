package database

func (db *appdbimpl) CreateUser(us string) (User, error) {
	var u User
	res, err := db.c.Exec(`INSERT INTO user (username) VALUES (?)`, us)
	// In caso di errore nella creazione dell'utente
	if err != nil {
		u.ID = -1
		return u, err
	}

	// Prendo l'id del nuovo utente
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		u.ID = -1
		return u, err
	}

	u.ID = int(lastInsertID)
	u.USERNAME = us
	return u, nil

}
