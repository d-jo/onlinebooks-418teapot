package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// IndexHandler serves the static index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./pages/index.html")
	log.Println("indexhandler.index.served")
}

// CreateListingPOSTHandler POST T5
// TODO finish this
func CreateListingPOSTHandler(w http.ResponseWriter, r *http.Request) {
	//jsdata := DecodeJSONBody(r.Body)
	//log.Println(jsdata)
	var lst Listing
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(bytes))
	err = json.Unmarshal(bytes, &lst)
	log.Println(lst)
	hash, err := HashPassword(lst.ListingPassword)
	lst.ListingPassword = hash
	log.Println("CAN INSERT")
	log.Println(lst)
	//lst.Insert()
}

// CreateListingGETHandler POST T5
// serves the CreateListing page
func CreateListingGETHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./pages/create.html")
}

// PublicListingDataHandler GET T6
// uses template to serve public listing data page
func PublicListingDataHandler(w http.ResponseWriter, r *http.Request) {
	// use the lines below to get the data from URL {listing_id}
	//vars := mux.Vars(r)
	//vars["listing_id"]

	// get listing info from DB using ID
	// call RenderSingleListingTemplate with tmpl=listing.html

}

// UpdateListingGETHandler POST T7
// serves the update listing page using the template update.html
// similar to PublicListingDataHandler
func UpdateListingGETHandler(w http.ResponseWriter, r *http.Request) {
	// use the lines below to get the data from URL {listing_id}
	//vars := mux.Vars(r)
	//vars["listing_id"]

	// get the listing details from database
	// use RenderSingleListingTemplate with tmpl=update.html

}

// UpdateListingPOSTHandler POST T7
func UpdateListingPOSTHandler(w http.ResponseWriter, r *http.Request) {
	// use the lines below to get the data from URL {listing_id}
	//vars := mux.Vars(r)
	//vars["listing_id"]

	// decode the body
	// use DB update the listing if password is correct

}

// DeleteListingHandler POST T8
func DeleteListingHandler(w http.ResponseWriter, r *http.Request) {
	// use the lines below to get the data from URL {listing_id}
	//vars := mux.Vars(r)
	//vars["listing_id"]

	// decode password from body
	// check password is correct
	// if password is correct, delete listing with id

}

// ActiveListingsHandler GET T9
func ActiveListingsHandler(w http.ResponseWriter, r *http.Request) {

	// uses the DB to get all active listings
	// returns a json-encoded array of listings

}

// SearchListingsHandler POST T10
func SearchListingsHandler(w http.ResponseWriter, r *http.Request) {

	// uses get keyword from body
	// execute search query on SQL
	// return json-encoded array of listings
}

// PrivateListingDetailsHandler POST T11
func PrivateListingDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// use the lines below to get the data from URL {listing_id}
	//vars := mux.Vars(r)
	//vars["listing_id"]

	// checks the password
	// if match, get the private details (buyer info)

}

// PurchaseListingHandler POST T12
func PurchaseListingHandler(w http.ResponseWriter, r *http.Request) {
	// use the lines below to get the data from URL {listing_id}
	//vars := mux.Vars(r)
	//vars["listing_id"]

}
