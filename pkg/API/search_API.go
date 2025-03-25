package API

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"search-engine/pkg/indexer"
	"search-engine/pkg/search"
)

type FileResult struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Sentence string `json:"sentence"`
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // CORS
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "query is required", http.StatusBadRequest)
		return
	}

	results, err := search.Search(query)

	if err != nil || len(results) == 0 {
		return
	}

	modified_results := make([]FileResult, 0)
	for _, filepath := range results {
		modified_results = append(modified_results, FileResult{Name: path.Base(filepath), Path: filepath})
	}

	json.NewEncoder(w).Encode(modified_results)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("Start downloading files")
	if err := r.ParseMultipartForm(100 << 20); err != nil {
		http.Error(w, "Huge files size", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		http.Error(w, "No files", http.StatusBadRequest)
		return
	}

	if err := os.MkdirAll("./files", os.ModePerm); err != nil {
		http.Error(w, "Cant creat dir", http.StatusInternalServerError)
		return
	}

	files_to_index := make([]string, 0)

	for _, fileHeader := range files { // скачиваем каждый файл
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Cant open files", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		fullFilePath := "./files/" + fileHeader.Filename

		dst, err := os.Create(fullFilePath)
		if err != nil {
			http.Error(w, "Cant create file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err = io.Copy(dst, file); err != nil {
			http.Error(w, "Cant download file", http.StatusInternalServerError)
			return
		}

		files_to_index = append(files_to_index, fullFilePath)
	}

	w.WriteHeader(http.StatusOK)

	go indexer.IndexFiles(files_to_index)
}

func API_Init() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/api/search", searchHandler)
	http.ListenAndServe(":8080", nil)
}
