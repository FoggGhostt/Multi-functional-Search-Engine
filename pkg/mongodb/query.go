package mongodb

import (
	"context"
	"search-engine/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d *DB) UpsertTokenInfos(ctx context.Context, tokenInfos []models.TokenInfo) error {
	coll := d.Cli.Database(d.Cfg.DbName).Collection("InvertIndex")

	for _, tokenInfo := range tokenInfos {
		filter := bson.M{"token": tokenInfo.Token}

		update := bson.M{
			"$push": bson.M{
				"occures": bson.M{
					"$each": tokenInfo.Occures,
				},
			},
		}

		opts := options.Update().SetUpsert(true) //  Если документ не найдется - мы его создадим

		_, err := coll.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DB) UpsertDocInfos(ctx context.Context, docInfo models.DocumentInfo) error {
	coll := d.Cli.Database(d.Cfg.DbName).Collection("Index")
	filter := bson.M{"file_path": docInfo.Filepath}

	update := bson.M{
		"$push": bson.M{
			"tokens": bson.M{
				"$each": docInfo.Tokens,
			},
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err := coll.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) FindRelDocs(ctx context.Context, token string) (*models.TokenInfo, error) {
	coll := d.Cli.Database(d.Cfg.DbName).Collection("InvertIndex")
	filter := bson.M{"token": token}
	var tokenInfo models.TokenInfo
	err := coll.FindOne(ctx, filter).Decode(&tokenInfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &tokenInfo, nil
}

func (d *DB) GetFileIndex(ctx context.Context, filePath string) (*models.DocumentInfo, error) {
	coll := d.Cli.Database(d.Cfg.DbName).Collection("Index")
	filter := bson.M{"file_path": filePath}
	var docInfo models.DocumentInfo
	err := coll.FindOne(ctx, filter).Decode(&docInfo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &docInfo, nil
}
