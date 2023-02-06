package database

// VA BENE
func (db *appdbimpl) SetMyUserName(userId int, newUsername string) int {

	// Controllo se l'utente a cui si vuole cambiare il nome esiste
	if db.UserExist(userId) == -1 {
		return -1
	}

	// Controllo che l'username selezionato non sia già utilizzato
	if db.UsernamUsed(newUsername) == 0 {
		return -2
	}

	// Aggiorno il nickname
	_, err := db.c.Exec(`UPDATE user SET username = ? WHERE id = ?`, newUsername, userId)
	// In caso di errore nel cambio username
	if err != nil {
		return -3
	}

	return 0

}
