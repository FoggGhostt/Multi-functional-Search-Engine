package indexer

import (
	"fmt"
	"os"
	"search-engine/pkg/parser"
	"sync"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const FILE_BLOCK_LIMIT = 2 << 23
const ERROR_CHANEL_SIZE = 100

// type InvIndex []TokenInfo

// type TokenInfo struct {
//     Token string
//     Occures []OccureInfo
// }

// type OccureInfo struct {
// 	FilePath string
// 	OccureCount int64
// }

type InvIndex []TokenInfo

type TokenInfo struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Token   string             `bson:"token"`
	Occures []OccureInfo       `bson:"occures"`
}

type OccureInfo struct {
	FilePath    string `bson:"file_path"`
	OccureCount int64  `bson:"occure_count"`
}

func IndexFiles(filePaths []string) error {
	var syncInvertedIndex sync.Map
	var mtx sync.Mutex
	var wg sync.WaitGroup
	errCh := make(chan error, ERROR_CHANEL_SIZE)

	curFileBlockSize := 0
	curFileBlock := make([]string, 0)
	for i, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("cant open the file")
		}
		fi, err := file.Stat()
		if err != nil {
			return fmt.Errorf("cant get stat of the file")
		}
		curFileBlockSize += int(fi.Size())
		curFileBlock = append(curFileBlock, filePath)
		file.Close()
		if curFileBlockSize > FILE_BLOCK_LIMIT || i == len(filePaths)-1 {
			fileBlockCopy := make([]string, len(curFileBlock))
			copy(fileBlockCopy, curFileBlock)
			curFileBlock = make([]string, 0)
			curFileBlockSize = 0

			wg.Add(1)
			go func(filesBlock []string) { //  Одна горутина берет пак файлов, каждый из них парсит и потом заносит все изменения в общую мапу syncInvertedIndex
				defer wg.Done()
				for _, filePath := range filesBlock {
					sync_map, err := parser.ParseFile(filePath)
					if err != nil {
						errCh <- err
					}

					sync_map.Range(func(key, value any) bool {
						strKey, isCorrectType := key.(string)
						if !isCorrectType {
							errCh <- fmt.Errorf("incorrect type in sync_map")
							return false
						}
						intValue, isCorrectType := value.(int64)
						if !isCorrectType {
							errCh <- fmt.Errorf("incorrect type in sync_map")
							return false
						}

						mtx.Lock()

						syncValue, is_inside := syncInvertedIndex.Load(key)

						if is_inside {
							tokenInfo, isCorrectType := syncValue.(TokenInfo)
							if !isCorrectType {
								errCh <- fmt.Errorf("incorrect type in sync_map")
								return false
							}
							tokenInfo.Occures = append(tokenInfo.Occures, OccureInfo{FilePath: filePath, OccureCount: intValue})
							syncInvertedIndex.Store(strKey, tokenInfo)
						} else {
							tokenInfo := TokenInfo{Token: strKey, Occures: []OccureInfo{OccureInfo{FilePath: filePath, OccureCount: intValue}}}
							syncInvertedIndex.Store(strKey, tokenInfo)
						}

						mtx.Unlock()

						return true
					})
				}
			}(fileBlockCopy)
		}
	}
	wg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
	}

	// DataBase interaction ?
	mongoURI := "mongodb://localhost:27017"

	return nil
}
