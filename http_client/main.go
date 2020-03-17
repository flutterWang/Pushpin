package main

import "net/http"

func toDo(w http.ResponseWriter, req *http.Request) {

}

func main() {
	http.HandleFunc("/", toDo)
	http.ListenAndServe(":8080", nil)
}
