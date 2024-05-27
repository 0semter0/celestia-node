package shwap

import (
	"crypto/sha256"
	"fmt"

	"github.com/celestiaorg/celestia-app/pkg/wrapper"
	"github.com/celestiaorg/nmt"
	nmt_pb "github.com/celestiaorg/nmt/pb"
	"github.com/celestiaorg/rsmt2d"

	"github.com/celestiaorg/celestia-node/share"
	"github.com/celestiaorg/celestia-node/share/shwap/pb"
)

// RowNamespaceData holds shares and their corresponding proof for a single row within a namespace.
type RowNamespaceData struct {
	Shares []share.Share `json:"shares"` // Shares within the namespace.
	Proof  *nmt.Proof    `json:"proof"`  // Proof of the shares' inclusion in the namespace.
}

// RowNamespaceDataFromEDS extracts and constructs a RowNamespaceData from the row of given EDS
// identified by the index and the namespace.
func RowNamespaceDataFromEDS(
	eds *rsmt2d.ExtendedDataSquare,
	namespace share.Namespace,
	rowIdx int,
) (RowNamespaceData, error) {
	shares := eds.Row(uint(rowIdx))
	rowData, err := RowNamespaceDataFromShares(shares, namespace, rowIdx)
	if err != nil {
		return RowNamespaceData{}, err
	}

	return rowData, nil
}

// RowNamespaceDataFromShares extracts and constructs a RowNamespaceData from shares within the
// specified namespace.
func RowNamespaceDataFromShares(
	shares []share.Share,
	namespace share.Namespace,
	rowIndex int,
) (RowNamespaceData, error) {
	var from, count int
	for i := range len(shares) / 2 {
		if namespace.Equals(share.GetNamespace(shares[i])) {
			if count == 0 {
				from = i
			}
			count++
			continue
		}
		if count > 0 {
			break
		}
	}
	if count == 0 {
		// FIXME: This should return Non-inclusion proofs instead. Need support in app wrapper to generate
		// absence proofs.
		return RowNamespaceData{}, fmt.Errorf("no shares found in the namespace for row %d", rowIndex)
	}

	namespacedShares := make([]share.Share, count)
	copy(namespacedShares, shares[from:from+count])

	tree := wrapper.NewErasuredNamespacedMerkleTree(uint64(len(shares)/2), uint(rowIndex))
	for _, shr := range shares {
		if err := tree.Push(shr); err != nil {
			return RowNamespaceData{}, fmt.Errorf("failed to build tree for row %d: %w", rowIndex, err)
		}
	}

	proof, err := tree.ProveRange(from, from+count)
	if err != nil {
		return RowNamespaceData{}, fmt.Errorf("failed to generate proof for row %d: %w", rowIndex, err)
	}

	return RowNamespaceData{
		Shares: namespacedShares,
		Proof:  &proof,
	}, nil
}

// RowNamespaceDataFromProto constructs RowNamespaceData out of its protobuf representation.
func RowNamespaceDataFromProto(row *pb.RowNamespaceData) RowNamespaceData {
	var proof nmt.Proof
	if row.GetProof().GetLeafHash() != nil {
		proof = nmt.NewAbsenceProof(
			int(row.GetProof().GetStart()),
			int(row.GetProof().GetEnd()),
			row.GetProof().GetNodes(),
			row.GetProof().GetLeafHash(),
			row.GetProof().GetIsMaxNamespaceIgnored(),
		)
	} else {
		proof = nmt.NewInclusionProof(
			int(row.GetProof().GetStart()),
			int(row.GetProof().GetEnd()),
			row.GetProof().GetNodes(),
			row.GetProof().GetIsMaxNamespaceIgnored(),
		)
	}

	return RowNamespaceData{
		Shares: SharesFromProto(row.GetShares()),
		Proof:  &proof,
	}
}

// ToProto converts RowNamespaceData to its protobuf representation for serialization.
func (rnd RowNamespaceData) ToProto() *pb.RowNamespaceData {
	return &pb.RowNamespaceData{
		Shares: SharesToProto(rnd.Shares),
		Proof: &nmt_pb.Proof{
			Start:                 int64(rnd.Proof.Start()),
			End:                   int64(rnd.Proof.End()),
			Nodes:                 rnd.Proof.Nodes(),
			LeafHash:              rnd.Proof.LeafHash(),
			IsMaxNamespaceIgnored: rnd.Proof.IsMaxNamespaceIDIgnored(),
		},
	}
}

// Validate checks validity of the RowNamespaceData against the Root, Namespace and Row index.
func (rnd RowNamespaceData) Validate(dah *share.Root, namespace share.Namespace, rowIdx int) error {
	if rnd.Proof == nil || rnd.Proof.IsEmptyProof() {
		return fmt.Errorf("nil proof")
	}
	if len(rnd.Shares) == 0 && !rnd.Proof.IsOfAbsence() {
		return fmt.Errorf("empty shares with non-absence proof for row %d", rowIdx)
	}

	if len(rnd.Shares) > 0 && rnd.Proof.IsOfAbsence() {
		return fmt.Errorf("non-empty shares with absence proof for row %d", rowIdx)
	}

	if err := ValidateShares(rnd.Shares); err != nil {
		return fmt.Errorf("invalid shares: %w", err)
	}

	rowRoot := dah.RowRoots[rowIdx]
	if namespace.IsOutsideRange(rowRoot, rowRoot) {
		return fmt.Errorf("namespace out of range for row %d", rowIdx)
	}

	if !rnd.verifyInclusion(rowRoot, namespace) {
		return fmt.Errorf("inclusion proof failed for row %d", rowIdx)
	}
	return nil
}

// verifyInclusion checks the inclusion of the row's shares in the provided root using NMT.
func (rnd RowNamespaceData) verifyInclusion(rowRoot []byte, namespace share.Namespace) bool {
	leaves := make([][]byte, 0, len(rnd.Shares))
	for _, shr := range rnd.Shares {
		namespaceBytes := share.GetNamespace(shr)
		leaves = append(leaves, append(namespaceBytes, shr...))
	}
	return rnd.Proof.VerifyNamespace(
		sha256.New(),
		namespace.ToNMT(),
		leaves,
		rowRoot,
	)
}
