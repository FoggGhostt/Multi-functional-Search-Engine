package parser

import (
	"fmt"
	"mime"
	"path/filepath"
	"strings"
	"sync"
)

func ParseFile(file_path string) (*sync.Map, error) {
	ext := filepath.Ext(file_path)
	mimeType := mime.TypeByExtension(ext)
	mimeType = strings.Split(mimeType, ";")[0]

	var sync_map_ptr *sync.Map
	var err error

	fmt.Println(mimeType)

	if mimeType == "text/plain" {
		sync_map_ptr, err = Parse_txt_File(file_path)
	} else if mimeType == "application/pdf" {
		// sync_map_ptr, err = Parse_pdf_file(file_path)
	}
	return sync_map_ptr, err
}
