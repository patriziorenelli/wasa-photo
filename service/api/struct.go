package api

import (
	"git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database"
	"reflect"
)

type UserId struct {
	USERID int `json:"userId"`
}

type Username struct {
	USERNAME string `json:"username"`
}

type User struct {
	ID       int    `json:"userId"`
	USERNAME string `json:"username"`
}

type Post struct {
	ID     int    `json:"id"`
	USERID int    `json:"userId"`
	DATE   string `json:"date"`
}

type Ban struct {
	UID1 int `json:"userId1"`
	UID2 int `json:"userId2"`
}

type Follow struct {
	UID1 int `json:"userId1"`
	UID2 int `json:"userId2"`
}

type Like struct {
	PHID int `json:"photoId"`
	UID  int `json:"userId"`
}

type Comment struct {
	CID  int    `json:"commentId"`
	UID  int    `json:"userId"`
	PHID int    `json:"photoId"`
	TEXT string `json:"text"`
}

type Result struct {
	CODE int    `json:"code"`
	TEXT string `json:"result"`
}

type CommentText struct {
	TEXT string `json:"text"`
}

func (userId *UserId) UserIdIsValid() bool {
	var x = reflect.TypeOf(userId.USERID).String()
	return x == "int"

}

func (user *Username) UsernameIsValid() bool {
	return reflect.TypeOf(user.USERNAME).String() == "string" && len(user.USERNAME) >= 6 && len(user.USERNAME) <= 16
}

func (user *Username) UsernameToDatabase() database.Username {
	return database.Username{USERNAME: user.USERNAME}
}

func (comment *CommentText) CommentTextIsValid() bool {
	return len(comment.TEXT) > 0 && len(comment.TEXT) <= 100
}
