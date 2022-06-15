package model


import (
	"encoding/csv"
	"os"
	"log"
	"strconv"
)

type Delivery struct {  
    Id int `json:"id"`
    OrderItemId int `json:"order_item_id"`
    DeliveredQuantity int `json:"delivered_quantity"`
}


func ReadDeliveryCSV(filePath string) []Delivery {
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

	var delivery_list[]Delivery

	for i := 1; i < len(records); i++ {
    	id, _ :=  strconv.Atoi(records[i][0])
    	order_item_id, _ :=  strconv.Atoi(records[i][1])
		delivered_quantity, _ := strconv.Atoi(records[i][2])

		delivery_list = append(delivery_list, Delivery{id, order_item_id, delivered_quantity})
	}

    return delivery_list
}