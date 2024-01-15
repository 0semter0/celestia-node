package byzantine

import (
	"context"
	"crypto/sha256"
	"hash"
	"testing"
	"time"

	"github.com/ipfs/boxo/blockservice"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	mhcore "github.com/multiformats/go-multihash/core"
	"github.com/stretchr/testify/require"
	core "github.com/tendermint/tendermint/types"

	"github.com/celestiaorg/celestia-app/pkg/da"
	"github.com/celestiaorg/celestia-app/test/util/malicious"
	"github.com/celestiaorg/nmt"
	"github.com/celestiaorg/rsmt2d"

	"github.com/celestiaorg/celestia-node/header"
	"github.com/celestiaorg/celestia-node/share"
	"github.com/celestiaorg/celestia-node/share/eds/edstest"
	"github.com/celestiaorg/celestia-node/share/ipld"
	"github.com/celestiaorg/celestia-node/share/sharetest"
)

// TestIncorrectBadEncodingFraudProof asserts that BEFP is not generated for the correct data
func TestIncorrectBadEncodingFraudProof(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bServ := ipld.NewMemBlockservice()

	squareSize := 8
	shares := sharetest.RandShares(t, squareSize*squareSize)

	eds, err := ipld.AddShares(ctx, shares, bServ)
	require.NoError(t, err)

	dah, err := share.NewRoot(eds)
	require.NoError(t, err)

	// get an arbitrary row
	row := uint(squareSize / 2)
	rowShares := eds.Row(row)
	rowRoot := dah.RowRoots[row]

	shareProofs, err := GetProofsForShares(ctx, bServ, ipld.MustCidFromNamespacedSha256(rowRoot), rowShares)
	require.NoError(t, err)

	// create a fake error for data that was encoded correctly
	fakeError := ErrByzantine{
		Index:  uint32(row),
		Shares: shareProofs,
		Axis:   rsmt2d.Row,
	}

	h := &header.ExtendedHeader{
		RawHeader: core.Header{
			Height: 420,
		},
		DAH: dah,
		Commit: &core.Commit{
			BlockID: core.BlockID{
				Hash: []byte("made up hash"),
			},
		},
	}

	proof := CreateBadEncodingProof(h.Hash(), h.Height(), &fakeError)
	err = proof.Validate(h)
	require.Error(t, err)
}

func TestBEFP_ValidateOutOfOrderShares(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	t.Cleanup(cancel)

	size := 4
	eds := edstest.RandEDS(t, size)

	shares := eds.Flattened()
	shares[0], shares[4] = shares[4], shares[0] // corrupting eds

	bServ := newNamespacedBlockService()
	batchAddr := ipld.NewNmtNodeAdder(ctx, bServ, ipld.MaxSizeBatchOption(size*2))

	eds, err := rsmt2d.ImportExtendedDataSquare(shares,
		share.DefaultRSMT2DCodec(),
		malicious.NewConstructor(uint64(size), nmt.NodeVisitor(batchAddr.Visit)),
	)
	require.NoError(t, err, "failure to recompute the extended data square")

	err = batchAddr.Commit()
	require.NoError(t, err)

	dah, err := da.NewDataAvailabilityHeader(eds)
	require.NoError(t, err)

	var errRsmt2d *rsmt2d.ErrByzantineData
	err = eds.Repair(dah.RowRoots, dah.ColumnRoots)
	require.ErrorAs(t, err, &errRsmt2d)

	byzantine := NewErrByzantine(ctx, bServ, &dah, errRsmt2d)
	var errByz *ErrByzantine
	require.ErrorAs(t, byzantine, &errByz)

	befp := CreateBadEncodingProof([]byte("hash"), 0, errByz)
	err = befp.Validate(&header.ExtendedHeader{DAH: &dah})
	require.NoError(t, err)
}

// namespacedBlockService wraps `BlockService` and extends the verification part
// to avoid returning blocks that has out of order namespaces.
type namespacedBlockService struct {
	blockservice.BlockService
	// the data structure that is used on the networking level, in order
	// to verify the order of the namespaces
	prefix *cid.Prefix
}

func newNamespacedBlockService() *namespacedBlockService {
	sha256NamespaceFlagged := uint64(0x7701)
	// register the nmt hasher to validate the order of namespaces
	mhcore.Register(sha256NamespaceFlagged, func() hash.Hash {
		nh := nmt.NewNmtHasher(sha256.New(), share.NamespaceSize, true)
		nh.Reset()
		return nh
	})

	bs := &namespacedBlockService{}
	bs.BlockService = ipld.NewMemBlockservice()

	bs.prefix = &cid.Prefix{
		Version: 1,
		Codec:   sha256NamespaceFlagged,
		MhType:  sha256NamespaceFlagged,
		// equals to NmtHasher.Size()
		MhLength: sha256.New().Size() + 2*share.NamespaceSize,
	}
	return bs
}

func (n *namespacedBlockService) GetBlock(ctx context.Context, c cid.Cid) (blocks.Block, error) {
	block, err := n.BlockService.GetBlock(ctx, c)
	if err != nil {
		return nil, err
	}

	_, err = n.prefix.Sum(block.RawData())
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (n *namespacedBlockService) GetBlocks(ctx context.Context, cids []cid.Cid) <-chan blocks.Block {
	blockCh := n.BlockService.GetBlocks(ctx, cids)
	resultCh := make(chan blocks.Block)

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(resultCh)
				return
			case block, ok := <-blockCh:
				if !ok {
					close(resultCh)
					return
				}
				if _, err := n.prefix.Sum(block.RawData()); err != nil {
					continue
				}
				resultCh <- block
			}
		}
	}()
	return resultCh
}
