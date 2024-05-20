package backend

import (
	"context"
	"log/slog"

	pb "github.com/maria-mz/bash-battle-proto/proto"
	"google.golang.org/grpc/metadata"
)

type ServerEvent interface{}

type Backend struct {
	token            string
	grpcClient       pb.BashBattleClient
	stream           pb.BashBattle_StreamGameClient
	ServerEventsChan chan ServerEvent
	StreamEnded      chan bool
}

func NewBackend(grpcClient pb.BashBattleClient) *Backend {
	return &Backend{
		grpcClient:       grpcClient,
		ServerEventsChan: make(chan ServerEvent),
		StreamEnded:      make(chan bool),
	}
}

func (be *Backend) getAuthContext() context.Context {
	header := metadata.New(
		map[string]string{"authorization": be.token},
	)
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	return ctx
}

func (be *Backend) Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	resp, err := be.grpcClient.Login(context.Background(), in)

	if err == nil && resp.ErrorCode == nil {
		be.token = resp.Token
	}

	return resp, err
}

func (be *Backend) CreateGame(in *pb.CreateGameRequest) (*pb.CreateGameResponse, error) {
	ctx := be.getAuthContext()

	resp, err := be.grpcClient.CreateGame(ctx, in)

	return resp, err
}

func (be *Backend) JoinGame(in *pb.JoinGameRequest) (*pb.JoinGameResponse, error) {
	ctx := be.getAuthContext()

	resp, err := be.grpcClient.JoinGame(ctx, in)

	return resp, err
}

func (be *Backend) GetGameConfig(in *pb.ConfigRequest) (*pb.ConfigResponse, error) {
	ctx := be.getAuthContext()

	resp, err := be.grpcClient.GetGameConfig(ctx, in)

	return resp, err
}

func (be *Backend) LeaveGame() error {
	ctx := be.getAuthContext()

	_, err := be.grpcClient.LeaveGame(ctx, &pb.EmptyMessage{})

	return err
}

func (be *Backend) StreamGame() error {
	ctx := be.getAuthContext()

	stream, err := be.grpcClient.StreamGame(ctx, &pb.EmptyMessage{})

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
			be.StreamEnded <- true
			return
		}

		be.ServerEventsChan <- event
	}
}
