package bytearray

import (
	"encoding/binary"
	"encoding/hex"
)

func (ba *ByteArray) Write(bb []byte) {
	ba.bytes = append(ba.bytes, bb...)
}

func (ba *ByteArray) WriteHex(s string) error {
	bytes, err := hex.DecodeString(s)
	if err != nil {
		return err
	}

	ba.bytes = append(ba.bytes, bytes...)
	return nil
}

func (ba *ByteArray) WriteUint8(b byte) {
	ba.bytes = append(ba.bytes, b)
}

// Same as WriteUint8
func (ba *ByteArray) WriteByte(b byte) {
	ba.WriteUint8(b)
}

func (ba *ByteArray) WriteInt8(b int8) {
	ba.bytes = append(ba.bytes, uint8(b))
}

func (ba *ByteArray) WriteBool(is bool) {
	if is {
		ba.WriteUint8(1)
		return
	}
	ba.WriteUint8(0)
}

func (ba *ByteArray) WriteUint16(i uint16) {
	bb := make([]byte, 2)
	binary.BigEndian.PutUint16(bb, i)
	ba.bytes = append(ba.bytes, bb...)
}

// Same as WriteUint16
func (ba *ByteArray) WriteShort(i uint16) {
	ba.WriteUint16(i)
}

func (ba *ByteArray) WriteInt16(i int16) {
	ba.WriteUint16(uint16(i))
}

func (ba *ByteArray) WriteUint32(i uint32) {
	bb := make([]byte, 4)
	binary.BigEndian.PutUint32(bb, i)
	ba.bytes = append(ba.bytes, bb...)
}

func (ba *ByteArray) WriteInt(i int) {
	ba.WriteUint32(uint32(i))
}

func (ba *ByteArray) WriteInt32(i int32) {
	ba.WriteUint32(uint32(i))
}

func (ba *ByteArray) WriteUTF(s string) {
	ba.WriteUint16(uint16(len(s)))
	ba.bytes = append(ba.bytes, []byte(s)...)
}

func (ba *ByteArray) WriteString(s string) {
	ba.bytes = append(ba.bytes, append([]byte(s), 0)...)
}
