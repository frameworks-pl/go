package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"	
	"io/ioutil"
)

type Config struct {
	//Variable names must start from capital letter to be exported
	DbName string `json:"dbname"`
	DbPass string `json:"dbpass"`
}

func main() {

	configPath := "./config.json"
	jsonBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	s := string(jsonBytes[:])
	fmt.Printf(s)

	var config Config;
	err = json.Unmarshal(jsonBytes, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

    http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, config.DbName)
    })

    fmt.Printf("Server running (port=8081), route: http://localhost:8081/helloworld\n")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatal(err)
    }
}
