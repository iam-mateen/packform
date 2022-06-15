package model


import (
	"encoding/csv"
	"os"
	"log"
	"strconv"
)

type CustomerCompany struct {  
    CompanyId int `json:"company_id"`
    CompanyName string `json:"company_name"`
}

func ReadCustomerCompanyCSV(filePath string) []CustomerCompany {
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

	var customer_company_list[]CustomerCompany

	for i := 1; i < len(records); i++ {
		company_id, _ := strconv.Atoi(records[i][0])
		var company_name = records[i][1]
		customer_company_list = append(customer_company_list, CustomerCompany{company_id, company_name})
	}

    return customer_company_list
}
