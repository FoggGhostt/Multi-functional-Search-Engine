package parser

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

const BLOCK_SIZE = 2 << (10) //  пока не очень понимаю, как его формировать
const UTF8_START_BYTE = 0x80
const MIDDLE_UTF8_BYTE_SIZE = 6
const START_OF_MIDDLE_UTF8_BYTE = 0b10
const ERROR_CHANEL_SIZE = 100

func searchFile(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); err != nil {
		return false, err
	}
	return true, nil
}

func count_1_in_begining(b byte) int {
	if b&UTF8_START_BYTE == 0 {
		return 1
	}
	result := 0
	start_num := byte(1 << 7)
	for b&start_num != 0 {
		start_num = start_num >> 1
		result++
	}
	return result
}

func is_start_byte(b byte) bool {
	return b>>MIDDLE_UTF8_BYTE_SIZE != START_OF_MIDDLE_UTF8_BYTE
}

func Parse_txt_File(filePath string) (*sync.Map, error) {
	does_exists, err := searchFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error while file searching")
	}
	if !does_exists {
		return nil, fmt.Errorf("file %s didn't found", filePath)
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error while file opening")
	}
	buffer := make([]byte, BLOCK_SIZE)
	var sync_map sync.Map
	var wg sync.WaitGroup
	errCh := make(chan error, ERROR_CHANEL_SIZE)
	undecoded_tail_len := 0
	defer file.Close()
	// buf_reader := bufio.NewReader(file)  оно не работает,
	// seek некорректно сдвигает его границу (даже когда делаю это для файла а потом создаю новый ридер)
	for {
		// byte_read_count, err := buf_reader.Read(buffer)
		byte_read_count, err := file.Read(buffer)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error while file reading")
		}
		for i := 0; i < byte_read_count; i++ {
			if is_start_byte(buffer[byte_read_count-i-1]) {
				undecoded_tail_len = i + 1 //  Находим место, где начинается первый байт символа unicode
				break
			}
		}
		if undecoded_tail_len != 0 && undecoded_tail_len != count_1_in_begining(buffer[byte_read_count-undecoded_tail_len]) {
			if _, err := file.Seek(-int64(undecoded_tail_len), io.SeekCurrent); err == nil {
				// buf_reader = bufio.NewReader(file)
			} else {
				return nil, fmt.Errorf("error while file reading")
			}
		} else {
			undecoded_tail_len = 0
		}
		dataCopy := string(buffer[:byte_read_count-undecoded_tail_len])
		wg.Add(1)
		go func(data string) {
			defer wg.Done()
			err := Tokenize(data, &sync_map)
			if err != nil {
				errCh <- err
			}
		}(dataCopy)
	}
	wg.Wait()
	select {
	case err := <-errCh:
		return nil, err
	default:
		return &sync_map, nil
	}
}
