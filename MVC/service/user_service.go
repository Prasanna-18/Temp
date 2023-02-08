package service

import (
	"github.com/Prasanna-18/temp/MVC/domain"
	"github.com/Prasanna-18/temp/MVC/util"
)

func Getuser(query domain.Query) (*domain.User, *util.ApplicationError) {
	return domain.Getuser(query)
}
