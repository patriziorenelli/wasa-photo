package api

import "git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database"

type userId struct {
	id string `json:"userId"`
}

type username struct {
	username string `json:"username"`
}

type User struct {
	id  string `json:"id"`
	username  string `json:"username"`
}

type Post struct {
	id          string
	description string
	userId      string
	photo       string
}

type Ban struct {
	uid  string
	uid2 string
}

type Like struct {
	phid string
	uid  string
}

type Comment struct {
	cid  string
	uid  string
	phid string
	text string
}

func (userId *userId) userIdIsValid() bool {
	return 6 <= len(userId.id) && len(userId.id) <= 16
}

func (user *username) usernameIsValid() bool {
	return 6 <= len(user.username) && len(user.username) <= 16
}

// ToDatabase returns the user in a database-compatible representation
func (user *User) ToDatabase() database.User {
	return database.User{}
}

// FromDatabase populates the struct with data from the database, overwriting all values.
// You might think this is code duplication, which is correct. However, it's "good" code duplication because it allows
// us to uncouple the database and API packages.
// Suppose we were using the "database.Fountain" struct inside the API package; in that case, we were forced to conform
// either the API specifications to the database package or the other way around. However, very often, the database
// structure is different from the structure of the REST API.
// Also, in this way the database package is freely usable by other packages without the assumption that structs from
// the database should somehow be JSON-serializable (or, in general, serializable).
func (u *User) FromDatabase(user database.User) {
	u.id = user.id
}



