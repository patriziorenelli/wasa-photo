package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrUserDoesNotExist = errors.New("User does not exist")

type UserId struct {
	ID int
}

type Username struct {
	USERNAME string
}

type User struct {
	ID       int
	USERNAME string
}

type Post struct {
	ID     int
	USERID int
	DATE   string
}

type Ban struct {
	UID1 int
	UID2 int
}

type Follow struct {
	UID1 int
	UID2 int
}

type Like struct {
	PHID int
	UID  int
}

type Comment struct {
	CID  int
	UID  int
	PHID int
	TEXT string
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// Funzione che controlla se esiste un utente
	UserExist(userId int) int

	// Funzione che gestisce il login
	DoLogin(name Username) (User, error)

	// Crea un utente passandogli il suo nickname
	CreateUser(newNick string) (User, error)

	// Imposta un nuovo nickname ad un utente
	SetMyUserName(userId int, newUsername string) int

	// Segue un altro utente
	FollowUser(userId int, followId int) (int, Username)

	// Smette di seguire un altro utente
	UnfollowUser(userId int, unfollowId int) (int, Username)

	// Banna un utente
	BanUser(userId int, banId int) (int, Username)

	// Toglie il ban ad un utente
	UnBanUser(userId int, unbanId int) (int, Username)

	// Mette like ad un post
	LikePhoto(userId int, photoId int) int

	// Rimuove un like ad un post
	UnlikePhoto(userId int, photoId int) int

	// Aggiunge un commento ad un post
	CommentPhoto(userId int, photoId int, text string) int

	// Rimuove un commento da un post
	UncommentPhoto(userId int, photoId int, commentId int) int

	// Crea un record per una nuova foto all'interno del database
	UploadPhoto(userId int) (int, int)

	// Elimina il record relativo ad una foto dal database
	DeletePhotoRecord(photoId int) int

	// Elimina una foto
	DeletePhoto(userId int, photoId int) int

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='user';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {

		// cambiare tutti gli id in integer per l'autoincrement
		sqlStmt := `CREATE TABLE user (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, username TEXT NOT NULL);
					CREATE TABLE post (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, uid INTEGER NOT NULL, date TEXT NOT NULL, FOREIGN KEY(uid) REFERENCES user(id) ON DELETE CASCADE ON UPDATE NO ACTION);
					CREATE TABLE ban (uid INTEGER NOT NULL, uid2 INTEGER NOT NULL, PRIMARY KEY(uid, uid2), FOREIGN KEY(uid) REFERENCES user(id) ON DELETE CASCADE ON UPDATE NO ACTION, FOREIGN KEY(uid2) REFERENCES user(id) ON DELETE CASCADE ON UPDATE NO ACTION);
					CREATE TABLE like (phid INTEGER NOT NULL, uid INTEGER NOT NULL, PRIMARY KEY(phid, uid), FOREIGN KEY(uid) REFERENCES user(id) ON DELETE CASCADE ON UPDATE NO ACTION, FOREIGN KEY(phid) REFERENCES post(id) ON DELETE CASCADE ON UPDATE NO ACTION);
					CREATE TABLE comment (cid INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, uid INTEGER NOT NULL, phid INTEGER NOT NULL, text TEXT NOT NULL , date TEXT NOT NULL,FOREIGN KEY(uid) REFERENCES user(id) ON DELETE CASCADE ON UPDATE NO ACTION, FOREIGN KEY(phid) REFERENCES post(id) ON DELETE CASCADE ON UPDATE NO ACTION );
					CREATE TABLE follow (uid INTEGER NOT NULL, uid2 INTEGER NOT NULL, PRIMARY KEY(uid, uid2), FOREIGN KEY(uid) REFERENCES user(id) ON DELETE CASCADE ON UPDATE NO ACTION, FOREIGN KEY(uid2) REFERENCES user(id) ON DELETE CASCADE ON UPDATE NO ACTION);
					`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		const q = `PRAGMA foreign_keys = ON`
		_, err = db.Exec(q)

		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		popola := `INSERT INTO user(id, username) VALUES (0000000, "marione_12");
				   INSERT INTO user(id, username) VALUES (0000001, "luca_33");
				   INSERT INTO user(id, username) VALUES (0000002, "Giorgia_Na");
				   INSERT INTO post(id, uid, date) VALUES (0000000001, 0000000, "data1");
				   INSERT INTO post(id, uid, date) VALUES (0000000002,0000000, "data2" );
				   INSERT INTO post(id, uid, date) VALUES (0000000003, 0000001,  "data3" );
				   INSERT INTO post(id, uid, date) VALUES (0000000004, 0000000,  "data4");
				   INSERT INTO post(id, uid, date) VALUES (0000000005, 0000001,  "data5");
				   INSERT INTO follow(uid, uid2) VALUES (000000, 000001);
				   INSERT INTO ban(uid, uid2) VALUES (0000000,0000002 );
				   INSERT INTO like(phid, uid) VALUES (0000000001,0000000);
				   INSERT INTO like(phid, uid) VALUES (0000000001,0000002);
				   INSERT INTO like(phid, uid) VALUES (0000000001,0000001);
				   INSERT INTO like(phid, uid) VALUES (0000000004,0000001);
				   INSERT INTO like(phid, uid) VALUES (0000000003,0000000);
				   INSERT INTO like(phid, uid) VALUES (0000000003,0000001);
				   INSERT INTO like(phid, uid) VALUES (0000000002,0000000);
				   INSERT INTO like(phid, uid) VALUES (0000000002,0000001);
				   INSERT INTO comment(phid, uid, cid, text, date) VALUES (0000000002,0000000, 0000001, "Primo commento", "data1");
				   INSERT INTO comment(phid, uid, cid, text, date) VALUES (0000000002,0000001, 0000002, "Secondo commento", "data4");
				   INSERT INTO comment(phid, uid, cid, text, date) VALUES (0000000003,0000000,0000003, "Terzo commento", "data2");
				   INSERT INTO comment(phid, uid, cid, text, date) VALUES (0000000005,00000001,0000004, "Quarto commeto ", "data3");
				   `
		_, err = db.Exec(popola)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
