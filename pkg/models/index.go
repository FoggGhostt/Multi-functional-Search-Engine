package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenInfo struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Token   string             `bson:"token"`
	Occures []OccureInfo       `bson:"occures"`
}

type OccureInfo struct {
	FilePath    string `bson:"file_path"`
	OccureCount int64  `bson:"occure_count"`
}

type IndexTokenInfo struct {
	Token       string `bson:"token"`
	OccureCount int64  `bson:"occure_count"`
}

type DocumentInfo struct {
	Filepath string           `bson:"file_path"`
	Tokens   []IndexTokenInfo `bson:"tokens"`
}
