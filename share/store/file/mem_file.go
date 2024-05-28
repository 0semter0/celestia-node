package file

import (
	"context"

	"github.com/celestiaorg/celestia-app/pkg/wrapper"
	"github.com/celestiaorg/nmt"
	"github.com/celestiaorg/rsmt2d"

	"github.com/celestiaorg/celestia-node/share"
	"github.com/celestiaorg/celestia-node/share/ipld"
	"github.com/celestiaorg/celestia-node/share/shwap"
)

var _ EdsFile = (*MemFile)(nil)

type MemFile struct {
	Eds *rsmt2d.ExtendedDataSquare
}

func (f *MemFile) Close() error {
	return nil
}

func (f *MemFile) Size() int {
	return int(f.Eds.Width())
}

func (f *MemFile) Share(
	_ context.Context,
	rowIdx, colIdx int,
) (*shwap.Sample, error) {
	axisType := rsmt2d.Row
	axisIdx, shrIdx := rowIdx, colIdx

	shares := getAxis(f.Eds, axisType, axisIdx)
	tree := wrapper.NewErasuredNamespacedMerkleTree(uint64(f.Size()/2), uint(axisIdx))
	for _, shr := range shares {
		err := tree.Push(shr)
		if err != nil {
			return nil, err
		}
	}

	proof, err := tree.ProveRange(shrIdx, shrIdx+1)
	if err != nil {
		return nil, err
	}

	return &shwap.Sample{
		Share:     shares[shrIdx],
		Proof:     &proof,
		ProofType: axisType,
	}, nil
}

func (f *MemFile) AxisHalf(_ context.Context, axisType rsmt2d.Axis, axisIdx int) (AxisHalf, error) {
	return AxisHalf{
		Shares:   getAxis(f.Eds, axisType, axisIdx)[:f.Size()/2],
		IsParity: false,
	}, nil
}

func (f *MemFile) Data(_ context.Context, namespace share.Namespace, rowIdx int) (shwap.RowNamespaceData, error) {
	shares := getAxis(f.Eds, rsmt2d.Row, rowIdx)
	return ndDataFromShares(shares, namespace, rowIdx)
}

func (f *MemFile) EDS(_ context.Context) (*rsmt2d.ExtendedDataSquare, error) {
	return f.Eds, nil
}

func getAxis(eds *rsmt2d.ExtendedDataSquare, axisType rsmt2d.Axis, axisIdx int) []share.Share {
	switch axisType {
	case rsmt2d.Row:
		return eds.Row(uint(axisIdx))
	case rsmt2d.Col:
		return eds.Col(uint(axisIdx))
	default:
		panic("unknown axis")
	}
}

func ndDataFromShares(shares []share.Share, namespace share.Namespace, rowIdx int) (shwap.RowNamespaceData, error) {
	bserv := ipld.NewMemBlockservice()
	batchAdder := ipld.NewNmtNodeAdder(context.TODO(), bserv, ipld.MaxSizeBatchOption(len(shares)))
	tree := wrapper.NewErasuredNamespacedMerkleTree(uint64(len(shares)/2), uint(rowIdx),
		nmt.NodeVisitor(batchAdder.Visit))
	for _, shr := range shares {
		err := tree.Push(shr)
		if err != nil {
			return shwap.RowNamespaceData{}, err
		}
	}

	root, err := tree.Root()
	if err != nil {
		return shwap.RowNamespaceData{}, err
	}

	err = batchAdder.Commit()
	if err != nil {
		return shwap.RowNamespaceData{}, err
	}

	row, proof, err := ipld.GetSharesByNamespace(context.TODO(), bserv, root, namespace, len(shares))
	if err != nil {
		return shwap.RowNamespaceData{}, err
	}
	return shwap.RowNamespaceData{
		Shares: row,
		Proof:  proof,
	}, nil
}
