package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"main/model"
)
func GetOrders(c *gin.Context) {
	orders := model.ReadOrdersCSV("../data/Test task - Postgres - orders.csv")
    c.IndentedJSON(http.StatusOK, orders)
}
