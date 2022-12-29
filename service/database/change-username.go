package database


func (db *appdbimpl) SetMyUserName(userId int, newUsername string) (Username, error) {

	var use Username

	_, err := db.c.Exec(`UPDATE user SET username = ? WHERE id = ?`, newUsername, userId)

	// In caso di errore nel cambio username
	if err != nil {

		use.USERNAME = "nil"
		return use, err
	} else {
		use.USERNAME = newUsername
		return use, nil
	}

}
