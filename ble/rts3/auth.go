package rts3

import "github.com/digital-dream-labs/vector-bluetooth/rts"

// BuildAuthMessage builds the auth request
func BuildAuthMessage(key string) ([]byte, error) {
	return buildMessage(
		rts.NewRtsConnection_3WithRtsCloudSessionRequest(
			&rts.RtsCloudSessionRequest{
				SessionToken: key,
			},
		),
	)
}
