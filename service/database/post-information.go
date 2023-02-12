package database

func (db *appdbimpl) GetPhotoLike(userId int, photoId int) (int, []UserId) {

	// Variabile di tipo post per raccogliere le informazioni di un post
	var post Post

	// Controllo che l'utente indicato esista
	if db.UserExist(userId) == -1 {
		return -1, nil
	}

	// Controllo che il post indicato esista e raccolgo i dati del post
	post, exist := db.GetPhoto(photoId)
	if exist == -1 {
		return -2, nil
	}


	// Controllo che l'utente che ha pubblicato il post di cui si vogliono sapere i like non abbia bannato l'utente
	if db.CheckBan(post.USERID, userId) == 0 {
			return -3, nil
	}

		
	// Controllo che l'utente non abbia bannato l'utente che ha pubblicato la foto di cui si vogliono sapere i like
	if db.CheckBan(userId, post.USERID) == 0 {
		return -4, nil
	}


	var like Like

	// Prendo i followers
	row, err := db.c.Query(`SELECT * FROM like  WHERE phid = ?`, photoId)

	// Errore nella query
	if err != nil {
		return -5, nil
	}

	// Creo l'array che conterrà gli id dei follower
	var userLike []UserId
	// Riempo l'array con gli id dei follower
	for row.Next() {
		err = row.Scan(&like.PHID, &like.UID)
		if err == nil {
			userLike = append(userLike, UserId{int(like.UID)})
		}
	}

	return 0, userLike
}

func (db *appdbimpl) GetPhotoComment(userId int, phId int) (int, []UserId) {

	return 0, nil
}
