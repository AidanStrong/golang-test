package main

import (
	"net/http"
)

func main() {

}

func simpleHttp() {
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":8081", nil)
}
