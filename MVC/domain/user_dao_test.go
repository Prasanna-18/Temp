package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestGetuserNoUserFound(t *testing.T) {
	val := struct{
		Id int `form:"id"`
		Name string `form:"name"`
	}{
		Id: 0,
		Name: "ji",
	}
	user, err := Getuser(val)

	 assert.Nil(t,user,"we were not expecting user with id 0")
	 // if user != nil {
	// 	t.Error("we were not expecting user with id 0")
	// }

	 assert.NotNil(t,err,"we were expecting an error when user id is not there")
	// if err == nil {
	// 	t.Error("we were expecting an error when user id is not there")
	// }

	// assert.EqualValues(t,http.StatusNotFound,err.StatusCode)
	// assert.EqualValues(t,"Not found",err.Code)
	// assert.EqualValues(t, fmt.Sprintf("User %v not found", u),err.Message)

 
}
