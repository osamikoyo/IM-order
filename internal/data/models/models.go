package models

import "time"

type Product struct{
	Name string `bson:"name"`
	Price uint64 `bson:"price"`
}

type Order struct{
	ID uint64 `bson:"id"`
	Status string `bson:"status"`
	Price uint64 `bson:"price"`
	CreatedAt time.Time `bson:"created_at"`
	Products []Product `bson:"products"`
}