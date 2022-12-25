package database

import (
	"database/sql"
	"errors"
)
// VA BENE 
func (db *appdbimpl) DoLogin(name Username) (User, error) {

	const query = ` SELECT * from user where username = ?`
	// row := db.c.Query(query, name.USERNAME)
	var us User

	row := db.c.QueryRow(query, name.USERNAME)
	err := row.Scan(&us.ID, &us.USERNAME)

	if errors.Is(err, sql.ErrNoRows){
		us.USERNAME = name.USERNAME
		us.ID = -2
		return us, nil
	}else if err == nil{
		return us, nil
	}else{
		us.ID = -1
		return us, err
	}


}
