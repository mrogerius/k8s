package main

import "net/http"

func main() {
	http.HandlerFunc("/", Home)
	http.ListenAndServe(":3000", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Ola mundo do K8S</h1>"))
}
