package services

import (
	"sync"

	"github.com/Prasanna-18/temp/api/configu"
	"github.com/Prasanna-18/temp/api/domain/github"
	"github.com/Prasanna-18/temp/api/domain/repository"
	"github.com/Prasanna-18/temp/api/providers/github_provider"
	"github.com/Prasanna-18/temp/api/utils_error"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(request repository.CreateRepoRequest) (*repository.CreateRepoResponse, utilerrors.ApiError)
	CreateRepos(request []repository.CreateRepoRequest) (repository.CreateReposResponse, utilerrors.ApiError)
}

var (
	RepoService repoServiceInterface
)

func init() {
	RepoService = &repoService{}
}



func (s *repoService) CreateRepo(input repository.CreateRepoRequest) (*repository.CreateRepoResponse, utilerrors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	reponse, err := githubprovider.CreateRepo(configu.GetgithubAcessToken(), request)
	if err != nil {
		return nil, utilerrors.NewApiError(err.StatusCode, err.Message)
	}
	result := repository.CreateRepoResponse{
		Id:    reponse.Id,
		Name:  reponse.Name,
		Owner: reponse.Owner.Login,
	}
	return &result, nil
}

func (s *repoService) CreateRepos(request []repository.CreateRepoRequest) (repository.CreateReposResponse, utilerrors.ApiError) {
	
	input := make(chan repository.CreateRepoResult)
	output := make(chan repository.CreateReposResponse)
 defer close(output)

 var wg sync.WaitGroup

	go handleConcurrent(&wg, input, output)

	for _, current := range request {
		go s.createrepoConcurrent(current, input)
		wg.Add(1)
	}
	wg.Wait()
	close(input)

	result := <-output

	return result, nil
}

func handleConcurrent(wg *sync.WaitGroup, input chan repository.CreateRepoResult, output chan repository.CreateReposResponse) {
	var results repository.CreateReposResponse
for s := range input{
	repoResult:= repository.CreateRepoResult{
		Response: s.Response,
		Error: s.Error,
	}
results.Result = append(results.Result, repoResult)
wg.Done()
}
output <-results
}

func (s *repoService) createrepoConcurrent(current repository.CreateRepoRequest, output chan repository.CreateRepoResult) {
	if err := current.Validate(); err != nil {
		output <- repository.CreateRepoResult{Error: err}
		return
	}
	result, err := s.CreateRepo(current)
	if err != nil {
		output <- repository.CreateRepoResult{Error: err}
	}
	output <- repository.CreateRepoResult{Response: *result}
}
