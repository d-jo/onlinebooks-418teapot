package main

// This file holds the structs

// Listing stores information on listings
type Listing struct {
	ID              int     `json:"id"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	ISBN            string  `json:"isbn"`
	Price           float32 `json:"price"`
	Category        string  `json:"category"`
	SellerName      string  `json:"seller_name"`
	ListingPassword string  `json:"listing_password"`
	Status          string  `json:"status"`
	Buyer           string  `json:"buyer"`
	BillingInfo     string  `json:"billing_info"`
	ShippingInfo    string  `json:"shipping_info"`
}

// GenericJSON is used to read arbitrary JSON as a map
type GenericJSON struct {
	Data map[string]int `json:"data"`
}

// ConfigStruct struct holds non-sensitive information
// about configuration
type ConfigStruct struct {
	WebHost        string            `json:"host"`
	WebPort        string            `json:"port"`
	SQLHost        string            `json:"sql_host"`
	SQLPort        string            `json:"sql_port"`
	DBName         string            `json:"sql_dbname"`
	StartupQueries []string          `json:"sql_init_queries"`
	SQLQueries     map[string]string `json:"sql_queries"`
}

// CredsStruct struct holds sensitive information like
// credentials for the DB
type CredsStruct struct {
	DBUser string `json:"db_user"`
	DBPass string `json:"db_pass"`
}
