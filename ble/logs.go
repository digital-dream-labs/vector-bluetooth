package ble

import (
	"bytes"
	"compress/bzip2"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/digital-dream-labs/vector-bluetooth/rts"
)

// LogResponse is the unified response for log downloading
type LogResponse struct {
	Filename string `json:"filename,omitempty"`
}

// Marshal converts a LogResponse message to bytes
func (sr *LogResponse) Marshal() ([]byte, error) {
	return json.Marshal(sr)
}

// Unmarshal converts a LogResponse byte slice to a LogResponse
func (sr *LogResponse) Unmarshal(b []byte) error {
	return json.Unmarshal(b, sr)
}

// DownloadLogs is an appropriately named function.
func (v *VectorBLE) DownloadLogs() (*LogResponse, error) {
	if !v.state.authorized {
		return nil, errors.New(errNotAuthorized)
	}

	msg, err := rts.BuildLogRequestMessage(v.ble.Version())
	if err != nil {
		return nil, err
	}

	if err := v.ble.Send(msg); err != nil {
		return nil, err
	}

	b, err := v.watch()

	resp := LogResponse{}
	if err := resp.Unmarshal(b); err != nil {
		return nil, err
	}

	return &resp, err
}

func handleRtsLogResponse(v *VectorBLE, msg interface{}) ([]byte, bool, error) {

	var sr *rts.RtsLogResponse

	switch v.ble.Version() {

	case rtsV2:
		t, ok := msg.(*rts.RtsConnection_2)
		if !ok {
			return handlerUnsupportedTypeError()
		}
		sr = t.GetRtsLogResponse()

	case rtsV3:
		t, ok := msg.(*rts.RtsConnection_3)
		if !ok {
			return handlerUnsupportedTypeError()
		}
		sr = t.GetRtsLogResponse()

	case rtsV4:
		t, ok := msg.(*rts.RtsConnection_4)
		if !ok {
			return handlerUnsupportedTypeError()
		}
		sr = t.GetRtsLogResponse()

	case rtsV5:
		t, ok := msg.(*rts.RtsConnection_5)
		if !ok {
			return handlerUnsupportedTypeError()
		}
		sr = t.GetRtsLogResponse()

	default:
		return handlerUnsupportedVersionError()

	}

	if sr.FileId != 0 {
		v.state.filedownload = filedownload{
			FileID: sr.FileId,
			//File:   f,
		}
	}

	return nil, true, nil
}

func handleRtsFileDownload(v *VectorBLE, msg interface{}) ([]byte, bool, error) {

	fmt.Print(".")

	var sr *rts.RtsFileDownload

	switch v.ble.Version() {

	case rtsV2:
		t, ok := msg.(*rts.RtsConnection_2)
		if !ok {
			return handlerUnsupportedTypeError()
		}
		sr = t.GetRtsFileDownload()

	case rtsV3:
		t, ok := msg.(*rts.RtsConnection_3)
		if !ok {
			return handlerUnsupportedTypeError()
		}
		sr = t.GetRtsFileDownload()

	case rtsV4:
		t, ok := msg.(*rts.RtsConnection_4)
		if !ok {
			return handlerUnsupportedTypeError()
		}
		sr = t.GetRtsFileDownload()

	case rtsV5:
		t, ok := msg.(*rts.RtsConnection_5)
		if !ok {
			return handlerUnsupportedTypeError()
		}
		sr = t.GetRtsFileDownload()

	default:
		return handlerUnsupportedVersionError()
	}

	switch {

	case sr.FileId != v.state.filedownload.FileID:
		v.state.filedownload = filedownload{}
		return nil, false, errors.New("invalid file")

	case sr.PacketNumber < sr.PacketTotal:
		v.state.filedownload.Buffer = append(v.state.filedownload.Buffer, sr.FileChunk...)
		return nil, true, nil

	case sr.PacketNumber == sr.PacketTotal:
		fmt.Println("about to uncompress")
		v.state.filedownload.Buffer = append(v.state.filedownload.Buffer, sr.FileChunk...)

		fn, err := v.unBzip()
		if err != nil {
			return nil, false, errors.New("fatal error")
		}

		resp := LogResponse{
			Filename: fn,
		}

		b, err := resp.Marshal()
		if err != nil {
			return nil, false, errors.New("fatal error")
		}

		v.state.filedownload = filedownload{}

		return b, false, nil

	default:
		// something bad happened...
		return nil, false, errors.New("fatal error")
	}
}

func (v *VectorBLE) unBzip() (string, error) {

	t := time.Now()
	fn := t.Format(time.RFC3339)
	output, err := os.OpenFile(
		fmt.Sprintf("%s.tar.gz", fn),
		os.O_APPEND|os.O_RDWR|os.O_CREATE,
		0600,
	)
	if err != nil {
		return "", err
	}

	buf := make([]byte, len(v.state.filedownload.Buffer))

	r := bzip2.NewReader(
		bytes.NewReader(v.state.filedownload.Buffer),
	)

	_, err = r.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		fmt.Println("NO")
		return "", err
	}

	if _, err := output.Write(buf); err != nil {
		return "", err
	}

	_ = output.Close()

	return fn, nil

}
