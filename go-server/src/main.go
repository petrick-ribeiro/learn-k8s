package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Listening at port 3000\n")

	http.Handle("/", http.FileServer(http.Dir("./static/")))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
