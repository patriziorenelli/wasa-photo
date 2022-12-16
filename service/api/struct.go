package api

import( "git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database" 
   		"reflect" 
	)

type UserId struct {
	ID int `json:"userId"`
}

type Username struct {
	USERNAME string `json:"username"`
}

type User struct {
	ID  int `json:"id"`
	USERNAME  string `json:"username"`
}

type Post struct {
	ID          int
	DESCRIPTION string
	USERID      int
	PHOTO       string
}

type Ban struct {
	UID1  int
	UID2 int
}

type Like struct {
	PHID int
	UID  int
}

type Comment struct {
	CID  int
	UID  int
	PHID int
	TEXT string
}

func (userId *UserId) UserIdIsValid() bool {
	var x = reflect.TypeOf(userId.ID).String()
	return  x == "int"
		
}

func (user *Username) UsernameIsValid() bool {
	return reflect.TypeOf(user.USERNAME).String() == "string"
}

func (user *Username) UsernameToDatabase() database.Username {
	return database.Username{ USERNAME: user.USERNAME, }
}


func (id *UserId) FromUserDatabase( i database.User) {
	id.ID = i.ID
}



func (u *User) FromDatabase(us database.User) {
	u.ID = us.ID
	u.USERNAME = us.USERNAME
}


