package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	body := r.PostFormValue("topic")
	fmt.Fprintf(w,body)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8803", nil)
}
