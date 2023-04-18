package dto

type BookRented struct {
	Book        string   `json:"book"`
	Users       []string `json:"active_users"`
	TotalAmount float64  `json:"total_amount"`
}

type UserRented struct {
	Book string
	User string
}

type BookName struct {
	Name string `json:"book_name"`
}

type BookTransaction struct {
	Book        string  `json:"book"`
	TotalAmount float64 `json:"total_amount"`
}
