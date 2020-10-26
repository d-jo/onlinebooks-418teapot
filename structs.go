package main

// This file holds the structs

type Listing struct {
}

type GenericJson struct {
	Data map[string]int `json:"data"`
}

// ConfigStruct struct holds non-sensitive information
// about configuration
type ConfigStruct struct {
	WebHost        string   `json:"host"`
	WebPort        string   `json:"port"`
	SQLHost        string   `json:"sql_host"`
	SQLPort        string   `json:"sql_port"`
	DBName         string   `json:"sql_dbname"`
	StartupQueries []string `json:"sql_init_queries"`
}

// CredsStruct struct holds sensitive information like
// credentials for the DB
type CredsStruct struct {
	DBUser string `json:"db_user"`
	DBPass string `json:"db_pass"`
}
