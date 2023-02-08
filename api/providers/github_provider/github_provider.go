package githubprovider

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Prasanna-18/temp/api/clients/restclient"
	"github.com/Prasanna-18/temp/api/domain/github"
)

const (
	headerAuth    = "Authorization"
	headerAuthFmt = "token %s"
	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthHeader(accesstoken string) string {
	return fmt.Sprintf(headerAuthFmt, accesstoken)
}

func CreateRepo(accesstoken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {

	headers := http.Header{}
	headers.Set(headerAuth, getAuthHeader(accesstoken))

	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Printf("error when trying to create new repo in github %s", err.Error())
		return nil, &github.GithubErrorResponse{StatusCode:http.StatusInternalServerError,Message:err.Error()}
	}


	bytes, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError,Message: "invalid response body"}
	}

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse

		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError,Message: "invalid json response body"}
		}

		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse

	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Printf("error when trying to unmarshal create repo successful %s", err.Error())
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "error when trying to unmarshal github create repo response"}
	}
	return &result, nil
}
