package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var fmtstr = "%s:%s@tcp(%s:%s)/%s"

// InitDB connects the the database using info from the config
func InitDB() {
	log.Println("initdb.connect.start")
	dbh, err := sql.Open("mysql", fmt.Sprintf(fmtstr, Creds.DBUser, Creds.DBPass, Config.SQLHost, Config.SQLPort, Config.DBName))
	if err != nil {
		log.Println("initdb.connect.fail")
		panic(err)
	}
	db = dbh
	log.Println("initdb.connect.success")
}

// SQLInits executes the startup SQL statements in the config
func SQLInits() {
	log.Println("sqlinits.inits.start")
	for i, s := range Config.StartupQueries {
		_, err := db.Exec(s)
		if err != nil {
			log.Println("sqlinits.inits.fail")
			log.Println(i)
			panic(err)
		}
	}
	log.Println("sqlinits.inits.end")
}

func (lst Listing) Insert() {
	query := Config.SQLQueries["create_listing"]
	arg := []interface{}{lst.Title, lst.Description, lst.ISBN, lst.Price, lst.Category, lst.SellerName}
	stmt, _ := db.Prepare(query)

	_, err := stmt.Exec(stmt, arg)
	if err != nil {
		panic(err)
	}
}
