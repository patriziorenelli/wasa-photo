package database

// NON RITORNA TUTTE LE FOTO
func (db *appdbimpl) GetMyStream(userId int, limit int, startIndex int) (int, []CompletePost) {

	// Controllo se l'utente a cui si vuole cambiare il nome esiste
	if db.UserExist(userId) == -1 {
		return -1, nil
	}

	row, err := db.c.Query(`SELECT post.id, post.uid, post.date, user.username
								FROM post, user
								WHERE post.uid = user.id
								AND post.uid IN (
									SELECT uid2 FROM follow 
									WHERE follow.uid = ?
								)
								AND post.uid NOT IN (
									SELECT uid2 FROM ban 
									WHERE ban.uid = ?
								)
								ORDER BY date DESC
								LIMIT ?
								OFFSET ?`, userId, userId, limit, startIndex)

	if err != nil {
		return -2, nil
	}

	var posts []CompletePost
	var post CompletePost

	// Per ogni tuplo inizializzo una variabile di tipo post
	for row.Next() {
		err = row.Scan(&post.ID, &post.USERID, &post.DATE, &post.USERNAME)
		if err != nil {
			return -2, nil
		}
		// Prendo il numero di mi piace
		err, likes := db.GetPhotoLike(userId, post.ID)
		if err != 0 {
			return -2, nil
		}

		post.LIKES = len(likes)

		// Prendo il numero di commenti
		err, comments := db.GetPhotoComment(userId, post.ID)
		if err != 0 {
			return -2, nil
		}
		post.COMMENTS = len(comments)

		// Prendo l'username del proprietario della foto
		err, username := db.FindUsername(post.USERID)
		if err == -1 {
			return -2, nil
		}
		post.USERNAME = username

		// Aggiungo il post all'array
		posts = append(posts, post)

	}

	// Aggiunta del check errore
	if row.Err() != nil {
		return -2, nil
	}

	return 0, posts

}
