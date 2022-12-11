package database



func (db *appdbimpl) findUsername(username string ) (int, error) {
	
	

	// Ricerca un user in base al suo username 
	const query = `SELECT id FROM user WHERE username = ?`

	rows, err := db.c.Query(query,username)
	if err != nil {
		return -1, err
	}


	var us User 
	err = rows.Scan(&us.ID, us.USERNAME )
	if err != nil{
		return -1, err
	}
	
	return us.ID,nil

}
