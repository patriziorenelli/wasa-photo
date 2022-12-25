package database
import (	"errors"
			"database/sql" 
			"fmt"
		)

func (db *appdbimpl) FollowUser(userId int, followId int) (int, error) {

	/* 
	DEVO CONTROLLARE:
		- CHE USERID ESISTA  ->  SQLITE3 NON FA CONTROLLO SULLE FK 
		- CHE FOLLOWID ESISTA -> SQLITE3 NON FA CONTROLLO SULLE FK 
		- CHE L'UTENTE FOLLOWID NON MI ABBIA BANNATO 
		- USERID NON ABBIA BANNATO FOLLOWID 
	FARE LA INSERT IN ban

	*/

	var us User

	// Controllo che l'utente indicato esista
	row := db.c.QueryRow(`SELECT * from user where id = ?`,  userId)
	err := row.Scan(&us.ID, &us.USERNAME)

	if errors.Is(err, sql.ErrNoRows){
		fmt.Print("Utente non esistente")
		return -1, nil
	}

	// Controllo che l'utente che si vuole seguire esista
	row = db.c.QueryRow(`SELECT * from user where id = ?`,  followId)
	err = row.Scan(&us.ID, &us.USERNAME)

	if errors.Is(err, sql.ErrNoRows){
		fmt.Print("Utente2 non esistente")
		return -2, nil
	}
	// Variabile di tipo Ban usata per i check
	var ban Ban 

	// Controllo che l'utente che si vuole seguire non abbia bloccato l'utente che lo vuole seguire 
	row = db.c.QueryRow(`SELECT * from ban where uid = ? and uid2 = ?`,  followId, userId)
	err = row.Scan(&ban.UID1,  &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows){
		fmt.Print("Utente bloccato")
		return -3, nil
	}

	// Controllo che l'utente non abbia bloccato chi vuole seguire 
	row = db.c.QueryRow(`SELECT uid from ban where uid = ? and uid2 = ?`, userId, followId)
	err = row.Scan(&ban.UID1,  &ban.UID2)

	if !errors.Is(err, sql.ErrNoRows){
		fmt.Print("Utente ha bloccato user2")
		return -4, nil
	}

	// Aggiungo il follow nel database 
	_ , err = db.c.Exec(`INSERT INTO follow VALUES (? , ?)`, userId, followId)

	// Caso in cui ci sia già quel follow
	if err != nil && ( err.Error() ) == "UNIQUE constraint failed: follow.uid, follow.uid2"{
		fmt.Print("DUPLICATO")
		return -5, nil
	} else if err != nil {
		return -6, nil
	}
	

	return 0, nil


}
