package search

import (
	"context"
	"fmt"
	"math"
	"search-engine/pkg/models"
	"search-engine/pkg/mongodb"
)

// Функция возвращает вектор и матрицу, вектор - вектор tf-idf для поискового запроса,
// а матрица - матрица tf-idf для документов, в которых встречаются токены из нашего запроса
// Также для удобства возвращаем сплошной список документов для последующей сортировки запросов
func Create_TF_IDF_Matrix(req_tokens []string, token_map map[string]int, rel_docs_info []models.TokenInfo) ([]float64, [][]float64, []string, error) {
	filePathMap := make(map[string]bool)
	filePaths := make([]string, 0) //  Создаем матрицу
	for _, tokenInfo := range rel_docs_info {
		for _, occureInfo := range tokenInfo.Occures {
			_, ok := filePathMap[occureInfo.FilePath]
			if !ok {
				filePaths = append(filePaths, occureInfo.FilePath)
				filePathMap[occureInfo.FilePath] = true
			}
		}
	}

	invert_token_map := make(map[string]int) // Создаем обратное соответствие (токен - номер координаты векторов)
	for i, token := range req_tokens {
		invert_token_map[token] = i
	}
	matrix := make([][]float64, len(filePaths)) //  Создаем матрицу)
	for i := range filePaths {
		fmt.Println(filePaths[i])
		matrix[i] = make([]float64, len(req_tokens))
	}
	req_vec_idf := make([]float64, len(req_tokens))    //  Вектор счетчиков вхождения токенов запроса в коллекцию документов
	req_vec_tf_idf := make([]float64, len(req_tokens)) //  Итоговый tf-idf вектор запроса

	files_lengths := make([]float64, len(filePaths))

	db, err := mongodb.GetDB()
	if err != nil {
		return nil, nil, nil, err
	}

	for i := range filePaths { //  Пробегаемся по документам, токенизируем их и заполняем матрицу tf-idf
		docInfo, err := db.GetFileIndex(context.Background(), filePaths[i])
		if err != nil {
			return nil, nil, nil, err
		}
		for _, tokInfo := range docInfo.Tokens {
			files_lengths[i] += float64(tokInfo.OccureCount) // Считаем общее количество токенов в документе
			tokenIndex, ok := invert_token_map[tokInfo.Token]
			if ok {
				req_vec_idf[tokenIndex] += 1.0
				matrix[i][tokenIndex] += float64(tokInfo.OccureCount)
			}
		}
	}
	for i := range req_tokens {
		req_vec_tf_idf[i] = float64(token_map[req_tokens[i]]) * math.Log((float64(len(filePaths))/req_vec_idf[i])+1.0)
	}
	for i := range req_tokens { //  Досчитали метрики tf-idf для матрицы
		for j := range filePaths {
			matrix[j][i] /= files_lengths[j]
			// matrix[j][i] *= math.Log(float64(len(filePaths)) / req_vec_idf[i])
		}
	}
	return req_vec_tf_idf, matrix, filePaths, nil
}
