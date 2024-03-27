package shwap

import (
	"context"
	"encoding/binary"
	"fmt"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"

	"github.com/celestiaorg/rsmt2d"

	"github.com/celestiaorg/celestia-node/share"
	"github.com/celestiaorg/celestia-node/share/store/file"
)

// TODO(@walldiss): maybe move into separate subpkg?

// RowIDSize is the size of the RowID in bytes
const RowIDSize = EdsIDSize + 2

// RowID is an unique identifier of a Row.
type RowID struct {
	EdsID

	// RowIndex is the index of the axis(row, col) in the data square
	RowIndex uint16
}

// NewRowID constructs a new RowID.
func NewRowID(height uint64, rowIdx uint16, root *share.Root) (RowID, error) {
	rid := RowID{
		EdsID: EdsID{
			Height: height,
		},
		RowIndex: rowIdx,
	}

	verifyFn := func(row Row) error {
		return row.Verify(root)
	}
	rowVerifiers.Add(rid, verifyFn)

	return rid, rid.Verify(root)
}

// RowIDFromCID coverts CID to RowID.
func RowIDFromCID(cid cid.Cid) (id RowID, err error) {
	if err = validateCID(cid); err != nil {
		return id, err
	}

	err = id.UnmarshalBinary(cid.Hash()[mhPrefixSize:])
	if err != nil {
		return id, fmt.Errorf("while unmarhaling RowID: %w", err)
	}

	return id, nil
}

// Cid returns RowID encoded as CID.
func (rid RowID) Cid() cid.Cid {
	data, err := rid.MarshalBinary()
	if err != nil {
		panic(fmt.Errorf("marshaling RowID: %w", err))
	}

	buf, err := mh.Encode(data, rowMultihashCode)
	if err != nil {
		panic(fmt.Errorf("encoding RowID as CID: %w", err))
	}

	return cid.NewCidV1(rowCodec, buf)
}

// MarshalTo encodes RowID into given byte slice.
// NOTE: Proto is avoided because
// * Its size is not deterministic which is required for IPLD.
// * No support for uint16
func (rid RowID) MarshalTo(data []byte) (int, error) {
	// TODO:(@walldiss): this works, only if data underlying array was preallocated with
	//  enough size. Otherwise Caller might not see the changes.
	data = binary.BigEndian.AppendUint64(data, rid.Height)
	data = binary.BigEndian.AppendUint16(data, rid.RowIndex)
	return RowIDSize, nil
}

// UnmarshalFrom decodes RowID from given byte slice.
func (rid *RowID) UnmarshalFrom(data []byte) (int, error) {
	rid.Height = binary.BigEndian.Uint64(data)
	rid.RowIndex = binary.BigEndian.Uint16(data[8:])
	return RowIDSize, nil
}

// MarshalBinary encodes RowID into binary form.
func (rid RowID) MarshalBinary() ([]byte, error) {
	data := make([]byte, 0, RowIDSize)
	n, err := rid.MarshalTo(data)
	return data[:n], err
}

// UnmarshalBinary decodes RowID from binary form.
func (rid *RowID) UnmarshalBinary(data []byte) error {
	if len(data) != RowIDSize {
		return fmt.Errorf("invalid RowID data length: %d != %d", len(data), RowIDSize)
	}
	_, err := rid.UnmarshalFrom(data)
	return err
}

// Verify verifies RowID fields.
func (rid RowID) Verify(root *share.Root) error {
	if err := rid.EdsID.Verify(root); err != nil {
		return err
	}

	sqrLn := len(root.RowRoots)
	if int(rid.RowIndex) >= sqrLn {
		return fmt.Errorf("RowIndex exceeds square size: %d >= %d", rid.RowIndex, sqrLn)
	}

	return nil
}

func (rid RowID) BlockFromFile(ctx context.Context, f file.EdsFile) (blocks.Block, error) {
	axisHalf, err := f.AxisHalf(ctx, rsmt2d.Row, int(rid.RowIndex))
	if err != nil {
		return nil, fmt.Errorf("while getting AxisHalf: %w", err)
	}

	shares := axisHalf.Shares
	// If it's a parity axis, we need to get the left half of the shares
	if axisHalf.IsParity {
		axis, err := axisHalf.Extended()
		if err != nil {
			return nil, fmt.Errorf("while getting extended shares: %w", err)
		}
		shares = axis[:len(axis)/2]
	}

	s := NewRow(rid, shares)
	blk, err := s.IPLDBlock()
	if err != nil {
		return nil, fmt.Errorf("while coverting to IPLD block: %w", err)
	}
	return blk, nil
}

func (rid RowID) Release() {
	rowVerifiers.Delete(rid)
}
