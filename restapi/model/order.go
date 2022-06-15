package model

import (
	"encoding/csv"
	"os"
	"log"
	"strconv"
)

type Order struct {  
    Id int `json:"id"`
    CreatedAt string `json:"created_at"`
    OrderName string `json:"order_name"`
    CustomerId string `json:"customer_id"`
}

func ReadOrdersCSV(filePath string) []Order {
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

	var order_list[]Order

	for i := 1; i < len(records); i++ {
		id, _ :=  strconv.Atoi(records[i][0])
		var created_at = records[i][1]
		var order_name = records[i][2]
		var customer_id = records[i][3]

		order_list = append(order_list, Order{id, created_at, order_name, customer_id})
	}

    return order_list
}