package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	
}



func Startapp() {
	mapUrls()

	// http.HandleFunc("/newusers", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != "GET" {
	// 		w.Write([]byte("invalid endpoint"))
	// 		return
	// 	}
	// 	w.Write([]byte("new user response"))
	// })
	if err := router.Run(":8081"); err != nil {
		panic(err)
	}
}
