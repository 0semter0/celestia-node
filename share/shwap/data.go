package shwap

import (
	"context"
	"fmt"

	blocks "github.com/ipfs/go-block-format"

	"github.com/celestiaorg/celestia-app/pkg/wrapper"
	"github.com/celestiaorg/nmt"
	nmtpb "github.com/celestiaorg/nmt/pb"
	"github.com/celestiaorg/rsmt2d"

	"github.com/celestiaorg/celestia-node/share"
	"github.com/celestiaorg/celestia-node/share/ipld"
	shwappb "github.com/celestiaorg/celestia-node/share/shwap/pb"
)

type Data struct {
	DataID
	share.NamespacedRow
}

// NewData constructs a new Data.
func NewData(id DataID, nr share.NamespacedRow) *Data {
	return &Data{
		DataID:        id,
		NamespacedRow: nr,
	}
}

// NewDataFromEDS samples the EDS and constructs Data for each row with the given namespace.
func NewDataFromEDS(
	square *rsmt2d.ExtendedDataSquare,
	height uint64,
	namespace share.Namespace,
) ([]*Data, error) {
	root, err := share.NewRoot(square)
	if err != nil {
		return nil, fmt.Errorf("while computing root: %w", err)
	}

	var datas []*Data //nolint:prealloc// we don't know how many rows with needed namespace there are
	for rowIdx, rowRoot := range root.RowRoots {
		if namespace.IsOutsideRange(rowRoot, rowRoot) {
			continue
		}

		id, err := NewDataID(height, uint16(rowIdx), namespace, root)
		if err != nil {
			return nil, err
		}

		shrs := square.Row(uint(rowIdx))
		// TDOD(@Wondertan): This will likely be removed
		nd, proof, err := ndFromShares(shrs, namespace, rowIdx)
		if err != nil {
			return nil, err
		}

		datas = append(datas, NewData(id, nd, proof))
	}

	return datas, nil
}

func ndFromShares(shrs []share.Share, namespace share.Namespace, axisIdx int) ([]share.Share, nmt.Proof, error) {
	bserv := ipld.NewMemBlockservice()
	batchAdder := ipld.NewNmtNodeAdder(context.TODO(), bserv, ipld.MaxSizeBatchOption(len(shrs)))
	tree := wrapper.NewErasuredNamespacedMerkleTree(uint64(len(shrs)/2), uint(axisIdx),
		nmt.NodeVisitor(batchAdder.Visit))
	for _, shr := range shrs {
		err := tree.Push(shr)
		if err != nil {
			return nil, nmt.Proof{}, err
		}
	}

	root, err := tree.Root()
	if err != nil {
		return nil, nmt.Proof{}, err
	}

	err = batchAdder.Commit()
	if err != nil {
		return nil, nmt.Proof{}, err
	}

	row, proof, err := ipld.GetSharesByNamespace(context.TODO(), bserv, root, namespace, len(shrs))
	if err != nil {
		return nil, nmt.Proof{}, err
	}
	return row, *proof, nil
}

// DataFromBlock converts blocks.Block into Data.
func DataFromBlock(blk blocks.Block) (*Data, error) {
	if err := validateCID(blk.Cid()); err != nil {
		return nil, err
	}
	return DataFromBinary(blk.RawData())
}

// IPLDBlock converts Data to an IPLD block for Bitswap compatibility.
func (s *Data) IPLDBlock() (blocks.Block, error) {
	data, err := s.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return blocks.NewBlockWithCid(data, s.Cid())
}

// MarshalBinary marshals Data to binary.
func (s *Data) MarshalBinary() ([]byte, error) {
	did := s.DataID.MarshalBinary()
	proof := &nmtpb.Proof{}
	proof.Nodes = s.DataProof.Nodes()
	proof.End = int64(s.DataProof.End())
	proof.Start = int64(s.DataProof.Start())
	proof.IsMaxNamespaceIgnored = s.DataProof.IsMaxNamespaceIDIgnored()
	proof.LeafHash = s.DataProof.LeafHash()

	return (&shwappb.Data{
		DataId:     did,
		DataShares: s.DataShares,
		DataProof:  proof,
	}).Marshal()
}

// DataFromBinary unmarshal Data from binary.
func DataFromBinary(data []byte) (*Data, error) {
	proto := &shwappb.Data{}
	if err := proto.Unmarshal(data); err != nil {
		return nil, err
	}

	did, err := DataIDFromBinary(proto.DataId)
	if err != nil {
		return nil, err
	}
	return NewData(did, proto.DataShares, nmt.ProtoToProof(*proto.DataProof)), nil
}

// Verify validates Data's fields and verifies Data inclusion.
func (s *Data) Verify(root *share.Root) error {
	if err := s.DataID.Verify(root); err != nil {
		return err
	}

	if len(s.DataShares) == 0 && s.DataProof.IsEmptyProof() {
		return fmt.Errorf("empty Data")
	}

	shrs := make([][]byte, 0, len(s.DataShares))
	for _, shr := range s.DataShares {
		shrs = append(shrs, append(share.GetNamespace(shr), shr...))
	}

	rowRoot := root.RowRoots[s.RowIndex]
	if !s.DataProof.VerifyNamespace(hashFn(), s.Namespace().ToNMT(), shrs, rowRoot) {
		return fmt.Errorf("invalid DataProof")
	}

	return nil
}
