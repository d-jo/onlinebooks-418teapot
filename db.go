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

// TODO
func UpdateListing(listingID string) {
	// get the statement from the config
	//query := Config.SQLQueries["update_listing"]
	// prepare the arguments using the lst paramater

	var tag Listing

	//not really sure what we are updating or how we are gettin the info
	//so essentially just change this test var to whatever is getting updated
	var test = "I just changed it"

	// Execute the query

	//old query that just selects the book based on id
	//err := db.QueryRow("SELECT id, title, description, isbn, price, category, seller, listing_password, status FROM Listings where id = ?", listingID).Scan(&tag.ID, &tag.Title, &tag.Description, &tag.ISBN, &tag.Price, &tag.Category, &tag.SellerName, &tag.ListingPassword, &tag.Status)

	//actual query to seupdate a title
	err := db.QueryRow("UPDATE Listings SET title = ? WHERE id = ?", test, listingID).Scan(&tag.Title)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	log.Println(tag)
	log.Println(tag.Title)

}

//func SelectActive() []Listing {
//}

// TODO
//func SelectPrivate(password string) Listing {
//}

// TODO
func Search(keyword string) {

	log.Println(keyword)

	results, err := db.Query("SELECT * FROM Listings WHERE title LIKE '%?%'OR description LIKE '%?%'OR isbn LIKE '%?%'OR category LIKE '%?%'OR seller LIKE '%?%'", keyword, keyword, keyword, keyword, keyword)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var book Listing
		//for each row, scan the result into our Listing obj
		err = results.Scan(&book.ID, &book.Title, &book.Description, &book.ISBN, &book.Price, &book.Category, &book.SellerName, &book.ListingPassword, &book.Status)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(book.Title)

	}

}
