package data

import (
	"context"
	"time"

	"github.com/osamikoyo/IM-order/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct{
	coll *mongo.Collection
	ctx context.Context
}

func New(cfg *config.Config) (*Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	url := options.Client().ApplyURI(cfg.MognoURl)
	client, err := mongo.Connect(ctx, url)
	if err != nil{
		return nil,err
	}

	return &Repository{
		coll: client.Database("im").Collection(),
	}, nil
}