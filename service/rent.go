package service

import (
	"encoding/json"
	"fmt"
	"hw1/dto"
	"hw1/pkg"
	"hw1/storage"
	"log"
)

type Rent struct {
	Repo *storage.Storage
}

func (s *Rent) RentedBooks() ([]dto.BookRented, error) {
	res, err := s.Repo.BookUser.List()
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	curName := res[0].BookName
	book := dto.BookRented{
		Book:        curName,
		Users:       make([]string, 0),
		TotalAmount: 0,
	}
	ret := make([]dto.BookRented, 0)
	for _, r := range res {
		if curName == r.BookName {
			book.Users = append(book.Users, r.Username)
		} else {
			ret = append(ret, book)
			book = dto.BookRented{
				Book:        r.BookName,
				Users:       []string{r.Username},
				TotalAmount: 0.00,
			}
			curName = r.BookName
		}
	}
	ret = append(ret, book)
	var bookTransactions []dto.BookTransaction
	response := pkg.MakeRequest("http://127.0.0.1:8000/transactions", "GET")
	if err != nil {
		return nil, err
	}
	resBytes := []byte(response)
	_ = json.Unmarshal(resBytes, &bookTransactions)
	if len(ret) != len(bookTransactions) {
		log.Println("book transactions and rented books size has different sizes")
	}
	j := 0
	for i := 0; i < len(ret); i++ {
		if j >= len(bookTransactions) {
			break
		}
		for ret[i].Book != bookTransactions[j].Book && j < len(bookTransactions) {
			j++
		}
		if j < len(bookTransactions) {
			ret[i].TotalAmount = bookTransactions[j].TotalAmount
		}
	}
	return ret, nil
}

func (s *Rent) RentBook(username string, bookName string) (dto.UserRented, error) {
	user, err := s.Repo.User.Retrieve(username)
	if err != nil {
		return dto.UserRented{}, nil
	}
	var data struct {
		UserID   int     `json:"user_id"`
		BookName string  `json:"book_name"`
		Amount   float64 `json:"amount"`
	}
	var body struct {
		UserID    int     `json:"user_id"`
		BookName  string  `json:"book_name"`
		Amount    float64 `json:"amount"`
		ID        int64   `json:"id"`
		CreatedAt string  `json:"created_at"`
	}
	data.BookName = bookName
	data.Amount = 100
	data.UserID = user.ID
	fmt.Println(data)
	b, _ := json.Marshal(&data)
	response := pkg.MakePostRequest("http://127.0.0.1:8000/transactions", b)
	resBytes := []byte(response)
	err = json.Unmarshal(resBytes, &body)
	if err != nil {
		fmt.Println("there is", err, body)
		return dto.UserRented{}, nil
	}
	res, err := s.Repo.BookUser.Create(bookName, user.ID)
	if err != nil {
		return dto.UserRented{}, err
	}
	var ret dto.UserRented
	ret.Book = res.BookName
	ret.User = res.Username
	return ret, nil
}

func NewRent(repo *storage.Storage) *Rent {
	return &Rent{Repo: repo}
}

type IRent interface {
	RentedBooks() ([]dto.BookRented, error)
	RentBook(string, string) (dto.UserRented, error)
}
