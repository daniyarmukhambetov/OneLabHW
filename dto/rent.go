package dto

type BookRented struct {
	Book       string
	Users      []string
	TotalPrice float64
}

type UserRented struct {
	Book string
	User string
}

type BookName struct {
	Name string `json:"book_name"`
}
