package database
import (	"errors"
			"database/sql" 
			"fmt"
			"strings"
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


	var ban Ban 
	// -----------------------------------

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


	//---------------------------------

	// Aggiunto il follow nel database 
	res , err := db.c.Exec(`INSERT INTO follow VALUES (? , ?)`, userId, followId)
	
	if strings.Contains(err.Error(), "UNIQUE constraint" ){
		fmt.Print("DUPLICATO")
		return -5, nil
	}

	
	
	fmt.Print(res)


	/*
	var use Username


	_, err := db.c.Exec(`UPDATE user SET username = ? WHERE id = ?`, newUsername, userId)

	// In caso di errore nel cambio username
	if err != nil {

		use.USERNAME = "nil"
		return use, err
	} else{
		use.USERNAME = newUsername
		fmt.Print(use)
		return use, nil
	}

*/

	return 0, nil


}
