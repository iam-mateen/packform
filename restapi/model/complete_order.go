type CompleteOrder struct {  
    CompanyId int `json:"company_id"`
    CompanyName string `json:"company_name"`
	UserId string `json:"user_id"`
    Login string `json:"login"`
    Password int `json:"password"`
    Name string `json:"name"`
	CreditCards string `json:"credit_cards"`
    OrderItemId int `json:"order_item_id"`
    DeliveredQuantity int `json:"delivered_quantity"`
    OrderId int `json:"order_id"`
    PricePerUnit float64 `json:"price_per_unit"`
    Quantity int `json:"quantity"`
	Product string `json:"product"`
	Id int `json:"id"`
    CreatedAt string `json:"created_at"`
    OrderName string `json:"order_name"`
    CustomerId string `json:"customer_id"`
}