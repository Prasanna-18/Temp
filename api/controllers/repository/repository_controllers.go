package repository

import (
	"net/http"
	"github.com/Prasanna-18/temp/api/domain/repository"
	"github.com/Prasanna-18/temp/api/repository"
	"github.com/Prasanna-18/temp/api/utils_error"
	"github.com/gin-gonic/gin"
)


func CreateRepo(c *gin.Context) {
	
	var request repository.CreateRepoRequest

	// json.NewDecoder(c.Request.Body).Decode(&request)
	
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := utilerrors.NewBadRequestError("INVALID JSON REQUEST")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	result, err := services.RepoService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func CreateRepos(c *gin.Context){
	var request []repository.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := utilerrors.NewBadRequestError("INVALID JSON REQUEST")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	result, err := services.RepoService.CreateRepos(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}