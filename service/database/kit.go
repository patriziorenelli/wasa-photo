package database

import (
	"database/sql"
	"errors"
)

// Funzione per verificare se un utente esiste
func (db *appdbimpl) UserExist(userId int) int {
	var us User
	row := db.c.QueryRow(`SELECT * from user where id = ?`, userId)
	err := row.Scan(&us.ID, &us.USERNAME)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return -1
	} else {
		return 0
	}
}

// Funzione per trovare l'username di un utente dal suo id
func (db *appdbimpl) FindUsername(userId int) (int, string) {
	var us User
	row := db.c.QueryRow(`SELECT * from user where id = ?`, userId)
	err := row.Scan(&us.ID, &us.USERNAME)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return -1, ""
	} else {
		return 0, us.USERNAME
	}

}

// Funzione per verificare se l'utente userId ha bannato l'user banId
func (db *appdbimpl) CheckBan(userId int, banId int) int {
	// Variabile di tipo Ban usata per i check
	var ban Ban
	row := db.c.QueryRow(`SELECT * from ban where uid = ? and uid2 = ?`, userId, banId)
	err := row.Scan(&ban.UID1, &ban.UID2)

	if err == nil || !errors.Is(err, sql.ErrNoRows) {
		return 0
	} else {
		return -1
	}
}

// Funzione per verificare se un username è già in uso
func (db *appdbimpl) UsernamUsed(newUsername string) int {
	var us User
	row := db.c.QueryRow(`SELECT * from user where username = ?`, newUsername)
	err := row.Scan(&us.ID, &us.USERNAME)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return -1
	} else {
		return 0
	}
}

// Funzione che verifica se userId segue followId
func (db *appdbimpl) CheckFollow(userId int, followId int) int {
	var follow Follow
	row := db.c.QueryRow(`SELECT * FROM follow WHERE uid = ? AND uid2 = ?`, userId, followId)
	err := row.Scan(&follow.UID1, &follow.UID2)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return -1
	} else {
		return 0
	}
}

// Funzione che verifica se un post esiste e ritorna il post in caso esista
func (db *appdbimpl) GetPhoto(photoId int) (Post, int) {
	var post Post
	row := db.c.QueryRow(`SELECT * from post where id = ?`, photoId)
	err := row.Scan(&post.ID, &post.USERID, &post.DATE)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return post, -1
	} else {
		return post, 0
	}
}

// Elimina un post dal database
func (db *appdbimpl) DeletePhotoRecord(photoId int) int {
	_, err := db.c.Exec(`PRAGMA foreign_keys = ON; DELETE FROM post WHERE id = ? `, photoId)
	if err != nil {
		return -1
	} else {
		return 0
	}
}
