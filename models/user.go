package models

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type LoginReq struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}