package main

import (
	"fmt"
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "Zchat/api/v1"
	"Zchat/config"
)

func main() {
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}
	corsConfig.ExposeHeaders = []string{"X-Custom-Header"}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))
	r.POST("/login", v1.Login)
	r.POST("/register", v1.Register)
	r.POST("/createGroup", v1.CreateGroup)
	conf := config.GetConfig()
	host := conf.MainConfig.Host
	port := conf.MainConfig.Port

	if err := r.Run(fmt.Sprintf("%s:%d", host, port)); err != nil {
		log.Fatal("server running failed")
		return
	}
}