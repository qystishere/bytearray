package bytearray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestByteArrayRead(t *testing.T) {
	r := require.New(t)

	b := New()
	b.Write([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	first4, err := b.Read(4)
	r.NoError(err)
	r.ElementsMatch(first4, []byte{1, 2, 3, 4})
	r.ElementsMatch(b.bytes, []byte{5, 6, 7, 8, 9, 10})
	_, err = b.Read(7)
	r.Error(err)
}
