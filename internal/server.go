package internal

import (
	"context"
	"net"

	"github.com/fluxx1on/finance_transaction_system/internal/config"
	"github.com/fluxx1on/finance_transaction_system/internal/mq"
	"github.com/fluxx1on/finance_transaction_system/internal/mq/producer"
	"github.com/fluxx1on/finance_transaction_system/internal/repo"
	"github.com/fluxx1on/finance_transaction_system/internal/service"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/middleware"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/pb/transfer"
	"github.com/fluxx1on/finance_transaction_system/internal/transport/rpc/pb/user"
	"github.com/jackc/pgx/v5/pgxpool"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Node struct {
	listener net.Listener
	server   *grpc.Server

	sched *mq.ManagerMQ
	prod  *producer.Producer

	db *repo.CreditDB
}

func (n *Node) StartUp(cfg *config.Config) {

	// PostgreSQL
	dbConn, err := pgxpool.New(context.Background(), cfg.PostgreSQL.URL)
	if err != nil {
		slog.Error("PostgreSQL unreached", err)
		panic("startup")
	}

	// CreditDB
	n.db = repo.NewCreditDB(dbConn)

	// RabbitMQ
	rabbitMQ, err := amqp.Dial(cfg.RabbitMQ.URL)
	if err != nil {
		slog.Error("RabbitMQ unreached", err)
		panic("startup")
	}

	// Scheduler setup
	n.sched, err = mq.NewManagerMQ(
		rabbitMQ,
		cfg.RabbitMQ,
		n.db,
	)
	if err != nil {
		slog.Error("Scheduler initialization failed", err)
		panic("startup")
	}

	// Producer setup
	n.prod, err = producer.NewProducer(
		rabbitMQ,
		cfg.RabbitMQ,
	)
	if err != nil {
		slog.Error("Producer initialization failed", err)
		panic("startup")
	}

	// Listener starting
	n.listener, err = net.Listen(cfg.ListenerProtocol, cfg.ServerAddress)
	if err != nil {
		slog.Error("Failed to listen", err)
		panic("startup")
	}

	slog.Info(n.listener.Addr().String())

	// gRPC creating
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(middleware.AuthInterceptor),
	}

	n.server = grpc.NewServer(
		opts...,
	)
	reflection.Register(n.server)

	n.MakeRPC()

	// Servers starting
	go func() {
		if err := n.server.Serve(n.listener); err != nil {
			slog.Error("Failed to serve: %v", err)
		}
	}()

	// Start consumer
	go n.sched.StartWorkers()

	slog.Info("gRPC server started on address:", cfg.ServerAddress)
}

func (n *Node) MakeRPC() {
	trService := rpc.NewTransferService(service.NewTransferFetcher(n.db, n.prod))
	transfer.RegisterTransferServiceServer(n.server, trService)

	opService := rpc.NewOperationService(service.NewOperationFetcher(n.db))
	transfer.RegisterOperationServiceServer(n.server, opService)

	usService := rpc.NewUserService(service.NewUserFetcher(n.db))
	user.RegisterUserServiceServer(n.server, usService)
}

func (n *Node) Stop() {
	if n.sched != nil {
		n.sched.ShutdownJob()
	}
	n.server.Stop()
	n.listener.Close()
}
