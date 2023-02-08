package domain

import (
	"fmt"
	"net/http"

	"github.com/Prasanna-18/temp/MVC/util"
)

var (
	users = map[int64]*User{
		123: {Id: 123, Firstname: "HI", Lastname: "Bye", Email: "abc@gmail.com"},
	}
)

func Getuser(query Query) (*User, *util.ApplicationError) {
	if user := users[int64(query.Id)]; user != nil {
		if user.Firstname == query.Name {
			return user, nil
		}
		return nil, &util.ApplicationError{
			Message:    fmt.Sprintf("User %v found but name %s mismatch", query.Id,query.Name),
			StatusCode: http.StatusNotFound,
			Code:       "Not found",
		}
	}
	return nil, &util.ApplicationError{
		Message:    fmt.Sprintf("User %v not found", query.Id),
		StatusCode: http.StatusNotFound,
		Code:       "Not found",
	}

}
