package controllers

import (
	"net/http"

	"github.com/Prasanna-18/temp/MVC/domain"
	"github.com/Prasanna-18/temp/MVC/service"
	"github.com/gin-gonic/gin"
)

func Getuser(c *gin.Context) {

	var query domain.Query
	if err := c.BindQuery(&query); err != nil {
		c.JSON(400, gin.H{"status": false, "error": err.Error()})
		return
	}


	user, apiErr := service.Getuser(query)
	if apiErr != nil {
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	c.JSON(http.StatusOK, user)
}
