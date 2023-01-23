package share

import (
	"context"
	"github.com/celestiaorg/celestia-node/share/eds"
	"github.com/ipfs/go-datastore"
	ds_sync "github.com/ipfs/go-datastore/sync"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"

	"github.com/celestiaorg/celestia-app/pkg/da"
	"github.com/celestiaorg/celestia-node/share"
)

func Test_EmptyCARExists(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	tmpDir := t.TempDir()
	ds := ds_sync.MutexWrap(datastore.NewMapDatastore())
	edsStore, err := eds.NewStore(tmpDir, ds, os.DirFS(tmpDir))
	require.NoError(t, err)
	err = edsStore.Start(ctx)
	require.NoError(t, err)

	eds := share.EmptyExtendedDataSquare()
	dah := da.NewDataAvailabilityHeader(eds)

	// add empty EDS to store
	err = ensureEmptyCARExists(ctx, edsStore)
	assert.NoError(t, err)

	// assert that the empty car exists
	has, err := edsStore.Has(ctx, dah.Hash())
	assert.True(t, has)
	assert.NoError(t, err)

	// assert that the empty car is, in fact, empty
	emptyEds, err := edsStore.Get(ctx, dah.Hash())
	assert.Equal(t, eds.Flattened(), emptyEds.Flattened())
	assert.NoError(t, err)
}
