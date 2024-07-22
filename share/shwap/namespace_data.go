package shwap

import (
	"fmt"

	"github.com/celestiaorg/celestia-node/share"
)

// NamespacedData stores collections of RowNamespaceData, each representing shares and their proofs
// within a namespace.
type NamespacedData []RowNamespaceData

// Flatten combines all shares from all rows within the namespace into a single slice.
func (ns NamespacedData) Flatten() []share.Share {
	var shares []share.Share
	for _, row := range ns {
		shares = append(shares, row.Shares...)
	}
	return shares
}

// Validate checks the integrity of the NamespacedData against a provided root and namespace.
func (ns NamespacedData) Validate(root *share.AxisRoots, namespace share.Namespace) error {
	rowIdxs := share.RowsWithNamespace(root, namespace)
	if len(rowIdxs) != len(ns) {
		return fmt.Errorf("expected %d rows, found %d rows", len(rowIdxs), len(ns))
	}

	for i, row := range ns {
		if err := row.Validate(root, namespace, rowIdxs[i]); err != nil {
			return fmt.Errorf("validating row: %w", err)
		}
	}
	return nil
}
