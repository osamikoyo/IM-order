package server

import (
	"github.com/osamikoyo/IM-order/internal/data"
	"github.com/osamikoyo/IM-order/pkg/proto/pb"
)

type Server struct{
	pb.UnimplementedOrderServiceServer
	repo *data.Repository
}

