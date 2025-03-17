package mongodb

import (
	"context"
	"search-engine/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
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
