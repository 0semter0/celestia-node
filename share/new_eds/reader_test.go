package eds

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/celestiaorg/celestia-node/share"
	"github.com/celestiaorg/celestia-node/share/eds/edstest"
)

func TestSharesReaderMany(t *testing.T) {
	// create io.Writer that write random data
	for i := 0; i < 10000; i++ {
		TestSharesReader(t)
	}
}

func TestSharesReader(t *testing.T) {
	// create io.Writer that write random data
	odsSize := 16
	eds := edstest.RandEDS(t, odsSize)
	getShare := func(rowIdx, colIdx int) ([]byte, error) {
		fmt.Println("get", rowIdx, colIdx)
		return eds.GetCell(uint(rowIdx), uint(colIdx)), nil
	}

	reader := NewSharesReader(odsSize, getShare)
	readBytes, err := readWithRandomBuffer(reader, 1024)
	require.NoError(t, err)
	expected := make([]byte, 0, odsSize*odsSize*share.Size)
	for _, share := range eds.FlattenedODS() {
		expected = append(expected, share...)
	}
	require.Len(t, readBytes, len(expected))
	require.Equal(t, expected, readBytes)
}

// testRandReader reads from reader with buffers of random sizes.
func readWithRandomBuffer(reader io.Reader, maxBufSize int) ([]byte, error) {
	// create buffer of random size
	data := make([]byte, 0, maxBufSize)
	for {
		bufSize := rand.Intn(maxBufSize-1) + 1
		buf := make([]byte, bufSize)
		n, err := reader.Read(buf)
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		}
		if n < bufSize {
			buf = buf[:n]
		}
		data = append(data, buf...)
		if errors.Is(err, io.EOF) {
			fmt.Println("eof?")
			break
		}
	}
	return data, nil
}
