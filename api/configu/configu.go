package configu

import (
	"os"
	
)

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAcessToken = os.Getenv(apiGithubAccessToken)
	Loglevel = "info"
)

func GetgithubAcessToken() string{
	return githubAcessToken
}

func IsProductionn() bool{
return os.Getenv("GO_ENVIRONMENT")== "Production"
}

