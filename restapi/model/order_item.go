package model


import (
	"encoding/csv"
	"os"
	"log"
	"strconv"
)

type OrderItem struct {  
    Id int `json:"id"`
    OrderId int `json:"order_id"`
    PricePerUnit float64 `json:"price_per_unit"`
    Quantity int `json:"quantity"`
	Product string `json:"product"`
}

func ReadOrderItemsCSV(filePath string) []OrderItem {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

	var delivery_list[]OrderItem

	for i := 1; i < len(records); i++ {
		id, _ :=  strconv.Atoi(records[i][0])
		order_id, _ :=  strconv.Atoi(records[i][1])
		price_per_unit, _ := strconv.ParseFloat(records[i][2], 32)
    	quantity, _ :=  strconv.Atoi(records[i][3])
		var product = records[i][4] 

		delivery_list = append(delivery_list, OrderItem{id, order_id, price_per_unit, quantity, product})
	}

    return delivery_list
}