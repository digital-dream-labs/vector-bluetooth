package rts4

import "github.com/digital-dream-labs/vector-bluetooth/rts"

// BuildAuthMessage builds the auth request
func BuildAuthMessage(key string) ([]byte, error) {
	return buildMessage(
		rts.NewRtsConnection_4WithRtsCloudSessionRequest(
			&rts.RtsCloudSessionRequest{
				SessionToken: key,
			},
		),
	)
}
