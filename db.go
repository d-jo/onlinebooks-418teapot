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
	// get the statement from the config
	query := Config.SQLQueries["create_listing"]
	// prepare the arguments using the lst paramater
	arg := []interface{}{lst.Title, lst.Description, lst.ISBN, lst.Price, lst.Category, lst.SellerName, lst.ListingPassword, lst.Status}
	// prepare the statement
	//stmt, _ := db.Prepare(query)
	log.Println(arg)

	// execute the statement with the args in the array
	_, err := db.Exec(query, arg...)
	// error check
	if err != nil {
		panic(err)
	}
}

func SelectActive() []Listing {
	// get the statement from the config
	query := Config.SQLQueries["select_all_active_listings"]

	// prepare the statement
	//stmt, _ := db.Prepare(query)
	log.Println()

	// execute the query
	results, err := db.Query(query)
	// error check
	if err != nil {
		panic(err)
	}	

	var listings []Listing
	var count = 0

	for results.Next() {
		var listing Listing
		listings[count] = Listing{
			Title: listing.Title, 
			Description: listing.Description, 
			ISBN: listing.ISBN, 
			Price: listing.Price, 
			Category: listing.Category, 
			SellerName: listing.SellerName,
		}

		if err != nil {
			panic(err)
		}

		count++
	}
	return listings
}

// TODO
//func SelectPrivate(password string) Listing {
//}

// TODO
//func Search(keyword string) {
//}

// TODO
//func UpdateListing() {
//
//}
