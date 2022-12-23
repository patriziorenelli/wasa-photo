package database
import "fmt"
func (db *appdbimpl) SetMyUserName( userId int, newUsername string) (Username, error) {

	var use Username


	_, err := db.c.Exec(`UPDATE user SET username = ? WHERE id = ?`, newUsername, userId)

	// In caso di errore nel cambio username
	if err != nil {

		use.USERNAME = "nil"
		return use, nil
	} else{
		use.USERNAME = newUsername
		fmt.Print(use)
		return use, nil
	}




}
