package search

import (
	"context"
	"fmt"
	"math"
	"os"
	"search-engine/pkg/models"
	"search-engine/pkg/mongodb"
	"sort"
)

const K_RELEVANT_DOC = 10

type VecInfo struct {
	vec      []float64
	filePath string
}

func findAngle(vec1, vec2 []float64) float64 {
	scalarProd := 0.0
	for i := 0; i < len(vec1); i++ {
		scalarProd += vec1[i] * vec2[i]
	}
	return scalarProd / (findVecMod(vec1) * findVecMod(vec2))
}

func findVecMod(vec []float64) float64 {
	sum := 0.0
	for i := range vec {
		sum += (vec[i] * vec[i])
	}
	return math.Sqrt(sum)
}

func Search(req string) ([]string, error) {
	reqTokensWithMatch, err := TokenizeSearchRequest(req)
	reqTokensNoMatch := make([]string, 0)
	if err != nil {
		return nil, err
	}
	relevant_docs_info := make([]models.TokenInfo, 0)
	token_map := make(map[string]int)

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017" // Значение по умолчанию для локального запуска без Docker
	}

	cnf := mongodb.DefaultConfig()
	cnf.DbName = "InvertIndex"

	db, err := mongodb.Init(mongoURI, cnf)
	if err != nil {
		return nil, err
	}

	for _, token := range reqTokensWithMatch {
		if _, ok := token_map[token]; ok {
			token_map[token]++
			continue
		}
		token_map[token] = 1
		reqTokensNoMatch = append(reqTokensNoMatch, token)
		tokenInfo, err := db.FindRelDocs(context.Background(), token)
		if err != nil {
			return nil, err
		}
		if tokenInfo != nil {
			relevant_docs_info = append(relevant_docs_info, *tokenInfo)
		}
	}

	req_vec_tf_idf, matrix, filePaths, err := Create_TF_IDF_Matrix(reqTokensNoMatch, token_map, relevant_docs_info)
	if err != nil {
		return nil, err
	}

	fmt.Println(matrix[0][0], filePaths[0], matrix[1][0], filePaths[1])

	docVecsToSort := make([]VecInfo, 0)
	for i := range filePaths {
		docVecsToSort = append(docVecsToSort, VecInfo{vec: matrix[i], filePath: filePaths[i]})
	}

	sort.Slice(docVecsToSort, func(i, j int) bool {
		cos_i := findAngle(docVecsToSort[i].vec, req_vec_tf_idf)
		cos_j := findAngle(docVecsToSort[j].vec, req_vec_tf_idf)
		fmt.Println(cos_i, cos_j)
		if cos_i == cos_j {
			return findVecMod(docVecsToSort[i].vec) > findVecMod(docVecsToSort[j].vec)
		}
		return findAngle(docVecsToSort[i].vec, req_vec_tf_idf) > findAngle(docVecsToSort[j].vec, req_vec_tf_idf)
	})

	fmt.Println(docVecsToSort[0].filePath, docVecsToSort[1].filePath)

	searchResult := make([]string, len(docVecsToSort))

	for i := range docVecsToSort {
		searchResult[i] = docVecsToSort[i].filePath
	}

	return searchResult, nil
}
