package database

func (db *appdbimpl) DoLogin(name Username) (User, error) {
	var us User
	const query = ` SELECT * from user where username = ?`
	row, err := db.c.Query(query, name.USERNAME)

	if err != nil {
		us.ID = -1
		return us, err
	}
	// Read all fountains in the resultset

	err = row.Scan(&us.ID, &us.USERNAME)

	if err != nil {
		us.ID = -1
		return us, err
	}

	return us, nil
}
