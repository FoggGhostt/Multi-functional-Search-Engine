package search

import (
	"context"
	"fmt"
	"math"
	"testing"

	"search-engine/pkg/models"
	"search-engine/pkg/mongodb"
)

type FakeDB struct {
	docs       map[string]models.DocumentInfo
	tokenInfos map[string]models.TokenInfo
}

func (f *FakeDB) GetFileIndex(ctx context.Context, filePath string) (*models.DocumentInfo, error) {
	if doc, ok := f.docs[filePath]; ok {
		return &doc, nil
	}
	return &models.DocumentInfo{}, fmt.Errorf("doc not found")
}

func (f *FakeDB) FindRelDocs(ctx context.Context, token string) (*models.TokenInfo, error) {
	if tokenInfo, ok := f.tokenInfos[token]; ok {
		return &tokenInfo, nil
	}
	return nil, nil
}

func (f *FakeDB) UpsertDocInfos(ctx context.Context, docInfo models.DocumentInfo) error {
	return nil
}
func (f *FakeDB) UpsertTokenInfos(ctx context.Context, tokenInfos []models.TokenInfo) error {
	return nil
}

var originalGetDB = mongodb.GetDB

func TestCreateTFIDFMatrix(t *testing.T) {
	defer func() { mongodb.GetDB = originalGetDB }()

	tests := []struct {
		name           string
		reqTokens      []string
		tokenMap       map[string]int
		relDocsInfo    []models.TokenInfo
		fakeDocs       map[string]models.DocumentInfo
		expectedTFIDF  []float64
		expectedMatrix [][]float64
	}{
		{
			name:      "2x2",
			reqTokens: []string{"a", "b"},
			tokenMap:  map[string]int{"a": 2, "b": 3},
			relDocsInfo: []models.TokenInfo{
				{
					Token: "a",
					Occures: []models.OccureInfo{
						{FilePath: "doc1", OccureCount: 1},
					},
				},
				{
					Token: "b",
					Occures: []models.OccureInfo{
						{FilePath: "doc2", OccureCount: 1},
					},
				},
			},
			fakeDocs: map[string]models.DocumentInfo{
				"doc1": {
					Filepath: "doc1",
					Tokens: []models.IndexTokenInfo{
						{Token: "a", OccureCount: 10},
						{Token: "b", OccureCount: 5},
					},
				},
				"doc2": {
					Filepath: "doc2",
					Tokens: []models.IndexTokenInfo{
						{Token: "a", OccureCount: 2},
						{Token: "b", OccureCount: 8},
					},
				},
			},
			expectedMatrix: [][]float64{
				{10.0 / 15.0, 5.0 / 15.0},
				{2.0 / 10.0, 8.0 / 10.0},
			},
			expectedTFIDF: []float64{
				2 * math.Log(2),
				3 * math.Log(2),
			},
		},
		{
			name:      "3x4",
			reqTokens: []string{"a", "b", "c"},
			tokenMap:  map[string]int{"a": 1, "b": 2, "c": 3},
			relDocsInfo: []models.TokenInfo{
				{
					Token: "a",
					Occures: []models.OccureInfo{
						{FilePath: "doc1", OccureCount: 1},
						{FilePath: "doc4", OccureCount: 1},
					},
				},
				{
					Token: "b",
					Occures: []models.OccureInfo{
						{FilePath: "doc2", OccureCount: 1},
						{FilePath: "doc4", OccureCount: 1},
					},
				},
				{
					Token: "c",
					Occures: []models.OccureInfo{
						{FilePath: "doc3", OccureCount: 1},
						{FilePath: "doc4", OccureCount: 1},
					},
				},
			},
			fakeDocs: map[string]models.DocumentInfo{
				"doc1": {
					Filepath: "doc1",
					Tokens: []models.IndexTokenInfo{
						{Token: "a", OccureCount: 10},
					},
				},
				"doc4": {
					Filepath: "doc4",
					Tokens: []models.IndexTokenInfo{
						{Token: "a", OccureCount: 1},
						{Token: "b", OccureCount: 1},
						{Token: "c", OccureCount: 1},
					},
				},
				"doc2": {
					Filepath: "doc2",
					Tokens: []models.IndexTokenInfo{
						{Token: "b", OccureCount: 6},
					},
				},
				"doc3": {
					Filepath: "doc3",
					Tokens: []models.IndexTokenInfo{
						{Token: "c", OccureCount: 7},
					},
				},
			},
			expectedMatrix: [][]float64{
				{1.0, 0, 0},
				{1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0},
				{0, 1.0, 0},
				{0, 0, 1.0},
			},
			expectedTFIDF: []float64{
				1 * math.Log(3),
				2 * math.Log(3),
				3 * math.Log(3),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mongodb.GetDB = func() (models.DBInterface, error) {
				return &FakeDB{docs: test.fakeDocs}, nil
			}

			reqTFIDF, matrix, _, err := Create_TF_IDF_Matrix(test.reqTokens, test.tokenMap, test.relDocsInfo)
			if err != nil {
				t.Fatalf("Unexpected err: %v", err)
			}

			// req vec check
			for i := range reqTFIDF {
				if math.Abs(reqTFIDF[i]-test.expectedTFIDF[i]) > 1e-6 {
					t.Errorf("Incorrect tf-idf on certain index %d: expected %v, recieved %v", i, test.expectedTFIDF[i], reqTFIDF[i])
				}
			}

			// matrix check
			for i := range matrix {
				for j := range matrix[i] {
					if math.Abs(matrix[i][j]-test.expectedMatrix[i][j]) > 1e-6 {
						t.Errorf("Несовпадение матрицы в [%d][%d]: expected %v, recieved %v", i, j, test.expectedMatrix[i][j], matrix[i][j])
					}
				}
			}
		})
	}
}
