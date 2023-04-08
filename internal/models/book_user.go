package models

type BookUser struct {
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	BookName  string `json:"book_name"`
	Returned  bool   `json:"not_returned"`
	TakenDate string `json:"taken_date"`
}

type UserBook struct {
	BookName  string
	Returned  bool
	TakenDate string
}

type UserBookCount struct {
	Username string
	Count    string
}
