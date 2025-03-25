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
	if err := r.ParseMultipartForm(100 << 20); err != nil { // лимит в 100 мб
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

		fullFilePath := "./files/" + fileHeader.Filename // папка с загружаемыми файлами

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

	indexer.IndexFiles(files_to_index)
}

func downloadFileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	filePath := r.URL.Query().Get("path")
	if filePath == "" {
		http.Error(w, "No path", http.StatusBadRequest)
		return
	}

	f, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		http.Error(w, "No file info", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=\""+path.Base(filePath)+"\"")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	io.Copy(w, f)
}

func API_Init() {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/api/search", searchHandler)
	http.HandleFunc("/download", downloadFileHandler)
	http.ListenAndServe(":8080", nil)
}
