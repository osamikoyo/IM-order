package rpc

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
	"github.com/osamikoyo/IM-order/internal/config"
	"github.com/osamikoyo/IM-order/internal/data"
	"github.com/osamikoyo/IM-order/internal/data/models"
	"github.com/osamikoyo/IM-order/pkg/loger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RpcServer struct{
	Channel *amqp.Channel
    Queue amqp.Queue
	Repo *data.Repository
	Logger loger.Logger
}

func New(cfg *config.Config) (*RpcServer, error) {
	logger := loger.New()

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
	if err != nil{
		return err
	}


	repo,err := data.New(cfg)
	if err != nil{
		return nil,err
	}

	return &RpcServer{
		Channel: ch,
		Queue: que,
		Repo: repo,
		Logger: logger,
	}, err
}

func (r *RpcServer) Listen() error {
	err := r.Channel.Qos(1,0,false)
	if err != nil{
		return err
	}

	msgs, err := r.Channel.Consume(
		r.Queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil{
		return err
	}

	var forever chan struct{}

	go func ()  {
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		defer cancel()
		
		for d := range msgs{
			var req models.Request

			if err = sonic.Unmarshal(d.Body, &req);err != nil{
				r.Logger.Error().Err(err)
			}

			err = r.Repo.UpdateStatus(req.Id, req.Status)
			if err != nil{
				r.Logger.Error().Err(err)
			}

			err = r.Channel.PublishWithContext(ctx,
				"",
				d.ReplyTo,
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					CorrelationId: d.CorrelationId,
					Body: []byte("success!"),
				},
			)

			d.Ack(false)
		}

		<- forever
	}()
	return nil
}