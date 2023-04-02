package in_memory

import (
	"errors"
	"fmt"
)

const IDNotFound = "incorrect user id"
const EmailDoesNotExist = "email does not exist"
const UsernameDoesNotExist = "username does not exist"
const EmailDuplication = "this email already taken"
const UsernameDuplication = "this username already taken"

var IDNotFoundError error
var EmailDoesNotExistError error
var UsernameDoesNotExistError error
var EmailDuplicationError error
var UsernameDuplicationError error

func init() {
	fmt.Println("hello from err")
	IDNotFoundError = errors.New(IDNotFound)
	EmailDoesNotExistError = errors.New(EmailDoesNotExist)
	UsernameDoesNotExistError = errors.New(UsernameDoesNotExist)
	EmailDuplicationError = errors.New(EmailDuplication)
	UsernameDuplicationError = errors.New(UsernameDuplication)
	//logger.Logger().Println(IDNotFoundError)
}
