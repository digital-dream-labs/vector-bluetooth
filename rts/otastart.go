package rts

import "errors"

// BuildOTAStartMessage builds the ota start message
func BuildOTAStartMessage(version int, url string) ([]byte, error) {
	switch version {
	case rtsv5:
		return buildMessage5(
			NewRtsConnection_5WithRtsOtaUpdateRequest(
				&RtsOtaUpdateRequest{
					Url: url,
				},
			),
		)
	default:
		return nil, errors.New(errUnsupportedVersion)
	}
}
