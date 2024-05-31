package eds

import (
	"context"
	"io"

	"github.com/celestiaorg/rsmt2d"

	"github.com/celestiaorg/celestia-node/share"
	"github.com/celestiaorg/celestia-node/share/shwap"
)

// EDS is an interface for accessing extended data square data.
type EDS interface {
	io.Closer
	// Size returns square size of the file.
	Size() int
	// Sample returns share and corresponding proof for the given axis and share index in this axis.
	Sample(ctx context.Context, rowIdx, colIdx int) (shwap.Sample, error)
	// AxisHalf returns Shares for the first half of the axis of the given type and index.
	AxisHalf(ctx context.Context, axisType rsmt2d.Axis, axisIdx int) (AxisHalf, error)
	// RowData returns data for the given namespace and row index.
	RowData(ctx context.Context, namespace share.Namespace, rowIdx int) (shwap.RowNamespaceData, error)
	// EDS returns extended data square stored in the file.
	EDS(ctx context.Context) (*rsmt2d.ExtendedDataSquare, error)
}
