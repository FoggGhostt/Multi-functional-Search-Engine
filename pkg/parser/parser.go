package parser

import (
	"path/filepath"
	"sync"
)

func ParseFile(file_path string) (*sync.Map, error) {
	format := filepath.Ext(file_path)
	var sync_map_ptr *sync.Map
	var err error
	switch format {
	case ".txt":
		sync_map_ptr, err = Parse_txt_File(file_path)
	}
	return sync_map_ptr, err
}
