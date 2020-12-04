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

		err = results.Scan(&listing.ID, &listing.Title, &listing.Description, &listing.ISBN, &listing.Price, &listing.Category, &listing.SellerName, &listing.Status)

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

		err := res.Scan(&listing.ID, &listing.Title, &listing.Description, &listing.ISBN, &listing.Price, &listing.Category, &listing.SellerName, &listing.Status)

		if err != nil {
			panic(err)
		}

		listings = append(listings, listing)
	}

	return listings
}

//SelectPassword function will call select_password query in config.json
//parameter: id. type int
//return value: the hash. type string
func SelectPassword(id int) string {

	query := Config.SQLQueries["select_password"]

	var hash string

	err := db.QueryRow(query, id).Scan(&hash) //Scan puts the result into the variable hash

	if err != nil {
		panic(err)
	}
	//log.Println("inside the SelectPassword function. The returned hash value is " + hash + "\n")
	return hash
}

//DeleteListing deletes a listing with a specific id.
func DeleteListing(id int) {

	query := Config.SQLQueries["delete_listing"]

	_, err := db.Query(query, id)

	if err != nil {
		panic(err.Error())
	} else {
		log.Println("DELETE LISTING in db.go SUCCESSFUL \n")
	}

}

// SelectPrivate selects private listing details and returns a listing
func SelectPrivate(id int) Listing {
	query := Config.SQLQueries["select_listing_private"]
	var lst Listing
	var buyer sql.NullString
	var billing sql.NullString
	var shipping sql.NullString

	err := db.QueryRow(query, id).Scan(&buyer, &billing, &shipping)

	if err != nil {
		panic(err)
		return lst
	}

	lst.Buyer = buyer.String
	lst.BillingInfo = billing.String
	lst.ShippingInfo = shipping.String

	return lst
}

// TODO
func Search(keyword string) []Listing {
	log.Println(keyword)
	var buyer sql.NullString
	var billing sql.NullString
	var shipping sql.NullString

	query := Config.SQLQueries["search_listings"]

	results, err := db.Query(query, keyword, keyword, keyword)

	if err != nil {
		panic(err.Error())
	}
	var books []Listing
	for results.Next() {
		var book Listing

		err = results.Scan(&book.ID, &book.Title, &book.Description, &book.ISBN, &book.Price, &book.Category, &book.SellerName, &book.ListingPassword, &book.Status, &buyer, &billing, &shipping)
		if err != nil {
			panic(err)
		}
		book.Buyer = buyer.String
		book.BillingInfo = billing.String
		book.ShippingInfo = shipping.String
		books = append(books, book)
		log.Println(book.Title)
	}

	return books
}

// TODO
func UpdateListing(listingId int, title string, description string, isbn string, price float32, category string, seller string) {
	//var book Listing

	query := Config.SQLQueries["update_listing"]

	_, err := db.Query(query, title, description, isbn, price, category, seller, listingId)

	if err != nil {
		panic(err)
	}
}

func PurchaseListing(buyer string, billInfo string, shipInfo string, id int) {
	query := Config.SQLQueries["purchase_listing"]

	_, err := db.Query(query, buyer, billInfo, shipInfo, id)

	if err != nil {
		panic(err)
	}
}
