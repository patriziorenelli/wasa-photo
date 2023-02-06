package database

import (
	"database/sql"
	"errors"
)

// Funzione per verificare se un utente esiste 
func (db *appdbimpl) UserExist(userId int) (int){
	var us User
	row := db.c.QueryRow(`SELECT * from user where id = ?`, userId)
	err := row.Scan(&us.ID, &us.USERNAME)
	if errors.Is(err, sql.ErrNoRows) { return -1 } else {  return 0 }
}

// Funzione per trovare l'username di un utente dal suo id 
func (db *appdbimpl) FindUsername(userId int) (int, string){
	var us User
	row := db.c.QueryRow(`SELECT * from user where id = ?`, userId)
	err := row.Scan(&us.ID, &us.USERNAME)

	if errors.Is(err, sql.ErrNoRows) {
		return -1, ""
	} else { return 0, us.USERNAME}
	 
}

// Funzione per verificare se l'utente userId ha bannato l'user banId 
func (db *appdbimpl) CheckBan(userId int, banId int ) (int){
	// Variabile di tipo Ban usata per i check
	var ban Ban
	row := db.c.QueryRow(`SELECT * from ban where uid = ? and uid2 = ?`, userId, banId)
	err := row.Scan(&ban.UID1, &ban.UID2)
	if !errors.Is(err, sql.ErrNoRows) { return 0 } else { return -1 }
}