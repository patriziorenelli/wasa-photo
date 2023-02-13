package database

// VA BENE
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

	// Prendo i like alla foto
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

// VA BENE
func (db *appdbimpl) GetPhotoComment(userId int, photoId int) (int, []Comment) {

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

	// Controllo che l'utente che ha pubblicato il post di cui si vogliono sapere i commenti non abbia bannato l'utente
	if db.CheckBan(post.USERID, userId) == 0 {
		return -3, nil
	}

	// Controllo che l'utente non abbia bannato l'utente che ha pubblicato la foto di cui si vogliono sapere i commenti
	if db.CheckBan(userId, post.USERID) == 0 {
		return -4, nil
	}

	// Prendo i commenti
	row, err := db.c.Query(`SELECT cid, uid, phid, text, date, username FROM comment, user  WHERE phid = ? AND id = uid`, photoId)

	// Errore nella query
	if err != nil {
		return -5, nil
	}

	// Variabile usata per la scan
	var comment Comment

	// Creo l'array che conterrà i vari commenti dei post
	var commentList []Comment
	// Riempo l'array con i commenti
	for row.Next() {
		err = row.Scan(&comment.CID, &comment.UID, &comment.PHID, &comment.TEXT, &comment.DATE, &comment.NAME)
		if err == nil {
			commentList = append(commentList, Comment{comment.UID, comment.NAME, comment.TEXT, comment.CID, comment.DATE, comment.PHID})
		}
	}

	return 0, commentList
}
