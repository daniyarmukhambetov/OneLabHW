package models

type JWT struct {
	Token string `json:"token"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
