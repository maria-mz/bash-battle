package backend

import (
	"context"

	"github.com/maria-mz/bash-battle-proto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Backend struct {
	token      string
	conn       *grpc.ClientConn
	grpcClient proto.BashBattleClient
	stream     *Stream
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
		conn:       conn,
		grpcClient: client,
	}, nil
}

func (be *Backend) Shutdown() {
	if be.conn != nil {
		be.conn.Close()
	}

	if be.stream != nil {
		be.stream.Shutdown()
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
	// TODO: Return error if stream already active
	ctx := be.getAuthContext()

	streamClient, err := be.grpcClient.Stream(ctx)

	if err != nil {
		return err
	}

	be.stream = NewStream(streamClient)

	go be.stream.Recv()

	endStreamMsg := <-be.stream.EndStreamMsgs // Blocking

	return endStreamMsg.Err
}

func (be *Backend) SendRoundLoadedAck() {
	ack := BuildRoundLoadedAck()
	be.stream.SendAck(ack)
}

func (be *Backend) SendScoreSubmissionAck(roundStats *proto.RoundStats) {
	ack := BuildRoundSubmissionAck(roundStats)
	be.stream.SendAck(ack)
}

func (be *Backend) GetServerEvents() <-chan *proto.Event {
	return be.stream.ServerEvents
}
