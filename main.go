package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var responseData interface{}
	json.NewDecoder(r.Body).Decode(&responseData)
	fmt.Println(responseData)
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/event", handler)
	log.Fatal(http.ListenAndServe(":3030", nil))
}
