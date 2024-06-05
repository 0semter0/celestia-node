package shwap

import (
	"errors"
	"fmt"

	"github.com/celestiaorg/nmt"
	nmt_pb "github.com/celestiaorg/nmt/pb"
	"github.com/celestiaorg/rsmt2d"

	"github.com/celestiaorg/celestia-node/share"
	"github.com/celestiaorg/celestia-node/share/shwap/pb"
)

// ErrFailedVerification is returned when inclusion proof verification fails. It is returned
// when the data and the proof do not match trusted data root.
var ErrFailedVerification = errors.New("failed to verify inclusion")

// Sample represents a data share along with its Merkle proof, used to validate the share's
// inclusion in a data square.
type Sample struct {
	share.Share             // Embeds the Share which includes the data with namespace.
	Proof       *nmt.Proof  // Proof is the Merkle Proof validating the share's inclusion.
	ProofType   rsmt2d.Axis // ProofType indicates whether the proof is against a row or a column.
}

// SampleFromProto converts a protobuf Sample back into its domain model equivalent.
func SampleFromProto(s *pb.Sample) Sample {
	proof := nmt.NewInclusionProof(
		int(s.GetProof().GetStart()),
		int(s.GetProof().GetEnd()),
		s.GetProof().GetNodes(),
		s.GetProof().GetIsMaxNamespaceIgnored(),
	)
	return Sample{
		Share:     ShareFromProto(s.GetShare()),
		Proof:     &proof,
		ProofType: rsmt2d.Axis(s.GetProofType()),
	}
}

// ToProto converts a Sample into its protobuf representation for serialization purposes.
func (s Sample) ToProto() *pb.Sample {
	return &pb.Sample{
		Share: &pb.Share{Data: s.Share},
		Proof: &nmt_pb.Proof{
			Start:                 int64(s.Proof.Start()),
			End:                   int64(s.Proof.End()),
			Nodes:                 s.Proof.Nodes(),
			LeafHash:              s.Proof.LeafHash(),
			IsMaxNamespaceIgnored: s.Proof.IsMaxNamespaceIDIgnored(),
		},
		ProofType: pb.AxisType(s.ProofType),
	}
}

// Validate checks the inclusion of the share using its Merkle proof under the specified root.
// Returns an error if the proof is invalid or does not correspond to the indicated proof type.
func (s Sample) Validate(dah *share.Root, rowIdx, colIdx int) error {
	if s.Proof == nil || s.Proof.IsEmptyProof() {
		return errors.New("nil proof")
	}
	if err := share.ValidateShare(s.Share); err != nil {
		return err
	}
	if s.ProofType != rsmt2d.Row && s.ProofType != rsmt2d.Col {
		return fmt.Errorf("invalid SampleProofType: %d", s.ProofType)
	}
	if !s.verifyInclusion(dah, rowIdx, colIdx) {
		return ErrFailedVerification
	}
	return nil
}

// verifyInclusion checks if the share is included in the given root hash at the specified indices.
func (s Sample) verifyInclusion(dah *share.Root, rowIdx, colIdx int) bool {
	size := len(dah.RowRoots)
	namespace := inclusionNamespace(s.Share, rowIdx, colIdx, size)
	rootHash := share.RootHashForCoordinates(dah, s.ProofType, uint(rowIdx), uint(colIdx))
	return s.Proof.VerifyInclusion(
		share.NewSHA256Hasher(),
		namespace.ToNMT(),
		[][]byte{s.Share},
		rootHash,
	)
}

// inclusionNamespace returns the namespace for the share based on its position in the square.
// Shares from extended part of the square are considered parity shares. It means that
// parity shares are located outside of first quadrant of the square. According to the nmt
// specification, the parity shares are prefixed with the namespace of the parity shares.
func inclusionNamespace(sh share.Share, rowIdx, colIdx, squareSize int) share.Namespace {
	isParity := colIdx >= squareSize/2 || rowIdx >= squareSize/2
	if isParity {
		return share.ParitySharesNamespace
	}
	return share.GetNamespace(sh)
}
