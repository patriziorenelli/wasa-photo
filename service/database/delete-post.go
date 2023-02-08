package database

func (db *appdbimpl) DeletePhoto(userId int, photoId int) int {

	post, exist := db.GetPhoto(photoId)
	// Controllo l'esistenza della foto da eliminare
	if exist == -1 {
		return -1
	}
	// Controllo che l'utente che sta richiedendo l'eliminazione è il proprietario del post
	if post.USERID != userId {
		return -2
	}
	// Si tenta l'eliminazione della foto
	if db.DeletePhotoRecord(photoId) != 0 {
		return -3
	}

	return 0
}
