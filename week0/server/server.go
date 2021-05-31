package main

import (
	"fmt"
	"log"
	"net/http"
)


func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hi there, I love %s %s",r.URL.Path,r.URL.Path[1:])
}


func main() {
	http.HandleFunc("/",index)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
