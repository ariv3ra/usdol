package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.kimonolabs.com/api/2oowt0xq?apikey=2wCIm7WwnFkQJboi08Ttzg1UrBTHEBmq")

	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	var result map[string]interface{}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&result); err != nil {
		// handle error
		log.Println("This is an error")
	}

	// log.Println(result["results"].(map[string]interface{})["est_search"].(map[string]interface{})["est_name"], err)
	// log.Println(result["results"].(map[string]interface{})["est_search"], err)

	//log.Println(result, err)

	for _, item := range result["results"].(map[string]interface{})["est_search"].([]interface{}) {
		log.Println(item.(map[string]interface{})["sic"])
	}

}
