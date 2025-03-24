package indexer

import (
	"context"
	"fmt"
	"os"
	"search-engine/pkg/models"
	"search-engine/pkg/mongodb"
	"search-engine/pkg/parser"
	"sync"
)

const FILE_BLOCK_LIMIT = 2 << 23
const ERROR_CHANEL_SIZE = 100

func IndexFiles(filePaths []string) error {
	var syncInvertedIndex sync.Map
	var mtx sync.Mutex
	var wg sync.WaitGroup
	errCh := make(chan error, ERROR_CHANEL_SIZE)

	curFileListSize := 0
	curFileList := make([]string, 0)
	for i, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("cant open the file")
		}
		fi, err := file.Stat()
		if err != nil {
			return fmt.Errorf("cant get stat of the file")
		}
		curFileListSize += int(fi.Size())
		curFileList = append(curFileList, filePath)
		file.Close()
		if curFileListSize > FILE_BLOCK_LIMIT || i == len(filePaths)-1 {
			fileListCopy := make([]string, len(curFileList))
			copy(fileListCopy, curFileList)
			curFileList = make([]string, 0)
			curFileListSize = 0

			wg.Add(1)
			go func(filesList []string) { //  Одна горутина берет пак файлов, каждый из них парсит и потом заносит все изменения в общую мапу syncInvertedIndex
				defer wg.Done()
				for _, filePath := range filesList {
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
						defer mtx.Unlock()

						syncValue, is_inside := syncInvertedIndex.Load(key)

						if is_inside {
							tokenInfo, isCorrectType := syncValue.(models.TokenInfo)
							if !isCorrectType {
								errCh <- fmt.Errorf("incorrect type in sync_map")
								return false
							}
							tokenInfo.Occures = append(tokenInfo.Occures, models.OccureInfo{
								FilePath:    filePath,
								OccureCount: intValue,
							})
							syncInvertedIndex.Store(strKey, tokenInfo)
						} else {
							tokenInfo := models.TokenInfo{
								Token: strKey,
								Occures: []models.OccureInfo{
									{FilePath: filePath, OccureCount: intValue},
								},
							}
							syncInvertedIndex.Store(strKey, tokenInfo)
						}

						return true
					})
				}
			}(fileListCopy)
		}
	}
	wg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
	}

	invertIndex := make([]models.TokenInfo, 0) //  Обратный индекс для хранения в базе данных

	syncInvertedIndex.Range(func(key, value any) bool {
		tokenInfo, isCorrectType := value.(models.TokenInfo)
		if !isCorrectType {
			// IDK what I should do there)) Не хочу создавать переменную для ошибки - коряво как-то, хелп)
			return false
		}
		invertIndex = append(invertIndex, tokenInfo)
		return true
	})

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = os.Getenv("LOCAL_MONGO_URI") // Значение по умолчанию для локального запуска без Docker
	}

	cnf := mongodb.DefaultConfig()
	cnf.DbName = "InvertIndex"

	db, err := mongodb.Init(mongoURI, cnf)
	if err != nil {
		return err
	}

	err = db.UpsertTokenInfos(context.Background(), invertIndex) // добавляю индекс в базу
	if err != nil {
		return err
	}

	return nil
}
