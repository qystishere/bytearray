package bytearray

import (
	"encoding/binary"
)

func (ba *ByteArray) Read(amount uint32) ([]byte, error) {
	if ba.Size() < amount {
		return nil, ErrOutOfBounds
	}
	bb := ba.bytes[:amount]
	ba.bytes = ba.bytes[amount:]
	return bb, nil
}

func (ba *ByteArray) ReadUint8() (byte, error) {
	if ba.Size() < 1 {
		return 0, ErrOutOfBounds
	}
	b := ba.bytes[0]
	ba.bytes = ba.bytes[1:]
	return b, nil
}

// Same as ReadUint8
func (ba *ByteArray) ReadByte() (byte, error) {
	return ba.ReadUint8()
}

func (ba *ByteArray) ReadInt8() (int8, error) {
	i, err := ba.ReadUint8()
	return int8(i), err
}

func (ba *ByteArray) ReadBool() (bool, error) {
	i, err := ba.ReadUint8()
	return i == 1, err
}

func (ba *ByteArray) ReadUint16() (uint16, error) {
	if ba.Size() < 2 {
		return 0, ErrOutOfBounds
	}
	s := ba.bytes[:2]
	ba.bytes = ba.bytes[2:]
	return binary.BigEndian.Uint16(s), nil
}

// Same as ReadUint16
func (ba *ByteArray) ReadShort() (uint16, error) {
	return ba.ReadUint16()
}

func (ba *ByteArray) ReadInt16() (int16, error) {
	i, err := ba.ReadUint16()
	return int16(i), err
}

func (ba *ByteArray) ReadUint32() (uint32, error) {
	if ba.Size() < 4 {
		return 0, ErrOutOfBounds
	}
	i := ba.bytes[:4]
	ba.bytes = ba.bytes[4:]
	return binary.BigEndian.Uint32(i), nil
}

// Same as ReadUint32
func (ba *ByteArray) ReadInt() (int, error) {
	i, err := ba.ReadUint32()
	return int(i), err
}

func (ba *ByteArray) ReadInt32() (int32, error) {
	i, err := ba.ReadUint32()
	return int32(i), err
}

func (ba *ByteArray) ReadUTF() (string, error) {
	length, err := ba.ReadUint16()
	if err != nil {
		return "", err
	}
	if length == 0 || len(ba.bytes) < int(length) {
		return "", ErrMalformedData
	}
	s := string(ba.bytes[:length])
	ba.bytes = ba.bytes[length:]
	return s, nil
}

func (ba *ByteArray) ReadString() string {
	var s []byte
	for _, b := range ba.bytes {
		if b == 0 {
			break
		}
		s = append(s, b)
	}
	ba.bytes = ba.bytes[len(s):]
	return string(s)
}
