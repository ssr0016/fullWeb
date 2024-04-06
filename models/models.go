package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/upper/db/v4"
)

var (
	ErrNoMoreRows     = errors.New("no records found")
	ErrDuplicateEmail = errors.New("email already in or database")
	ErrUserNotActive  = errors.New("your account is inactive")
	ErrInvalidLogin   = errors.New("invalid login")
)

type Models struct {
	Users UsersModel
	Post  PostsModel
}

func New(db db.Session) Models {
	return Models{
		Users: UsersModel{db: db},
		Post:  PostsModel{db: db},
	}
}

func convertUpperIDtoInt(id db.ID) int {
	idType := fmt.Sprintf("%T", id)
	if idType == "int64" {
		return int(id.(int64))
	}
	return id.(int)
}

func errHasDuplicate(errr error, key string) bool {
	str := fmt.Sprintf(`ERROR: duplicate key value violates unique constraint "%s"`, key)
	return strings.Contains(errr.Error(), str)
}
