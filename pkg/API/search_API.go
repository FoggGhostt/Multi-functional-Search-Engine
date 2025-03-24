package API

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FileResult struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // CORS
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "query is required", http.StatusBadRequest)
		return
	}

	results := []FileResult{{Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}, {Name: query, Path: "/hehe"}}
	fmt.Println("req")

	json.NewEncoder(w).Encode(results)
}

func API_Init() {
	http.HandleFunc("/api/search", searchHandler)
	http.ListenAndServe(":8080", nil)
}
