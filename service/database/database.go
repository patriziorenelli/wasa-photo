package database

import (
	"database/sql"
	"errors"
	"fmt"
)
var ErrUserDoesNotExist = errors.New("User does not exist")

type Id struct{
	id string
}

type User struct {
	id  string
	username  string
}

type Post struct {
	id  string
	description string
	userId string
	photo string 
}

type Ban struct {
	uid   string
	uid2  string
}

type Like struct {
	phid string
	uid  string
}

type Comment struct {
	cid  string
	uid  string
	phid string 
	text string
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	
	doLogin(username string) (Id , error)
	createUser(id string) (User, error)
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
		sqlStmt := `CREATE TABLE user (id TEXT NOT NULL PRIMARY KEY AUTOINCREMENT, username TEXT NOT NULL);
					CREATE TABLE post (id TEXT NOT NULL PRIMARY KEY AUTOINCREMENT,  description TEXT NOT NULL, uid TEXT NOT NULL, photo TEXT NOT NULL);
					CREATE TABLE ban (uid TEXT NOT NULL, uid2 TEXT NOT NULL, PRIMARY KEY(uid, uid2), FOREIGN KEY(uid) REFERENCES user(id) ON DELETE CASCADE ON UPDATE NO ACTION, FOREIGN KEY(uid2) REFERENCES user(id) ON DELETE CASCADE ON UPDATE NO ACTION);
					CREATE TABLE like (phid TEXT NOT NULL, uid TEXT NOT NULL, PRIMARY KEY(phid, uid), FOREIGN KEY(uid) REFERENCES user(id) ON DELETE CASCADE ON UPDATE NO ACTION, FOREIGN KEY(phid) REFERENCES post(id) ON DELETE CASCADE ON UPDATE NO ACTION);
					CREATE TABLE comment (cid TEXT NOT NULL PRIMARY KEY AUTOINCREMENT, uid TEXT NOT NULL, phid TEXT NOT NULL, text TEXT NOT NULL,FOREIGN KEY(uid) REFERENCES user(id) ON DELETE CASCADE ON UPDATE NO ACTION, FOREIGN KEY(phid) REFERENCES post(id) ON DELETE CASCADE ON UPDATE NO ACTION );
		`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	    popola := `INSERT INTO user(id, username,description) VALUES ("0000000", "marione_12");
				   INSERT INTO user(id, username,description) VALUES ("0000001", "luca_33");
				   INSERT INTO user(id, username,description) VALUES ("0000002", "Giorgia_Na");
				   INSERT INTO post(id, description, uid, photo) VALUES ("0000000001", "la 1 foto","0000000", "LINK FOTO 1" );
				   INSERT INTO post(id, description, uid, photo) VALUES ("0000000002", "la 2 foto","0000000", "LINK FOTO 2" );
				   INSERT INTO post(id, description, uid, photo) VALUES ("0000000003", "la 3 foto","0000001", "LINK FOTO 3" );
				   INSERT INTO post(id, description, uid, photo) VALUES ("0000000004", "la 4 foto","0000000", "LINK FOTO 4" );
				   INSERT INTO post(id, description, uid, photo) VALUES ("0000000005", "la 5 foto","0000001", "LINK FOTO 5" );
				   INSERT INTO ban(uid, uid2) VALUES ("0000000","0000002" );
				   INSERT INTO like(phid, uid) VALUES ("0000000001","0000000");
				   INSERT INTO like(phid, uid) VALUES ("0000000001","0000002");
				   INSERT INTO like(phid, uid) VALUES ("0000000001","0000001");
				   INSERT INTO like(phid, uid) VALUES ("0000000004","0000001");
				   INSERT INTO like(phid, uid) VALUES ("0000000003","0000000");
				   INSERT INTO like(phid, uid) VALUES ("0000000003","0000001");
				   INSERT INTO like(phid, uid) VALUES ("0000000002","0000000");
				   INSERT INTO like(phid, uid) VALUES ("0000000002","0000001");
				   INSERT INTO comment(phid, uid, cid, text) VALUES ("0000000002","0000000", "0000001", "Primo commento");
				   INSERT INTO comment(phid, uid, cid, text) VALUES ("0000000002","0000001", "0000002", "Secondo commento");
				   INSERT INTO comment(phid, uid, cid, text) VALUES ("0000000003","0000000","0000003", "Terzo commento");
				   INSERT INTO comment(phid, uid, cid, text) VALUES ("0000000005","00000001","0000004", "Quarto commeto ");
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
		   

