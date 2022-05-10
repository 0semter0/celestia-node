package p2p

import (
	"fmt"

	header_pb "github.com/celestiaorg/celestia-node/header/pb"
)

// ExtendedHeaderRequest is the packet format for nodes to request ExtendedHeaders
// from the network.
type ExtendedHeaderRequest struct {
	Origin uint64 // block height from which to request ExtendedHeaders
	Amount uint64 // amount of desired ExtendedHeaders starting from Origin, syncing in ascending order
}

// MarshalBinary marshals ExtendedHeaderRequest to binary.
func (ehr *ExtendedHeaderRequest) MarshalBinary() ([]byte, error) {
	return MarshalExtendedHeaderRequest(ehr)
}

func (ehr *ExtendedHeaderRequest) UnmarshalBinary(data []byte) error {
	if ehr == nil {
		return fmt.Errorf("header: cannot UnmarshalBinary - nil ExtendedHeader")
	}

	out, err := UnmarshalExtendedHeaderRequest(data)
	if err != nil {
		return err
	}

	*ehr = *out
	return nil
}

func (ehr *ExtendedHeaderRequest) ToProto() *header_pb.ExtendedHeaderRequest {
	return &header_pb.ExtendedHeaderRequest{
		Origin: ehr.Origin,
		Amount: ehr.Amount,
	}
}
