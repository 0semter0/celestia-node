package getters

import (
	"context"
	"fmt"

	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	"golang.org/x/sync/errgroup"

	"github.com/celestiaorg/celestia-node/share"
	"github.com/celestiaorg/celestia-node/share/eds"
	"github.com/celestiaorg/celestia-node/share/ipld"
	"github.com/celestiaorg/nmt"

	"github.com/celestiaorg/nmt/namespace"
)

var _ share.Getter = (*IPLDGetter)(nil)

// IPLDGetter is a share.Getter that retrieves shares from the IPLD network. An EDS store can be
// provided to store shares retrieved from GetShares. Otherwise, result caching is handled by the
// provided blockservice.
type IPLDGetter struct {
	rtrv  *eds.Retriever
	bServ blockservice.BlockService

	// store is an optional eds store that can be used to cache retrieved shares.
	store *eds.Store
}

// NewIPLDGetter creates a new share.Getter that retrieves shares from the IPLD network.
func NewIPLDGetter(bServ blockservice.BlockService) share.Getter {
	return NewIPLDGetterWithStore(bServ, nil)
}

// NewIPLDGetterWithStore creates a new IPLDGetter with an EDS store for storing EDSes from
// GetShares.
func NewIPLDGetterWithStore(bServ blockservice.BlockService, store *eds.Store) share.Getter {
	return &IPLDGetter{
		rtrv:  eds.NewRetriever(bServ),
		bServ: bServ,
		store: store,
	}
}

func (ig *IPLDGetter) GetShare(ctx context.Context, dah *share.Root, row, col int) (share.Share, error) {
	root, leaf := ipld.Translate(dah, row, col)
	nd, err := share.GetShare(ctx, ig.bServ, root, leaf, len(dah.RowsRoots))
	if err != nil {
		return nil, err
	}

	return nd, nil
}

func (ig *IPLDGetter) GetShares(ctx context.Context, root *share.Root) ([][]share.Share, error) {
	eds, err := ig.rtrv.Retrieve(ctx, root)
	if err != nil {
		return nil, err
	}

	// store the EDS if a store was provided
	if ig.store != nil {
		err = ig.store.Put(ctx, root.Hash(), eds)
		if err != nil {
			return nil, fmt.Errorf("failed to cache retrieved shares: %w", err)
		}
	}

	origWidth := int(eds.Width() / 2)
	shares := make([][]share.Share, origWidth)

	for i := 0; i < origWidth; i++ {
		row := eds.Row(uint(i))
		shares[i] = make([]share.Share, origWidth)
		for j := 0; j < origWidth; j++ {
			shares[i][j] = row[j]
		}
	}

	return shares, nil
}

func (ig *IPLDGetter) GetSharesByNamespace(
	ctx context.Context,
	root *share.Root,
	nID namespace.ID,
) (share.NamespaceShares, error) {
	if len(nID) != share.NamespaceSize {
		return nil, fmt.Errorf("expected namespace ID of size %d, got %d", share.NamespaceSize, len(nID))
	}

	rowRootCIDs := make([]cid.Cid, 0)
	for _, row := range root.RowsRoots {
		if !nID.Less(nmt.MinNamespace(row, nID.Size())) && nID.LessOrEqual(nmt.MaxNamespace(row, nID.Size())) {
			rowRootCIDs = append(rowRootCIDs, ipld.MustCidFromNamespacedSha256(row))
		}
	}
	if len(rowRootCIDs) == 0 {
		return nil, nil
	}

	errGroup, ctx := errgroup.WithContext(ctx)
	shares := make([]share.RowNamespaceShares, len(rowRootCIDs))
	for i, rootCID := range rowRootCIDs {
		// shadow loop variables, to ensure correct values are captured
		i, rootCID := i, rootCID
		errGroup.Go(func() (err error) {
			var proof *ipld.Proof
			row, err := share.GetSharesByNamespace(ctx, ig.bServ, rootCID, nID, len(root.RowsRoots), proof)
			shares[i] = share.RowNamespaceShares{
				Shares: row,
				Proof:  proof,
			}
			return
		})
	}

	if err := errGroup.Wait(); err != nil {
		return nil, err
	}

	return shares, nil
}
