package database

import "errors"

var (
	ErrDatabaseNotFound = errors.New("not found")
)

type Database interface {
	Connect(params ...interface{}) error
	GetUserCount(userName string) (int, error)
	SetUserCount(userName string, count int) error
	UpdateUserCount(userName string, diff int) error
}
