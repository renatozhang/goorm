package main

import (
	"github.com/gin-gonic/gin"
	"github.com/renatozhang/goorm/controller/account"
	"github.com/renatozhang/goorm/id_gen"
)

func main() {
	router := gin.Default()

	err := id_gen.Init(1)
	if err != nil {
		panic(err)
	}

	router.POST("/api/user/register", account.RegisterHandler)
	router.POST("/api/user/login", account.LoginHandler)

	router.Run(":9090")
}
