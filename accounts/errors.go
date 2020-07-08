package accounts

import "errors"

var (
	ErrNoNilDB = errors.New("database can't be nil")
)
