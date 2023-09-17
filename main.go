package main

import (
	"github.com/BatyrhanB/gin-edu/database"
	"github.com/BatyrhanB/gin-edu/helper"
	"github.com/gin-gonic/gin"
	//"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	log.Info().Msg("Started server !")
	routes := gin.Default()
	routes.Use(gin.Logger(), gin.Recovery())

	routes.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "<h1>Hello, world</h1>")
	})
	server := http.Server{
		Addr:    ":8000",
		Handler: routes,
	}

	err := database.Connect()
	helper.ErrorPanic(err)

	ServerErr := server.ListenAndServe()
	helper.ErrorPanic(ServerErr)
}
