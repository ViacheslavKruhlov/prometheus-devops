package main

import "net/http"

func main() {

	http.ListenAndServe(":18080", http.FileServer(http.Dir("./html")))
}
