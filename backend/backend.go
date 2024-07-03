package backend

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/maria-mz/bash-battle-proto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (be *Backend) Connect(request *proto.ConnectRequest) error {
	resp, err := be.grpcClient.Connect(context.Background(), request)

	if err == nil {
		be.token = resp.Token
	}

	return err
}

func (be *Backend) JoinGame() error {
	ctx := be.getAuthContext()

	_, err := be.grpcClient.JoinGame(ctx, &emptypb.Empty{})

	return err
}

func (be *Backend) GetGameConfig() (*proto.GameConfig, error) {
	ctx := be.getAuthContext()

	gameConfig, err := be.grpcClient.GetGameConfig(ctx, &emptypb.Empty{})

	return gameConfig, err
}

func (be *Backend) GetPlayers() (*proto.Players, error) {
	ctx := be.getAuthContext()

	players, err := be.grpcClient.GetPlayers(ctx, &emptypb.Empty{})

	return players, err
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
