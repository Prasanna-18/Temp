package repository

import (
	"strings"

	"github.com/Prasanna-18/temp/api/utils_error"
)
type DeleteRepoRequest struct{
	Owner string `json:"owner"`
	Repo  string `json:"repo"`
}

type CreateRepoRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

type DeleteRepoResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type CreateRepoResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
type CreateReposResponse struct {
	StatusCode int `json:"status"`
	Result []CreateRepoResult `json:"result"`
}

type CreateRepoResult struct{
	Response CreateRepoResponse `json:"response"`
	Error utilerrors.ApiError `json:"error"`	
}

func (b CreateRepoRequest)Validate() utilerrors.ApiError{
	b.Name = strings.TrimSpace(b.Name)
	if b.Name == ""{
		return utilerrors.NewBadRequestError("invalid repo name")
	}
	return nil
}