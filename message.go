package tescat

import (
	"encoding/hex"
	"fmt"
)

type Message struct {
	Type    uint8
	Opcode  uint8
	Data    []byte
	Unknown []byte
}

func NewHexMessage(messageType uint8, opCode uint8, hexData string) (m *Message, err error) {
	m = new(Message)
	m.Type = messageType
	m.Opcode = opCode
	s, err := hex.DecodeString(hexData)
	if err != nil {
		return nil, err
	}
	m.Data = s
	return
}

func NewRawMessage(data []byte) (m *Message) {
	m = new(Message)
	m.Unknown = data[0:2]
	m.Type = data[2]
	m.Opcode = data[3]
	m.Data = data[4:]

	return
}

func (m *Message) String() string {
	return fmt.Sprintf("%02X.%02X %X", m.Type, m.Opcode, m.Data)
}

func (m *Message) Bytes() []byte {
	var length = 4 + len(m.Data)
	out := make([]byte, length)
	out[2] = m.Type
	out[3] = m.Opcode
	copy(out[4:], m.Data)
	return out
}
