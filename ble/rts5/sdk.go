package rts5

import (
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

// BuildSDKMessage builds an SDK message
func BuildSDKMessage(token, id, urlpath, json string) ([]byte, error) {
	return buildMessage(
		rts.NewRtsConnection_5WithRtsSdkProxyRequest(
			&rts.RtsSdkProxyRequest{
				ClientGuid: token,
				MessageId:  id,
				UrlPath:    urlpath,
				Json:       json,
			},
		),
	)
}
