package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var (
	Creds  CredsStruct
	Config ConfigStruct
)

// HashPassword hashes a password are returns the hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// ComparePassword compares a password and a hash, returns true if match
func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// DecodeJSONBody takes a reader and retuns a generic JSON
func DecodeJSONBody(body io.Reader) GenericJSON {
	log.Println("decodejson.decode.start")
	// read the bodt
	bytedata, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println("decodejson.read.fail")
		panic(err)
	}

	// delcare buffer
	var jsdata GenericJSON

	// unmarshal the data into the struct
	err = json.Unmarshal(bytedata, &jsdata)

	if err != nil {
		log.Println("decodejson.decode.fail")
		panic(err)
	}
	log.Println("decodejson.decode.success")
	return jsdata
}

func loadConfigs() {
	log.Println("loadconfig.config.start_load")
	// Get file handler
	cfgFle, err := os.Open("config.json")
	if err != nil {
		log.Println("loadconfig.config.fail_handle")
		panic(err)
	}
	// defer closing the file to the end of the func
	defer cfgFle.Close()

	// read bytes from file handler
	bytes, err := ioutil.ReadAll(cfgFle)
	if err != nil {
		log.Println("loadconfig.config.fail_load")
		panic(err)
	}
	// use GoLang's json function to load
	// the config values into the config struct
	json.Unmarshal(bytes, &Config)
	log.Println("loadconfig.config.success_load")

	log.Println("loadconfig.creds.start_load")
	// get handler to cred file
	credFle, err := os.Open("creds.json")
	if err != nil {
		log.Println("loadconfig.creds.fail_handle")
		panic(err)
	}
	// defer closing file to end of func
	defer credFle.Close()

	// read bytes in cred file
	bytes, err = ioutil.ReadAll(credFle)
	if err != nil {
		log.Println("loadconfig.creds.fail_load")
		panic(err)
	}
	// load the values into the creds struct
	json.Unmarshal(bytes, &Creds)
	log.Println("loadconfig.creds.success_load")
}

func main() {
	// start server
	log.Println("server.init")
	// load configs
	loadConfigs()
	// initialize DB connection
	InitDB()
	// Run startup queries
	SQLInits()

	// http
	rootRouter := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static"))
	rootRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// JSON Endpoints
	rootRouter.HandleFunc("/create_listing", CreateListingHandler)
	rootRouter.HandleFunc("/listing/{listing_id}", PublicListingDataHandler)
	rootRouter.HandleFunc("/listing/{listing_id}/update", UpdateListingHandler)
	rootRouter.HandleFunc("/delete", DeleteListingHandler)
	rootRouter.HandleFunc("/active", ActiveListingsHandler)
	rootRouter.HandleFunc("/search", SearchListingsHandler)
	rootRouter.HandleFunc("/private_details", PrivateListingDetailsHandler)
	rootRouter.HandleFunc("/listing/{listing_id}/purchase", PurchaseListingHandler)

	rootRouter.HandleFunc("/", IndexHandler)

	// register the router
	http.Handle("/", rootRouter)

	// start listening
	http.ListenAndServe(fmt.Sprintf("%s:%s", Config.WebHost, Config.WebPort), nil)
}
