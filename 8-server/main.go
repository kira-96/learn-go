package main

import (
	"fmt"
	"log"
	"net/http"
)

type yuan float32

func (t yuan) String() string {
	return fmt.Sprintf("￥%.2f", t)
}

type database map[string]yuan

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, &price)
	}
}

func main() {
	db := database{"Go T-Shirt": 99, "Go Jacket": 199}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
