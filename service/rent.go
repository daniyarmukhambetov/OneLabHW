package service

import (
	"hw1/dto"
	"hw1/storage"
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
		Book:       curName,
		Users:      make([]string, 0),
		TotalPrice: 0,
	}
	ret := make([]dto.BookRented, 0)
	for _, r := range res {
		if curName == r.BookName {
			book.Users = append(book.Users, r.Username)
		} else {
			ret = append(ret, book)
			book = dto.BookRented{
				Book:       r.BookName,
				Users:      []string{r.Username},
				TotalPrice: 0.00,
			}
			curName = r.BookName
		}
	}
	ret = append(ret, book)
	return ret, nil
}

func (s *Rent) RentBook(username string, bookName string) (dto.UserRented, error) {
	user, err := s.Repo.User.Retrieve(username)
	if err != nil {
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
