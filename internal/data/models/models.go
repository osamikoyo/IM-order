package models

import (
	"time"

	"github.com/osamikoyo/IM-order/pkg/proto/pb"
)

type Product struct{
	Name string `bson:"name"`
	Price uint64 `bson:"price"`
	ID uint64 `bson:"id"`
}

type Order struct{
	UserID uint64 `bson:"user_id"`
	ID uint64 `bson:"id"`
	Status string `bson:"status"`
	Price uint64 `bson:"price"`
	CreatedAt time.Time `bson:"created_at"`
	Products []Product `bson:"products"`
}

func ToPbProducts(p []Product) []*pb.Product {
	var products []*pb.Product
	for _, v := range p {
		products = append(products, &pb.Product{
			Name: v.Name,
			Price: v.Name,
			ID: v.ID,
		})
	}
	return products
}

func ToPb(order *Order) *pb.Order {
	return &pb.Order{
		ID: order.ID,
		UserId: order.UserID,
		Price: order.Price,
		Status: order.Status,
		CreatedAt: order.CreatedAt.GoString(),
		Prodcuts: ToPbProducts(order.Products),
	}
}