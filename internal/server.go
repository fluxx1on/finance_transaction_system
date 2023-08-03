package internal

import (
	"context"
	"net"
	"net/url"

	"github.com/fluxx1on/finance_transaction_system/internal/config"
	"github.com/fluxx1on/finance_transaction_system/internal/database"
	"github.com/fluxx1on/finance_transaction_system/internal/mq"
	"github.com/fluxx1on/finance_transaction_system/internal/mq/producer"
	"github.com/fluxx1on/finance_transaction_system/internal/rpc/operation"
	"github.com/fluxx1on/finance_transaction_system/internal/rpc/pb"
	"github.com/fluxx1on/finance_transaction_system/internal/rpc/transfer"
	"github.com/fluxx1on/finance_transaction_system/internal/rpc/user"
	"github.com/jackc/pgx/v5"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Node struct {
	listener  net.Listener
	server    *grpc.Server
	scheduler *mq.Scheduler
	producer  *producer.Producer
	db        *database.CreditDB
}

func (n *Node) StartUp(cfg *config.Config) {

	// PostgreSQL
	dbConn, err := pgx.Connect(context.Background(), cfg.PostgreSQL)
	if err != nil {
		slog.Error("PostgreSQL unreached", err)
		panic("startup")
	}

	// CreditDB
	n.db = database.NewCreditDB(dbConn)

	// RabbitMQ
	mqURL := &url.URL{
		Scheme: "amqp",
		Host:   cfg.RabbitMQ.Address,
		User:   url.UserPassword(cfg.RabbitMQ.User, cfg.RabbitMQ.Password),
	}
	rabbitMQ, err := amqp.Dial(mqURL.String())
	if err != nil {
		slog.Error("RabbitMQ unreached", err)
		panic("startup")
	}

	// Producer setup
	Producer := producer.NewProducer(
		rabbitMQ,
		cfg.RabbitMQ.TransactionExchangeName,
		cfg.RabbitMQ.RoutingKey,
	)
	n.producer = Producer

	// Scheduler setup
	WorkerScheduler := mq.NewScheduler(
		rabbitMQ,
		cfg.RabbitMQ.QueueAmount,
		cfg.RabbitMQ.WorkerByChannelAmount,
		cfg.RabbitMQ.TransactionExchangeName,
		cfg.RabbitMQ.RoutingKey,
		n.db,
	)
	n.scheduler = WorkerScheduler

	// Listener starting
	n.listener, err = net.Listen(cfg.ListenerProtocol, cfg.ServerAddress)
	if err != nil {
		slog.Error("Failed to listen", err)
		panic("startup")
	}

	slog.Info(n.listener.Addr().String())

	// gRPC creating
	n.server = grpc.NewServer()
	reflection.Register(n.server)

	n.MakeRPC()

	// Server starting
	go func() {
		if err := n.server.Serve(n.listener); err != nil {
			slog.Error("Failed to serve: %v", err)
		}
	}()

	// Start consumer
	go n.scheduler.StartWorkers()

	slog.Info("gRPC server started on address:", cfg.ServerAddress)
}

func (n *Node) MakeRPC() {
	tr_srv := transfer.NewTransferService(transfer.NewTransferFetcher(n.db, n.producer))
	pb.RegisterTransferServiceServer(n.server, tr_srv)

	op_srv := operation.NewOperationService(operation.NewOperationFetcher(n.db))
	pb.RegisterOperationServiceServer(n.server, op_srv)

	us_srv := user.NewUserService(user.NewUserFetcher(n.db))
	pb.RegisterUserServiceServer(n.server, us_srv)
}

func (n *Node) Stop() {
	if n.scheduler != nil {
		n.scheduler.ShutdownJob()
	}
	n.server.Stop()
	n.listener.Close()
}
