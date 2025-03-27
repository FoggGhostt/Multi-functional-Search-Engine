package mongodb

import (
	"context"
	"fmt"
	"search-engine/pkg/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	DbName          string
	MaxPoolSize     uint64
	MinPoolSize     uint64
	MaxConnIdleTime time.Duration
}

type DB struct {
	Cfg *Config
	Cli *mongo.Client
}

var (
	DeFaultMaxPoolSize     uint64        = 100
	DeFaultMinPoolSize     uint64        = 10
	DeFaultMaxConnIdleTime time.Duration = 10 * time.Second
)

func Init(url string, config *Config) (*DB, error) {
	dboption := options.Client().ApplyURI(url)
	fmt.Println(url)

	if config == nil {
		return nil, fmt.Errorf("db config is empty")
	}

	cfg := config

	if cfg.DbName == "" {
		return nil, fmt.Errorf("dbname is empty")
	}

	if cfg.MaxPoolSize == 0 {
		cfg.MaxPoolSize = DeFaultMaxPoolSize

	}

	if cfg.MinPoolSize == 0 {
		cfg.MinPoolSize = DeFaultMinPoolSize

	}

	if cfg.MaxConnIdleTime == 0 {
		cfg.MaxConnIdleTime = DeFaultMaxConnIdleTime

	}

	dboption.SetMaxPoolSize(cfg.MaxPoolSize)
	dboption.SetMinPoolSize(cfg.MinPoolSize)
	dboption.SetMaxConnIdleTime(cfg.MaxConnIdleTime)

	cli, err := mongo.Connect(context.Background(), dboption)

	if err != nil {
		panic(err)
	}

	err = cli.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return &DB{
		Cfg: cfg,
		Cli: cli,
	}, nil
}

func DefaultConfig() *Config {
	return &Config{
		MaxPoolSize:     DeFaultMaxPoolSize,
		MinPoolSize:     DeFaultMinPoolSize,
		MaxConnIdleTime: DeFaultMaxConnIdleTime,
	}
}

func (d *DB) Close() {
	d.Cli.Disconnect(context.TODO())
}

func GetDB() (*DB, error) {
	config, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	mongoURI := config.DBConfig.MongoURI

	cnf := DefaultConfig()
	cnf.DbName = "InvertIndex"

	db, err := Init(mongoURI, cnf)
	if err != nil {
		return nil, err
	}
	return db, nil
}
