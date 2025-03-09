package server

import (
	"context"

	"github.com/osamikoyo/IM-order/internal/data"
	"github.com/osamikoyo/IM-order/pkg/proto/pb"
)

type Server struct{
	pb.UnimplementedOrderServiceServer
	repo *data.Repository
}

func (s *Server) Create(_ context.Context,req *pb.CreateReq) (*pb.Response, error){
	
}

func (s *Server) Delete(_ context.Context,req *pb.DeleteReq) (*pb.Response, error){

}

func (s *Server) Get(_ context.Context,req *pb.GetReq) (*pb.GetResp, error){

}