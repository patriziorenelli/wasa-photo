package database
import (	"errors"
			"database/sql" 
		)

func (db *appdbimpl) FollowUser(userId int, followId int) ( int, Username) {

	var username Username
	var us User

	// Controllo che l'utente indicato esista
	row := db.c.QueryRow(`SELECT * from user where id = ?`,  userId)
	err := row.Scan(&us.ID, &us.USERNAME)

	if errors.Is(err, sql.ErrNoRows){
		return -1, username
	}

	// Controllo che l'utente che si vuole seguire esista
	row = db.c.QueryRow(`SELECT * from user where id = ?`,  followId)
	err = row.Scan(&us.ID, &us.USERNAME)

	if errors.Is(err, sql.ErrNoRows){
		return -2, username
	}

	username.USERNAME = us.USERNAME

	// Variabile di tipo Ban usata per i check
	var ban Ban 

	// Controllo che l'utente che si vuole seguire non abbia bloccato l'utente che lo vuole seguire 
	row = db.c.QueryRow(`SELECT * from ban where uid = ? and uid2 = ?`,  followId, userId)
	err = row.Scan(&ban.UID1,  &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows){
		return -3, username
	}

	// Controllo che l'utente non abbia bloccato chi vuole seguire 
	row = db.c.QueryRow(`SELECT uid from ban where uid = ? and uid2 = ?`, userId, followId)
	err = row.Scan(&ban.UID1,  &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows){
		return -4, username
	}

	// Aggiungo il follow nel database 
	_ , err = db.c.Exec(`INSERT INTO follow VALUES (? , ?)`, userId, followId)

	// Caso in cui ci sia già quel follow
	if err != nil && ( err.Error() ) == "UNIQUE constraint failed: follow.uid, follow.uid2"{
		return -5, username
	} else if err != nil {
		return -6, username
	}
	


	return 0, username


}
