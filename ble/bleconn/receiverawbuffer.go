package bleconn

import (
	"fmt"
)

const (
	sizeBits = 0b00111111

	msgStart    = 0b10
	msgContinue = 0b00
	msgEnd      = 0b01
	msgSolo     = 0b11
	// msgBits     = 0b11 << 6
)

type bleBuffer struct {
	Buf   []byte `json:"buf,omitempty"`
	State uint   `json:"state,omitempty"`
}

func (b *bleBuffer) receiveRawBuffer(buf []byte) []byte {

	headerByte := buf[0]
	sizeByte := getSize(headerByte)
	multipartState := getMultipartBits(headerByte)

	if int(sizeByte) != len(buf)-1 {
		fmt.Println("SIZE ERROR")
		fmt.Printf("expecting %v, got %v\n", int(sizeByte), len(buf)-1)
		return nil
	}

	// fmt.Printf("buffer state is %v\n", multipartState)
	// fmt.Printf("chunk size is %v\n", int(sizeByte))

	switch multipartState {

	case msgStart:
		// fmt.Println("message start")
		b.Buf = []byte{}
		b.append(buf, int(sizeByte))
		b.State = msgContinue

	case msgContinue:
		// fmt.Println("message continue")
		b.append(buf, int(sizeByte))
		b.State = msgContinue

	case msgEnd:
		// fmt.Println("message end")
		b.append(buf, int(sizeByte))
		b.State = msgStart
		t := b.Buf
		b.Buf = []byte{}
		return t

	case msgSolo:
		// fmt.Println("message solo")
		b.append(buf, int(sizeByte))
		b.State = msgStart
		t := b.Buf
		b.Buf = []byte{}
		return t

	default:
		fmt.Printf("invalid message: %v\n", multipartState)

	}

	return nil
}

func (b *bleBuffer) append(buf []byte, size int) {
	b.Buf = append(b.Buf, buf[1:]...)
}

func getSize(msgSize byte) byte {
	return msgSize & sizeBits
}

func getMultipartBits(msgSize byte) byte {
	//nolint
	return msgSize >> 6
}
