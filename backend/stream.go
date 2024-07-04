package backend

import "github.com/maria-mz/bash-battle-proto/proto"

type EndStreamMsg struct {
	Info string
	Err  error
}

type Stream struct {
	streamClient proto.BashBattle_StreamClient
	done         bool

	ServerEvents  chan *proto.Event
	EndStreamMsgs chan EndStreamMsg
}

func NewStream(streamClient proto.BashBattle_StreamClient) *Stream {
	return &Stream{
		streamClient:  streamClient,
		ServerEvents:  make(chan *proto.Event),
		EndStreamMsgs: make(chan EndStreamMsg),
	}
}

func (s *Stream) Recv() {
	if s.done {
		return
	}

	for {
		event, err := s.streamClient.Recv()

		if err != nil {
			s.EndStreamMsgs <- EndStreamMsg{Err: err}
			return
		}

		s.ServerEvents <- event
	}
}

func (s *Stream) Shutdown() {
	s.streamClient.CloseSend()
	close(s.EndStreamMsgs)
	close(s.ServerEvents)
	s.done = true
}

func (s *Stream) SendAck(ack *proto.AckMsg) {
	if s.done {
		return
	}

	if err := s.streamClient.Send(ack); err != nil {
		s.EndStreamMsgs <- EndStreamMsg{Err: err}
	}
}
