package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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
	fmt.Println(lst.Title) //take out later
	lst.ListingPassword = hash
	log.Println("CAN INSERT")
	log.Println(lst)
	lst.Status = "active"
	newid, err := lst.Insert()
	//w.Header().Set("new_id", string(newid))
	if err != nil {
		// write error
		w.WriteHeader(503)
		fmt.Fprintf(w, "%s", "Internal Server Error")
	} else {
		// good, send good response
		w.WriteHeader(200)
		fmt.Fprintf(w, "%d", newid)
	}
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
	vars := mux.Vars(r)
	selectedID := vars["listing_id"]

	intID, err := strconv.Atoi(selectedID)

	if err != nil {
		// bad id
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if intID < 1 {
		// bad id
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// get listing info from DB using ID
	selectedListing := SelectPublicListingDetails(intID)

	if len(selectedListing) == 0 {
		// 404
		w.WriteHeader(http.StatusNotFound)
	} else {
		// good
		RenderSingleListingTemplate(w, "listing.html", selectedListing[0])
	}

	// call RenderSingleListingTemplate with tmpl=listing.html

}

// UpdateListingGETHandler POST T7
// serves the update listing page using the template update.html
// similar to PublicListingDataHandler
func UpdateListingGETHandler(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "./pages/update.html")
	vars := mux.Vars(r)
	selectedID := vars["listing_id"]
	ID, err := strconv.Atoi(selectedID)

	if err != nil {
		//ahh
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// get listing info from DB using ID
	selectedListing := SelectPublicListingDetails(ID)

	if len(selectedListing) == 0 {
		// 404
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		// good
		RenderSingleListingTemplate(w, "update.html", selectedListing[0])
	}
}

// UpdateListingPOSTHandler POST T7
func UpdateListingPOSTHandler(w http.ResponseWriter, r *http.Request) {
	// use the lines below to get the data from URL {listing_id}
	vars := mux.Vars(r)
	listingId := vars["listing_id"]
	intID, err := strconv.Atoi(listingId)
	log.Print(listingId)
	var lst Listing
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(string(bytes))
	err = json.Unmarshal(bytes, &lst)

	//hash, err := HashPassword(lst.ListingPassword)
	fmt.Println(lst.Title) //take out later
	//lst.ListingPassword = hash
	log.Println(lst)

	if err != nil {
		fmt.Fprintf(w, "fail")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		passwordIn := lst.ListingPassword

		//call SelectPassword get the hash of the Listing with intID.
		hashToCheckAgainst := SelectPassword(intID)
		//log.Printf("hashToCheckAgainst(the result of SelectPassword) is %s \n", hashToCheckAgainst)

		// check password is correct
		passwordIsCorrect := ComparePassword(passwordIn, hashToCheckAgainst)

		// if password is correct ... go to update page?
		if passwordIsCorrect == true {

			log.Println("update complete")
			UpdateListing(intID, lst.Title, lst.Description, lst.ISBN, lst.Price, lst.Category, lst.SellerName)

			w.WriteHeader(200)
			fmt.Fprintf(w, "%t", true)

		} else { // wrong password

			w.WriteHeader(200)
			fmt.Fprintf(w, "%t", false)
		}
	}

}

// DeleteListingHandler POST T8
func DeleteListingHandler(w http.ResponseWriter, r *http.Request) {
	//get the listing ID
	vars := mux.Vars(r)
	selectedID := vars["listing_id"]
	intID, err := strconv.Atoi(selectedID)
	//log.Printf("inside the deleteLIstinghandler. intID is %d \n", intID)

	var lst Listing
	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}
	//log.Println("Printing string(btyes) in delete handler:")
	log.Println(string(bytes))

	err = json.Unmarshal(bytes, &lst)
	log.Println(lst)
	//log.Printf("The password in the delete handler is %s", lst.ListingPassword)

	if err != nil {
		fmt.Fprintf(w, "fail")
	} else {
		passwordIn := lst.ListingPassword

		//call SelectPassword get the hash of the Listing with intID.
		hashToCheckAgainst := SelectPassword(intID)
		//log.Printf("hashToCheckAgainst(the result of SelectPassword) is %s \n", hashToCheckAgainst)

		// check password is correct
		passwordIsCorrect := ComparePassword(passwordIn, hashToCheckAgainst)

		//log.Printf("the result of ComparePassword is %t \n", passwordIsCorrect)

		// if password is correct, delete listing with id
		if passwordIsCorrect == true {
			//log.Println("password matches. It should delete listing\n")

			DeleteListing(intID) //delete the listing from table

			w.WriteHeader(200)
			fmt.Fprintf(w, "%t", true)
			log.Println("CAN DELETE \n")

		} else { // wrong password
			w.WriteHeader(200)
			fmt.Fprintf(w, "%t", false)
		}
	}
}

// ActiveListingsHandler GET T9
func ActiveListingsHandler(w http.ResponseWriter, r *http.Request) {
	// uses the DB to get all active listings
	activeListings := SelectActive()
	// returns a json-encoded array of listings
	js, err := json.Marshal(activeListings)
	if err != nil {
		// TODO not panic
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// SearchListingsHandler POST T10
func SearchListingsHandler(w http.ResponseWriter, r *http.Request) {
	// uses get keyword from body
	// execute search query on SQL
	// return json-encoded array of list
	var keyword string
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(bytes))

	keyword = string(bytes)
	res1 := strings.SplitAfter(keyword, ":\"")
	withQ := strings.TrimSuffix(res1[1], "}")
	cleanedString := strings.TrimSuffix(withQ, "\"")
	searchResults := Search(cleanedString)

	js, err := json.Marshal(searchResults)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// PrivateListingDetailsHandler POST T11
func PrivateListingDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// use the lines below to get the data from URL {listing_id}
	vars := mux.Vars(r)

	lstID, err := strconv.Atoi(vars["listing_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var lst Listing
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bytes, &lst)

	if err != nil {

		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", "false_json_unmarshal")

	} else {
		dbPassword := SelectPassword(lstID)
		passwordIsCorrect := ComparePassword(lst.ListingPassword, dbPassword)

		if passwordIsCorrect {
			// select private listing details
			lst := SelectPrivate(lstID)
			js, err := json.Marshal(lst)
			if err != nil {
				w.WriteHeader(200)
				fmt.Fprintf(w, "%s", "false_json")

			} else {
				w.WriteHeader(200)
				w.Header().Set("Content-Type", "application/json")
				w.Write(js)
			}
		} else {
			w.WriteHeader(200)
			fmt.Fprintf(w, "%s", "false_password")
		}
	}
}

// PurchaseListingHandler POST T12
func PurchaseListingHandler(w http.ResponseWriter, r *http.Request) {
	// use the lines below to get the data from URL {listing_id}
	vars := mux.Vars(r)
	log.Println("listing.purchase")

	var lst Listing
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &lst)
	lst.Status = "purchased"
	lst.ID, err = strconv.Atoi(vars["listing_id"])

	if err != nil {
		fmt.Fprintf(w, "fail")

	} else {
		if lst.ID > -1 {
			PurchaseListing(lst.Buyer, lst.BillingInfo, lst.ShippingInfo, lst.ID)
			fmt.Fprintf(w, "true")
		} else {
			fmt.Fprintf(w, "false")
		}
	}

}
