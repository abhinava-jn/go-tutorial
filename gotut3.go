package main

import("fmt"
	   "net/http")

func index_handler(w http.ResponseWriter,r *http.Request) {
 fmt.Fprintf(w,"whoa, Go is neat!")
}

func about_handler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"number 1")
}


func abhinav_handler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"number 2")
}


func king_handler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"number 3")
}
func main() {
	http.HandleFunc("/",index_handler)
	http.HandleFunc("/about/",about_handler)
	http.HandleFunc("/abhinav/",abhinav_handler)
	http.HandleFunc("/about/king",king_handler)
	http.ListenAndServe(":8000",nil)
}

