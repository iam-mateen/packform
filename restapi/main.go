package main

import (
	"github.com/gin-gonic/gin"
	"main/repo"
	"main/controller"
)

func main() {
	var url = "postgres://xmrsrssh:Lao1zyOdww1DBddgPzAyu26zLEujo7-F@heffalump.db.elephantsql.com/xmrsrssh"
	var port = 5432
	var user = "admin"
	var password = "Lao1zyOdww1DBddgPzAyu26zLEujo7-F"
	var dbname = "xmrsrssh"

	repo.InitializePostgres(url, port, user,password, dbname)

   	router := gin.Default()
    router.GET("/orders", controller.GetOrders)
    router.Run("localhost:8083")
}