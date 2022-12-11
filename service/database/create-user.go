package database

func (db *appdbimpl) createUser(u User) (User, error) {
	
	res, err := db.c.Exec(`INSERT USER fountains ( nickname) VALUES (?)`, u.username)
	if err != nil {
		return u, err
	}
    // QUI DEVO RITORNARE u con il valore id settato con : res.LastInsertId() in modo da ritornare il nuovo utente 
	return  u, err

}
