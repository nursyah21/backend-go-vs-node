// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

type Music struct {
	ID     int64  `json:"id"`
	Userid int64  `json:"userid"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Link   string `json:"link"`
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
