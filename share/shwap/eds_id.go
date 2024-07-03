package shwap

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// EdsIDSize defines the byte size of the EdsID.
const EdsIDSize = 8

// ErrOutOfBounds is returned whenever an index is out of bounds.
var (
	ErrInvalidShwapID = errors.New("invalid shwap ID")
	ErrOutOfBounds    = fmt.Errorf("index out of bounds: %w", ErrInvalidShwapID)
)

// EdsID represents a unique identifier for a row, using the height of the block
// to identify the data square in the chain.
type EdsID struct {
	Height uint64 // Height specifies the block height.
}

// NewEdsID creates a new EdsID using the given height.
func NewEdsID(height uint64) (EdsID, error) {
	eid := EdsID{
		Height: height,
	}
	return eid, eid.Validate()
}

// EdsIDFromBinary decodes a byte slice into an EdsID, validating the length of the data.
// It returns an error if the data slice does not match the expected size of an EdsID.
func EdsIDFromBinary(data []byte) (EdsID, error) {
	if len(data) != EdsIDSize {
		return EdsID{}, fmt.Errorf("invalid EdsID data length: %d != %d", len(data), EdsIDSize)
	}
	eid := EdsID{
		Height: binary.BigEndian.Uint64(data),
	}
	if err := eid.Validate(); err != nil {
		return EdsID{}, fmt.Errorf("validating EdsID: %w", err)
	}

	return eid, nil
}

// MarshalBinary encodes an EdsID into its binary form, primarily for storage or network
// transmission.
func (eid EdsID) MarshalBinary() ([]byte, error) {
	data := make([]byte, 0, EdsIDSize)
	return eid.appendTo(data), nil
}

// Validate checks the integrity of an EdsID's fields against the provided Root.
// It ensures that the EdsID is not constructed with a zero Height and that the root is not nil.
func (eid EdsID) Validate() error {
	if eid.Height == 0 {
		return fmt.Errorf("%w: Height == 0", ErrInvalidShwapID)
	}
	return nil
}

// appendTo helps in the binary encoding of EdsID by appending the binary form of Height to the
// given byte slice.
func (eid EdsID) appendTo(data []byte) []byte {
	return binary.BigEndian.AppendUint64(data, eid.Height)
}
