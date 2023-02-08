package app

import "github.com/Prasanna-18/temp/api/controllers/repository"

func mapUrls() {
	
router.POST("/respository", repository.CreateRepo)
router.POST("/respositories", repository.CreateRepos)

}