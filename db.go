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

// Insert is called on a listing, and the data in that
// instance of the object will be added to the database
func (lst Listing) Insert() (int64, error) {
	// get the statement from the config
	query := Config.SQLQueries["create_listing"]
	// prepare the arguments using the lst paramater
	arg := []interface{}{lst.Title, lst.Description, lst.ISBN, lst.Price, lst.Category, lst.SellerName, lst.ListingPassword, lst.Status}
	// prepare the statement
	//stmt, _ := db.Prepare(query)
	log.Println(arg)

	// execute the statement with the args in the array
	res, err := db.Exec(query, arg...)
	// error check
	if err != nil {
		panic(err)
	}
	return res.LastInsertId()
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

		err = results.Scan(&listing.ID, &listing.Title, &listing.Description, &listing.ISBN, &listing.Price, &listing.Category, &listing.SellerName)

		listings = append(listings, listing)

		if err != nil {
			panic(err)
		}

		count++
	}
	return listings
}

func SelectPublicListingDetails(id int) []Listing {
	query := Config.SQLQueries["select_listing_public"]

	res, err := db.Query(query, id)

	if err != nil {
		panic(err)
	}

	var listings []Listing

	for res.Next() {
		var listing Listing

		err := res.Scan(&listing.ID, &listing.Title, &listing.Description, &listing.ISBN, &listing.Price, &listing.Category, &listing.SellerName)

		if err != nil {
			panic(err)
		}

		listings = append(listings, listing)
	}

	return listings
}

// TODO
//func SelectPrivate(password string) Listing {
//}

// TODO
func Search(keyword string) []Listing {
	log.Println(keyword)

	query := Config.SQLQueries["search_listings"]

	results, err := db.Query(query, keyword)

	//results, err := db.Query("SELECT * FROM Listings WHERE title LIKE '%?%' OR description LIKE '%?%' OR isbn LIKE '%?%'")

	if err != nil {
		panic(err.Error())
	}
	var books []Listing
	for results.Next() {
		var book Listing

		err = results.Scan(&book.ID, &book.Title, &book.Description, &book.ISBN, &book.Price, &book.Category, &book.SellerName, &book.ListingPassword, &book.Status, &book.Buyer, &book.BillingInfo, &book.ShippingInfo)
		if err != nil {
			panic(err)
		}

		books = append(books, book)
		log.Println(book.Title)
	}

	return books
}

// TODO
//func UpdateListing() {
//
//}
