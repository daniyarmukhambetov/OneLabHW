package in_memory

import (
	"errors"
)

const IDNotFound = "incorrect user id"
const EmailDoesNotExist = "email does not exist"
const UsernameDoesNotExist = "username does not exist"
const EmailDuplication = "this email already taken"
const UsernameDuplication = "this username already taken"

var IDNotFoundError = errors.New(IDNotFound)
var EmailDoesNotExistError = errors.New(EmailDoesNotExist)
var UsernameDoesNotExistError = errors.New(UsernameDoesNotExist)
var EmailDuplicationError = errors.New(EmailDuplication)
var UsernameDuplicationError = errors.New(UsernameDuplication)
