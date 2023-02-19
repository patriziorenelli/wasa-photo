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

type CompleteUser struct {
	ID        int    `json:"userId"`
	USERNAME  string `json:"userName"`
	POST      int    `json:"numPosts"`
	FOLLOWER  int    `json:"numFollower"`
	FOLLOWING int    `json:"numFollowing"`
}

type Post struct {
	ID     int    `json:"photoId"`
	USERID int    `json:"userId"`
	DATE   string `json:"date"`
}

type CompletePost struct {
	ID       int    `json:"photoId"`
	USERID   int    `json:"userId"`
	USERNAME string `json:"name"`
	LIKES    int    `json:"likes"`
	COMMENTS int    `json:"comments"`
	DATE     string `json:"upladTime"`
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
	UID  int    `json:"userId"`
	NAME string `json:"name"`
	TEXT string `json:"comment"`
	CID  int    `json:"commentId"`
	DATE string `json:"date"`
	PHID int    `json:"photoId"`
}

type Result struct {
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
