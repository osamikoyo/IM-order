package rpc

import (
	"github.com/osamikoyo/IM-order/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RpcServer struct{
	Channel *amqp.Channel
    Queue amqp.Queue
}

func New(cfg *config.Config) (*RpcServer, error) {
    conn,err := amqp.Dial(cfg.AmqpUrl)
	if err != nil{
		return nil, err
	}

	ch,err := conn.Channel()
	if err != nil{
		return nil,err
	}

	que, err := ch.QueueDeclare(
		cfg.QueName,
		false,
		false,
		false,
		false,
		nil,
	)

	return &RpcServer{
		Channel: ch,
		Queue: que,
	}, err
}