package rts5

import "github.com/digital-dream-labs/vector-bluetooth/rts"

// BuildAuthMessage builds the auth request
func BuildAuthMessage(key string) ([]byte, error) {
	return buildMessage(
		rts.NewRtsConnection_5WithRtsCloudSessionRequest5(
			&rts.RtsCloudSessionRequest_5{
				SessionToken: key,
			},
		),
	)
}
