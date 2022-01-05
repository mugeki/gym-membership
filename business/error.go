package business

import "errors"

var (
	ErrDuplicateData = errors.New("Data already exist")
	ErrInvalidLoginInfo = errors.New("Username or password is invalid")
	ErrUserNotFound = errors.New("User not found")
	ErrInternalServer = errors.New("Something went wrong")
	ErrUnauthorized = errors.New("User Unauthorized")
	ErrProductNotFound = errors.New("Product not found")
)