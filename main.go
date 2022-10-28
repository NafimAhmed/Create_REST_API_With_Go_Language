package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Endpoint is hit")

}

type Article struct {
	Title   string `json"Title"`
	Desc    string `json"desc"`
	Content string `json"Content"`
}

type Articles []Article

func allArticle(w http.ResponseWriter, r *http.Request) {

	articles := Articles{
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
	}

	//fmt.Fprintf(w,"All Article End point is hit")
	json.NewEncoder(w).Encode(articles)
}

func testPostAllArticle(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintf(w, "Post End point hit")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var post Article
	json.Unmarshal(reqBody, &post)

	json.NewEncoder(w).Encode(post)

	//_, err := json.Marshal(post)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	//fmt.Println(string(newData))

	// 	// fmt.Fprintf(w, string(newData))
	// }

}

func main() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/article", allArticle).Methods("GET")
	myRouter.HandleFunc("/article", testPostAllArticle).Methods("POST")
	http.ListenAndServe(":8080", myRouter)

}
