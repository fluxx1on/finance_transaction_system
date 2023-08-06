package internal

import (
	"context"
	"net"

	"github.com/fluxx1on/finance_transaction_system/internal/config"
	"github.com/fluxx1on/finance_transaction_system/internal/database"
	"github.com/fluxx1on/finance_transaction_system/internal/mq"
	"github.com/fluxx1on/finance_transaction_system/internal/mq/producer"
	"github.com/fluxx1on/finance_transaction_system/internal/rpc/operation"
	"github.com/fluxx1on/finance_transaction_system/internal/rpc/pb"
	"github.com/fluxx1on/finance_transaction_system/internal/rpc/transfer"
	"github.com/fluxx1on/finance_transaction_system/internal/rpc/user"
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

	db *database.CreditDB
}

func (n *Node) StartUp(cfg *config.Config) {

	// PostgreSQL
	dbConn, err := pgxpool.New(context.Background(), cfg.PostgreSQL)
	if err != nil {
		slog.Error("PostgreSQL unreached", err)
		panic("startup")
	}

	// CreditDB
	n.db = database.NewCreditDB(dbConn)

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
	go n.sched.StartWorkers()

	slog.Info("gRPC server started on address:", cfg.ServerAddress)
}

func (n *Node) MakeRPC() {
	trService := transfer.NewTransferService(transfer.NewTransferFetcher(n.db, n.prod))
	pb.RegisterTransferServiceServer(n.server, trService)

	opService := operation.NewOperationService(operation.NewOperationFetcher(n.db))
	pb.RegisterOperationServiceServer(n.server, opService)

	usService := user.NewUserService(user.NewUserFetcher(n.db))
	pb.RegisterUserServiceServer(n.server, usService)
}

func (n *Node) Stop() {
	if n.sched != nil {
		n.sched.ShutdownJob()
	}
	n.server.Stop()
	n.listener.Close()
}
