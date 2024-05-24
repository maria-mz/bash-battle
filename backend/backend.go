package backend

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/maria-mz/bash-battle-proto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type ServerEvent interface{}

type Backend struct {
	token            string
	conn             *grpc.ClientConn
	grpcClient       proto.BashBattleClient
	stream           proto.BashBattle_StreamClient // NOTE: this is nil by default
	ServerEventsChan chan ServerEvent
	StreamEnded      chan bool
}

func NewBackend(serverAddr string) (*Backend, error) {
	conn, err := grpc.NewClient(
		serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		conn.Close()
		return nil, err
	}

	client := proto.NewBashBattleClient(conn)

	return &Backend{
		conn:             conn,
		grpcClient:       client,
		ServerEventsChan: make(chan ServerEvent),
		StreamEnded:      make(chan bool),
	}, nil
}

func (be *Backend) Shutdown() {
	// TODO: close channels (test)
	if be.conn != nil {
		be.conn.Close()
	}

	if be.stream != nil {
		be.stream.CloseSend()
	}
}

func (be *Backend) getAuthContext() context.Context {
	header := metadata.New(
		map[string]string{"authorization": be.token},
	)
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	return ctx
}

func (be *Backend) Login(request *proto.LoginRequest) (*proto.LoginResponse, error) {
	resp, err := be.grpcClient.Login(context.Background(), request)

	if err == nil && resp.Error == nil {
		be.token = resp.Token
	}

	return resp, err
}

func (be *Backend) Stream() error {
	ctx := be.getAuthContext()

	stream, err := be.grpcClient.Stream(ctx)

	if err != nil {
		return err
	}

	be.stream = stream

	go be.recvStream()

	return nil
}

func (be *Backend) recvStream() {
	for {
		event, err := be.stream.Recv()

		if err != nil {
			slog.Error("failed to receive from game stream")
			// be.StreamEnded <- true
			return
		}

		slog.Info(fmt.Sprintf("received event!!! %+v", event))

		// be.ServerEventsChan <- event
	}
}
