package main

import (
	"log"
	v1 "Zchat/api/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", v1.Login)
	r.POST("/register", v1.Register)
	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal("server running failed")
		return
	}
}