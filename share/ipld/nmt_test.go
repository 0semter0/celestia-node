package ipld

import (
	"bytes"
	"crypto/rand"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/celestiaorg/celestia-app/pkg/da"

	"github.com/celestiaorg/celestia-node/share"
)

// TestNamespaceFromCID checks that deriving the Namespaced hash from
// the given CID works correctly.
func TestNamespaceFromCID(t *testing.T) {
	var tests = []struct {
		randData [][]byte
	}{
		// note that the number of shares must be a power of two
		{randData: generateRandNamespacedRawData(4, share.NamespaceSize, share.Size-share.NamespaceSize)},
		{randData: generateRandNamespacedRawData(16, share.NamespaceSize, share.Size-share.NamespaceSize)},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// create DAH from rand data
			eds, err := da.ExtendShares(tt.randData)
			require.NoError(t, err)
			dah := da.NewDataAvailabilityHeader(eds)
			// check to make sure NamespacedHash is correctly derived from CID
			for _, row := range dah.RowRoots {
				c, err := CidFromNamespacedSha256(row)
				require.NoError(t, err)

				got := NamespacedSha256FromCID(c)
				assert.Equal(t, row, got)
			}
		})
	}
}

// generateRandNamespacedRawData returns random namespaced raw data for testing
// purposes. Note that this does not check that total is a power of two.
func generateRandNamespacedRawData(total, nidSize, leafSize uint32) [][]byte {
	data := make([][]byte, total)
	for i := uint32(0); i < total; i++ {
		nid := make([]byte, nidSize)

		_, _ = rand.Read(nid)
		data[i] = nid
	}
	sortByteArrays(data)
	for i := uint32(0); i < total; i++ {
		d := make([]byte, leafSize)

		_, _ = rand.Read(d)
		data[i] = append(data[i], d...)
	}

	return data
}

func sortByteArrays(src [][]byte) {
	sort.Slice(src, func(i, j int) bool { return bytes.Compare(src[i], src[j]) < 0 })
}
