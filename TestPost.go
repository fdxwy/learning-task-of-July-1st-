package main

import (
"fmt"
"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	method :=r.Method
	body := r.Body
	header := r.Header
	topic := r.PostFormValue("topic")
	fmt.Fprintln(w, method)
	fmt.Fprintln(w, body)
	fmt.Fprintln(w, header)
	fmt.Fprintf(w,topic)
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8803", nil)
}
