package router

import (
	"imdb/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3036"}
	r.Use(cors.New(config))

	r.GET("getmovie", controller.GetMovie)
	r.GET("showlogs", controller.ShowLogs)
	return r
}
