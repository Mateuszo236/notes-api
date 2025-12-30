package main

import (
	"fmt"
	"net/http"
	"log"
)

func healthHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK\n")
}

func main(){
	fmt.Println("Hello world")
	
	http.HandleFunc("/health", healthHandler)

	if err := http.ListenAndServe(":8080" , nil); err != nil {
		log.Fatal(err)
	}
	
}
