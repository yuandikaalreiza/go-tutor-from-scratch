package main

import (
	"encoding/json"
	"net/http"
)

type student struct {
	ID    string
	name  string
	grade int
}

var data = []student{
	student{"E001", "ethan", 21},
	student{"W001", "wick", 22},
	student{"B001", "bourne", 23},
	student{"B002", "bond", 23},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}

		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
