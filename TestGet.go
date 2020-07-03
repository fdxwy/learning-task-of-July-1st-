package main

import(
	"fmt"
	"net/http"
)

func getParameter(w http.ResponseWriter,r *http.Request){
	parameter1 := r.URL.Query().Get("aaa")
	parameter2 := r.URL.Query().Get("bbb")
	
	fmt.Fprintln(w,parameter1)
	fmt.Fprintln(w,parameter2)
}

func main(){
	http.HandleFunc("/",getParameter)
	http.ListenAndServe(":8803",nil)

}
