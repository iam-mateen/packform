package main

import (
	"github.com/gin-gonic/gin"
	"main/repo"
	"main/controller"
)

func main() {
	var url = "localhost"
	var port = 5432
	var user = "admin"
	var password = "admin123-F"
	var dbname = "packform"

	repo.InitializePostgres(url, port, user,password, dbname)

   	router := gin.Default()
    router.GET("/orders", controller.GetOrders)
    router.Run("localhost:8083")
}