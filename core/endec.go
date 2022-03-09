package core

import (
	"bytes"
	"encoding/binary"
)

var (
	HeadLen uint16 = 4
)

func getHeaderLen() uint16 {
	return HeadLen
}

func getLen() {

}

func Unpack(binaryData []byte) (*Message, error) {
	dataBuff := bytes.NewReader(binaryData)

	msg := &Message{}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	return msg, nil
}

func Pack(msg *Message) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}
