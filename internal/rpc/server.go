package rpc

import (
	"github.com/osamikoyo/IM-order/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RpcServer struct{
	Channel *amqp.Channel
    Queue *amqp.Queue
}

func New(cfg *config.Config) (*RpcServer, error) {
    
}