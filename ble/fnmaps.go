package ble

import (
	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

const (
	rtsV2 = 2
	rtsV3 = 3
	rtsV4 = 4
	rtsV5 = 5
)

var rts2Handlers = map[rts.RtsConnection_2Tag]func(v *VectorBLE, msg *rts.RtsConnection_2) ([]byte, bool, error){
	rts.RtsConnection_2Tag_RtsConnRequest:             handleRts2ConnRequest,
	rts.RtsConnection_2Tag_RtsNonceMessage:            handleRTS2NonceRequest,
	rts.RtsConnection_2Tag_RtsChallengeMessage:        handleRTS2ChallengeMessage,
	rts.RtsConnection_2Tag_RtsChallengeSuccessMessage: handleRTS2ChallengeSuccessMessage,
	rts.RtsConnection_2Tag_RtsStatusResponse2:         handleRST2StatusResponse,
	rts.RtsConnection_2Tag_RtsWifiScanResponse2:       handleRST2WifiScanResponse,
	rts.RtsConnection_2Tag_RtsWifiConnectResponse:     handleRST2WifiConnectionResponse,
	rts.RtsConnection_2Tag_RtsOtaUpdateResponse:       handleRST2OtaUpdateResponse,
}

var rts3Handlers = map[rts.RtsConnection_3Tag]func(v *VectorBLE, msg *rts.RtsConnection_3) ([]byte, bool, error){
	rts.RtsConnection_3Tag_RtsConnRequest:             handleRts3ConnRequest,
	rts.RtsConnection_3Tag_RtsNonceMessage:            handleRTS3NonceRequest,
	rts.RtsConnection_3Tag_RtsChallengeMessage:        handleRTS3ChallengeMessage,
	rts.RtsConnection_3Tag_RtsChallengeSuccessMessage: handleRTS3ChallengeSuccessMessage,
	rts.RtsConnection_3Tag_RtsStatusResponse3:         handleRST3StatusResponse,
	rts.RtsConnection_3Tag_RtsWifiScanResponse3:       handleRST3WifiScanResponse,
	rts.RtsConnection_3Tag_RtsWifiConnectResponse3:    handleRST3WifiConnectionResponse,
	rts.RtsConnection_3Tag_RtsOtaUpdateResponse:       handleRST3OtaUpdateResponse,
	rts.RtsConnection_3Tag_RtsCloudSessionResponse:    handleRST3CloudSessionResponse,
}

var rts4Handlers = map[rts.RtsConnection_4Tag]func(v *VectorBLE, msg *rts.RtsConnection_4) ([]byte, bool, error){
	rts.RtsConnection_4Tag_RtsConnRequest:             handleRts4ConnRequest,
	rts.RtsConnection_4Tag_RtsNonceMessage:            handleRTS4NonceRequest,
	rts.RtsConnection_4Tag_RtsChallengeMessage:        handleRTS4ChallengeMessage,
	rts.RtsConnection_4Tag_RtsChallengeSuccessMessage: handleRTS4ChallengeSuccessMessage,
	rts.RtsConnection_4Tag_RtsStatusResponse4:         handleRST4StatusResponse,
	rts.RtsConnection_4Tag_RtsWifiScanResponse3:       handleRST4WifiScanResponse,
	rts.RtsConnection_4Tag_RtsWifiConnectResponse3:    handleRST4WifiConnectionResponse,
	rts.RtsConnection_4Tag_RtsOtaUpdateResponse:       handleRST4OtaUpdateResponse,
	rts.RtsConnection_4Tag_RtsCloudSessionResponse:    handleRST4CloudSessionResponse,
}

var rts5Handlers = map[rts.RtsConnection_5Tag]func(v *VectorBLE, msg *rts.RtsConnection_5) ([]byte, bool, error){
	rts.RtsConnection_5Tag_RtsConnRequest:             handleRts5ConnRequest,
	rts.RtsConnection_5Tag_RtsNonceMessage:            handleRTS5NonceRequest,
	rts.RtsConnection_5Tag_RtsChallengeMessage:        handleRTS5ChallengeMessage,
	rts.RtsConnection_5Tag_RtsChallengeSuccessMessage: handleRTS5ChallengeSuccessMessage,
	rts.RtsConnection_5Tag_RtsStatusResponse5:         handleRST5StatusResponse,
	rts.RtsConnection_5Tag_RtsWifiScanResponse3:       handleRST5WifiScanResponse,
	rts.RtsConnection_5Tag_RtsWifiConnectResponse3:    handleRST5WifiConnectionResponse,
	rts.RtsConnection_5Tag_RtsOtaUpdateResponse:       handleRST5OtaUpdateResponse,
	rts.RtsConnection_5Tag_RtsCloudSessionResponse:    handleRST5CloudSessionResponse,
	rts.RtsConnection_5Tag_RtsSdkProxyResponse:        handleRST5SDKProxyResponse,
}
