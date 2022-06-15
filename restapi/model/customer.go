package model


import (
	"encoding/csv"
	"os"
	"log"
	"strconv"
)

type Customer struct {  
    UserId string `json:"user_id"`
    Login string `json:"login"`
    Password int `json:"password"`
    Name string `json:"name"`
	CompanyId int `json:"company_id"`
	CreditCards string `json:"credit_cards"`
}

func ReadCustomerCSV(filePath string) []Customer {
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

	var customer_list[]Customer

	for i := 1; i < len(records); i++ {
		var user_id = records[i][0]
    	var login = records[i][1]
    	password, _ :=  strconv.Atoi(records[i][2])
    	var name = records[i][3]
		company_id, _ := strconv.Atoi(records[i][4])
		var credit_cards = records[i][5]

		customer_list = append(customer_list, Customer{user_id, login, password, name, company_id, credit_cards})
	}
    return customer_list
}