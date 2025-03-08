package data

import (
	"context"
	"time"

	"github.com/osamikoyo/IM-order/internal/config"
	"github.com/osamikoyo/IM-order/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
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
		coll: client.Database(cfg.DBname).Collection(cfg.Collname),
		ctx: ctx,
	}, nil
}

func (r *Repository) Add(order *models.Order) error {
	_, err := r.coll.InsertOne(r.ctx, order)
	return err
}

func (r *Repository) Delete(id uint64) error {
	res := r.coll.FindOneAndDelete(r.ctx, bson.M{"id" : id})
	return res.Err()
}

func (r *Repository) GetAll() ([]models.Order, error){
	res, err := r.coll.Find(r.ctx, bson.M{})
	if err != nil{
		return nil,err
	}

	var result []models.Order
	for res.Next(r.ctx){
		var r models.Order
		res.Decode(&r)
		result = append(result, r)
	}

	return result, nil
}