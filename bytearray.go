package bytearray

import (
	"bytes"
	"compress/flate"
	"compress/zlib"
	"errors"
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

func (ba *ByteArray) Compress(algorithm compressAlgorithm) error {
	var in bytes.Buffer
	switch algorithm {
	case CompressAlgorithmZLIB:
		w := zlib.NewWriter(&in)
		_, err := w.Write(ba.bytes)
		if err != nil {
			return err
		}
		err = w.Close()
		if err != nil {
			return err
		}
		ba.bytes = in.Bytes()
	case CompressAlgorithmDeflate:
		w, err := flate.NewWriter(&in, flate.DefaultCompression)
		if err != nil {
			return err
		}
		_, err = w.Write(ba.bytes)
		if err != nil {
			return err
		}
		err = w.Close()
		if err != nil {
			return err
		}
		ba.bytes = in.Bytes()
	}
	return nil
}

func (ba *ByteArray) Decompress(algorithm compressAlgorithm) error {
	switch algorithm {
	case CompressAlgorithmZLIB:
		r, err := zlib.NewReader(bytes.NewReader(ba.bytes))
		if err != nil {
			return err
		}
		bb, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}
		err = r.Close()
		if err != nil {
			return err
		}
		ba.bytes = bb
	case CompressAlgorithmDeflate:
		r := flate.NewReader(bytes.NewReader(ba.bytes))
		bb, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}
		err = r.Close()
		if err != nil {
			return err
		}
		ba.bytes = bb
	}
	return nil
}
