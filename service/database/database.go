/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetName() (string, error)
	SetName(name string) error

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
		sqlStmt := `CREATE TABLE user (id INTEGER NOT NULL PRIMARY KEY, username TEXT NOT NULL, description TEXT NOT NULL);
					CREATE TABLE post (id INTEGER NOT NULL PRIMARY KEY, description TEXT NOT NULL, userId INTEGER NOT NULL, photo TEXT NOT NULL);
					CREATE TABLE ban (uid INTEGER NOT NULL, uid2 INTEGER NOT NULL, PRIMARY KEY(uid, uid2), FOREIGN KEY(uid) REFERENCES user(id) ON DELETE CASCADE ON UPDATE NO ACTION, FOREIGN KEY(uid2) REFERENCES user(id) ON DELETE CASCADE ON UPDATE NO ACTION);
					CREATE TABLE like (phid INTEGER NOT NULL, uid INTEGER NOT NULL, PRIMARY KEY(phid, uid), FOREIGN KEY(uid) REFERENCES user(id) ON DELETE CASCADE ON UPDATE NO ACTION, FOREIGN KEY(phid) REFERENCES post(id) ON DELETE CASCADE ON UPDATE NO ACTION);
					CREATE TABLE comment (cid INTEGER NOT NULL PRIMARY KEY, uid INTEGER NOT NULL, phid INTEGER NOT NULL, text TEXT NOT NULL,FOREIGN KEY(uid) REFERENCES user(id) ON DELETE CASCADE ON UPDATE NO ACTION, FOREIGN KEY(phid) REFERENCES post(id) ON DELETE CASCADE ON UPDATE NO ACTION );
		`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	    popola := `INSERT INTO user(id, username,description) VALUES (0000000, "marione_12", "sono mario di Roma");
				   INSERT INTO user(id, username,description) VALUES (0000001, "luca_33", "sono Luchino");
				   INSERT INTO user(id, username,description) VALUES (0000002, "Giorgia_Na", "Guardami negli occhi");
				   INSERT INTO post(id, description, userId, photo) VALUES (0000000001, "la 1 foto",0000000, "LINK FOTO 1" );
				   INSERT INTO post(id, description, userId, photo) VALUES (0000000002, "la 2 foto",0000000, "LINK FOTO 2" );
				   INSERT INTO post(id, description, userId, photo) VALUES (0000000003, "la 3 foto",0000001, "LINK FOTO 3" );
				   INSERT INTO post(id, description, userId, photo) VALUES (0000000004, "la 4 foto",0000000, "LINK FOTO 4" );
				   INSERT INTO post(id, description, userId, photo) VALUES (0000000005, "la 5 foto",0000001, "LINK FOTO 5" );
				   INSERT INTO ban(uid, uid2) VALUES (0000000,0000002 );
				   INSERT INTO like(phid, uid) VALUES (0000000001,0000000);
				   INSERT INTO like(phid, uid) VALUES (0000000001,0000002);
				   INSERT INTO like(phid, uid) VALUES (0000000001,0000001);
				   INSERT INTO like(phid, uid) VALUES (0000000004,0000001);
				   INSERT INTO like(phid, uid) VALUES (0000000003,0000000);
				   INSERT INTO like(phid, uid) VALUES (0000000003,0000001);
				   INSERT INTO like(phid, uid) VALUES (0000000002,0000000);
				   INSERT INTO like(phid, uid) VALUES (0000000002,0000001);
				   INSERT INTO comment(phid, uid, cid, text) VALUES (0000000002,0000000, 0000001, "Primo commento");
				   INSERT INTO comment(phid, uid, cid, text) VALUES (0000000002,0000001, 0000002, "Secondo commento");
				   INSERT INTO comment(phid, uid, cid, text) VALUES (0000000003,0000000,0000003, "Terzo commento");
				   INSERT INTO comment(phid, uid, cid, text) VALUES (0000000005,00000001,0000004, "Quarto commeto ");
				   `
		_, err = db.Exec(popola)
		if err != nil {
			return nil, fmt.Errorf("database population error: %w", err)
		}


	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
