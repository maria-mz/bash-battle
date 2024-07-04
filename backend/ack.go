package backend

import "github.com/maria-mz/bash-battle-proto/proto"

func BuildRoundLoadedAck() *proto.AckMsg {
	ack := &proto.AckMsg{
		Ack: &proto.AckMsg_RoundLoaded{
			RoundLoaded: &proto.RoundLoaded{},
		},
	}

	return ack
}

func BuildRoundSubmissionAck(roundStats *proto.RoundStats) *proto.AckMsg {
	ack := &proto.AckMsg{
		Ack: &proto.AckMsg_RoundSubmission{
			RoundSubmission: &proto.RoundSubmission{
				RoundStats: roundStats,
			},
		},
	}

	return ack
}
