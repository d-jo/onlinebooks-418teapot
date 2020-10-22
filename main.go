package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var (
	Creds  CredsStruct
	Config ConfigStruct
)

func loadConfigs() {
	log.Println("loadconfig.config.start_load")
	// Get file handler
	cfgFle, err := os.Open("config.json")
	if err != nil {
		panic(err)
		log.Println("loadconfig.config.fail_handle")
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
		panic(err)
		log.Println("loadconfig.creds.fail_handle")
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
	log.Println("server.init")
	loadConfigs()
	InitDB()
	log.Println(Creds)
	log.Println(Config)
}
