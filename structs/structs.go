package structs

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type Product struct {
	ID                  int    `json:"id"`
	Product_name        string `json:"product_name"`
	Product_description string `json:"product_description"`
	Product_stock       int    `json:"product_stock"`
	Product_price       int    `json:"product_price"`
	Created_at          string `json:"created_at"`
	Updated_at          string `json:"updated_at"`
}

type Orders struct {
	ID           int    `json:"id"`
	Users_id     string `json:"users_id"`
	Products_id  string `json:"products_id"`
	Total_price  string `json:"total_price"`
	Total_bought string `json:"total_bought"`
	Created_at   string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
}
