package rts3

import "github.com/digital-dream-labs/vector-bluetooth/rts"

// BuildChallengeResponse builds the challenge response
func BuildChallengeResponse(number uint32) ([]byte, error) {
	return buildMessage(
		rts.NewRtsConnection_3WithRtsChallengeMessage(
			&rts.RtsChallengeMessage{
				Number: number + 1,
			},
		),
	)
}
