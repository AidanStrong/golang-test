package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Play struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Year   int32  `json:"year"`
	Copies uint8  `json:"copies"`
}

var count int64 = 0

var play1 Play = Play{increment_id_count(), "Macbeth", 1623, 1}
var play2 Play = Play{increment_id_count(), "Hamlet", 1603, 2}
var play3 Play = Play{increment_id_count(), "Romeo and Juliet", 1597, 5}

var bookArray = []Play{play1, play2, play3}

// var bookArray = []Play{
// 	{
// 		Id: 1,
// 		Title: "Macbeth",
// 	},
// 	{}
// }

func main() {
	http.HandleFunc("/plays", playsHandler)
	http.HandleFunc("/plays/", playsHandler)
	//http.HandleFunc("/checkout/", checkoutHandler)
	fmt.Println("Running server on localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// //handler func
// func playsHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(bookArray)
// }

//handler func
func playsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idFromUrl, _ := strconv.ParseInt(getBookId(r.RequestURI), 10, 64)
	fmt.Println("After convert: ", idFromUrl)

	if idFromUrl == 0 {
		json.NewEncoder(w).Encode(bookArray)
		return
	}

	for _, book := range bookArray {
		if book.Id == idFromUrl {
			fmt.Fprintln(w, "There are", book.Copies, "copies left of", book.Title, "in the Shakespeare library")
		}
	}

	if idFromUrl > int64(len(bookArray)) {
		json.NewEncoder(w).Encode(bookArray)

	} else {
		json.NewEncoder(w).Encode(bookArray[idFromUrl-1])
	}

}

// func checkoutPlay(w http.ResponseWriter, r *http.Request) {

// }

// func returnPlay(w http.ResponseWriter, r *http.Request) {

// }

func getBookId(url string) string {
	split := strings.Split(url, "/")
	fmt.Println("SPLIT: ", split[len(split)-1])
	return split[len(split)-1]
}

func increment_id_count() int64 {
	count = count + 1
	return count
}
