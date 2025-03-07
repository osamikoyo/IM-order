package rpc

import amqp "github.com/rabbitmq/amqp091-go"

type RpcServer struct{
	Channel *amqp.Channel
    Queue *amqp.Queue
}
