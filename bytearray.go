package bytearray

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"errors"
	"io"
	"io/ioutil"
)

var (
	ErrOutOfBounds   = errors.New("out of bounds")
	ErrMalformedData = errors.New("malformed data")
)

type compressAlgorithm int

const (
	CompressAlgorithmZLIB compressAlgorithm = iota
	CompressAlgorithmDeflate
)

type ByteArray struct {
	bytes []byte
}

func New(b ...byte) *ByteArray {
	return &ByteArray{
		bytes: b,
	}
}

func (ba *ByteArray) Buffer() *bytes.Buffer {
	return bytes.NewBuffer(ba.bytes)
}

func (ba *ByteArray) Size() uint32 {
	return uint32(len(ba.bytes))
}

func (ba *ByteArray) Skip(n int) error {
	if len(ba.bytes) < n {
		return ErrOutOfBounds
	}
	ba.bytes = ba.bytes[n:]
	return nil
}

func (ba *ByteArray) Clear() {
	ba.bytes = []byte{}
}

func (ba *ByteArray) Bytes() []byte {
	return ba.bytes
}

func (ba *ByteArray) Compress(algorithm compressAlgorithm) (err error) {
	var (
		buffer bytes.Buffer
		writer io.WriteCloser
	)
	switch algorithm {
	case CompressAlgorithmZLIB:
		writer = zlib.NewWriter(&buffer)
	case CompressAlgorithmDeflate:
		writer, err = flate.NewWriter(&buffer, flate.DefaultCompression)
	}
	if err != nil {
		return err
	}
	if _, err = writer.Write(ba.bytes); err != nil {
		return err
	}
	if err = writer.Close(); err != nil {
		return err
	}
	ba.bytes = buffer.Bytes()
	return nil
}

func (ba *ByteArray) Decompress(algorithm compressAlgorithm) (err error) {
	var reader io.ReadCloser
	switch algorithm {
	case CompressAlgorithmZLIB:
		reader, err = zlib.NewReader(bytes.NewReader(ba.bytes))
	case CompressAlgorithmDeflate:
		reader = flate.NewReader(bytes.NewReader(ba.bytes))
	}
	if err != nil {
		return err
	}
	if bb, err := ioutil.ReadAll(reader); err != nil {
		return err
	} else {
		ba.bytes = bb
	}
	return reader.Close()
}
