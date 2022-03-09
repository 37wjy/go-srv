package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

var (
	HeadLen uint16 = 8
)

func getHeaderLen() uint16 {
	return HeadLen
}

func UnpackHead(binaryData []byte) (*Message, error) {
	dataBuff := bytes.NewReader(binaryData)

	msg := &Message{}
	fmt.Println("unpack :", binaryData)

	if err := binary.Read(dataBuff, binary.BigEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuff, binary.BigEndian, &msg.ID); err != nil {
		return nil, err
	}

	return msg, nil
}

func Pack(msg *Message) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})

	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.BigEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}
