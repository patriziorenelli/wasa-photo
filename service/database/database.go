package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrUserDoesNotExist = errors.New("User does not exist")

type UserId struct {
	USERID int `json:"userId"`
}

type Username struct {
	USERNAME string `json:"username"`
}

type User struct {
	ID       int    `json:"userId"`
	USERNAME string `json:"username"`
}

type CompleteUser struct {
	ID        int    `json:"userId"`
	USERNAME  string `json:"userName"`
	POST      int    `json:"numPosts"`
	FOLLOWER  int    `json:"numFollower"`
	FOLLOWING int    `json:"numFollowing"`
}

type Post struct {
	ID     int    `json:"photoId"`
	USERID int    `json:"userId"`
	DATE   string `json:"date"`
}

type CompletePost struct {
	ID       int    `json:"photoId"`
	USERID   int    `json:"userId"`
	USERNAME string `json:"name"`
	LIKES    int    `json:"likes"`
	COMMENTS int    `json:"comments"`
	DATE     string `json:"upladTime"`
}

type Ban struct {
	UID1 int `json:"userId1"`
	UID2 int `json:"userId2"`
}

type Follow struct {
	UID1 int `json:"userId1"`
	UID2 int `json:"userId2"`
}

type Like struct {
	PHID int `json:"photoId"`
	UID  int `json:"userId"`
}

type Comment struct {
	UID  int    `json:"userId"`
	NAME string `json:"name"`
	TEXT string `json:"comment"`
	CID  int    `json:"commentId"`
	DATE string `json:"date"`
	PHID int    `json:"photoId"`
}

type Result struct {
	TEXT string `json:"result"`
}

type CommentText struct {
	TEXT string `json:"text"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// Funzione che controlla se esiste un utente
	UserExist(userId int) int

	// Funzione che controlla l'esistenza di un ban
	CheckBan(userId int, banId int) int

	// Funzione che verifica se un post esiste e ritorna il post in caso esista
	GetPhoto(photoId int) (Post, int)

	// Funzione che gestisce il login
	DoLogin(name Username) (User, error)

	// Funzione che restituisce la lista degli id degli utenti che seguono userId2
	GetUserFollowers(userId int, userId2 int) (int, []UserId)

	// Funzione che restituisce la lista degli id degli utente che segue un utente
	GetUserFollowing(userId int, userId2 int) (int, []UserId)

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

	// Funzione che ritorna tutti i commenti di un post
	GetPhotoComment(userId int, phId int) (int, []Comment)

	// Funzione che ritorna la lista dei post di un utente
	GetUserPhotos(userId int, userId2 int) (int, []CompletePost)

	// Funzione che ritorna i like di un post
	GetPhotoLike(userId int, phId int) (int, []UserId)

	// Crea un record per una nuova foto all'interno del database
	UploadPhoto(userId int) (int, int)

	// Elimina il record relativo ad una foto dal database
	DeletePhotoRecord(photoId int) int

	// Elimina una foto
	DeletePhoto(userId int, photoId int) int

	// Ritorna il profilo di un utente
	GetUserProfile(userId int, userId2 int) (int, CompleteUser)

	// Ritorna una parte dello stream relativa agli utenti che lo user segue
	GetMyStream(userId int, limit int, startIndex int) (int, []CompletePost)

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

				   INSERT INTO post(id, uid, date) VALUES (0000000003, 0000001,  "data3" );
				   INSERT INTO post(id, uid, date) VALUES (0000000005, 0000001,  "data5");
				   INSERT INTO post(id, uid, date) VALUES (0000000006, 0000001,  "data7" );
				   INSERT INTO post(id, uid, date) VALUES (0000000009, 0000001,  "data6" );

				   INSERT INTO post(id, uid, date) VALUES (0000000010, 0000000,  "2023-02-28T21:50:02+01:00" );
				   INSERT INTO post(id, uid, date) VALUES (0000000011, 0000000,  "2023-02-28T21:50:27+01:00" );
				   INSERT INTO post(id, uid, date) VALUES (0000000012, 0000000,  "2023-02-28T21:58:36+01:00" );

				   
				   INSERT INTO follow(uid, uid2) VALUES (000000, 000001);
				   INSERT INTO follow(uid, uid2) VALUES (000002, 000001);
				   INSERT INTO ban(uid, uid2) VALUES (0000000,0000002 );
				



				   INSERT INTO like(phid, uid) VALUES (0000000010,0000000);
				   INSERT INTO like(phid, uid) VALUES (0000000010,0000001);
				   INSERT INTO like(phid, uid) VALUES (0000000012,0000001);


				   INSERT INTO like(phid, uid) VALUES (0000000003,0000000);
				   INSERT INTO like(phid, uid) VALUES (0000000003,0000001);
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
