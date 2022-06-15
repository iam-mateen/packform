package repo

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
	
)

func InitializePostgres(url string, port int, user string, password string, dbname string) {
	psqlconnection := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode= disable",url, port, user, password, dbname)

   // Connect to database
	db, err := sql.Open("postgres", psqlconnection)
	if err != nil {
		log.Fatal(err)
	}

	db.Query("CREATE DATABASE packform")
	
	//Company Table
	db.Query(`CREATE TABLE [IF NOT EXISTS] company (
		company_id INT PRIMARY KEY,
		company_name VARCHAR (50) NOT NULL,
	)`)

	//Customer Table
	db.Query(`CREATE TABLE [IF NOT EXISTS] customer (
		user_id VARCHAR (50) PRIMARY KEY,
		login INT NOT NULL,
		password VARCHAR (50) NOT NULL,
		name VARCHAR NOT NULL,
		company_id INT NOT NULL,
		credit_cards VARCHAR (200),
		FOREIGN KEY(company_id) REFERENCES company(company_id)
	)`)

	//Order Table Query
	db.Query(`CREATE TABLE [IF NOT EXISTS] order (
		id INT PRIMARY KEY,
		created_at VARCHAR (50) NOT NULL,
		order_name VARCHAR (50) NOT NULL,
		customer_id VARCHAR (50) NOT NULL,
		FOREIGN KEY(customer_id) REFERENCES customer(user_id)
	)`)

	//Order Item Table
	db.Query(`CREATE TABLE [IF NOT EXISTS] order_item (
		id INT PRIMARY KEY,
		order_id INT NOT NULL,
		price_per_unit real NOT NULL,
		quantity INT NOT NULL,
		product VARCHAR (50) NOT NULL,
		FOREIGN KEY(order_id) REFERENCES order(id)
	)`)

	//Delivery Table
	db.Query(`CREATE TABLE [IF NOT EXISTS] delivery (
		id INT PRIMARY KEY,
		order_item_id INT NOT NULL,
		delivered_quantity INT NOT NULL,
		FOREIGN KEY(order_item_id) REFERENCES order_item(id)
	)`)	
}

func GetOrders() {
	
}