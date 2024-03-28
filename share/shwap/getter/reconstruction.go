package shwap_getter

import (
	"context"
	"github.com/celestiaorg/celestia-node/header"
	"github.com/celestiaorg/celestia-node/share"
	"github.com/celestiaorg/rsmt2d"
)

type ReconstructionGetter struct {
	retriever edsRetriver
}

func (r ReconstructionGetter) GetShare(ctx context.Context, header *header.ExtendedHeader, row, col int) (share.Share, error) {
	return nil, share.ErrOperationNotSupported
}

func (r ReconstructionGetter) GetEDS(ctx context.Context, header *header.ExtendedHeader) (*rsmt2d.ExtendedDataSquare, error) {
	return r.retriever.Retrieve(ctx, header)
}

func (r ReconstructionGetter) GetSharesByNamespace(ctx context.Context, header *header.ExtendedHeader, namespace share.Namespace) (share.NamespacedShares, error) {
	return nil, share.ErrOperationNotSupported
}
