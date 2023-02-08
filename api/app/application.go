package app

import (
	"github.com/Prasanna-18/temp/api/log"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	
}



func Startsapp() {


	log.Log.Info("about to MAP URLS")

	mapUrls()


	log.Log.Info("URLS MAPPED")
	
	// http.HandleFunc("/newusers", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != "GET" {
	// 		w.Write([]byte("invalid endpoint"))
	// 		return
	// 	}
	// 	w.Write([]byte("new user response"))
	// })
	
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
