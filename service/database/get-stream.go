package database

// DA FARE
func (db *appdbimpl) getMyStream(userId int, limit int, startIndex int) (int, []CompletePost) {

	// Controllo se l'utente a cui si vuole cambiare il nome esiste
	if db.UserExist(userId) == -1 {
		return -1, nil
	}

	var posts []CompletePost

	/*
		row, err := db.c.Query(`SELECT id, uid, date,username,
									(
										SELECT COUNT(*) AS "likes" FROM like
										WHERE like.phid = post.id
									),
									(
										SELECT COUNT(*) AS "comments" FROM comment
										WHERE comment.phid = post.id
									),

									FROM post , user
									WHERE post.uid = user.id
									AND post.uid IN (
										SELECT uid2 FROM follow WHERE uid = ?
									)
									AND post.uid NOT IN (
										SELECT uid2 FROM ban WHERE ban.uid = ?
									)
									ORDER BY date DESC
									LIMIT ?
									OFFSET ?`, userId, userId, limit, startIndex)
		if err != nil{
			return -2, nil
		}
	*/

	return 0, posts

}
