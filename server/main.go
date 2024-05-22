package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	posts := make(map[int]SharkAttack)
	jsonData, jsonerr := readJSON()
	if jsonerr != nil {
		log.Println(jsonerr)
	} else {
		for i := 0; i < len(jsonData); i++ {
			posts[i] = jsonData[i]
		}
	}

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var attack SharkAttack
			err := json.NewDecoder(r.Body).Decode(&attack)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			posts[len(posts)] = attack
			fmt.Fprintf(w, "Post added\n")
		}
	})
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			for key, value := range posts {
				fmt.Fprintf(w, "%d: %v\n", key, value)
			}
		}
	})
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			var key int
			err := json.NewDecoder(r.Body).Decode(&key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			delete(posts, key)
			fmt.Fprintf(w, "Post deleted\n")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
