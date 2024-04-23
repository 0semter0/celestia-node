package file

import (
	"context"
	"errors"
	"io"
	"sync/atomic"

	"github.com/celestiaorg/rsmt2d"

	"github.com/celestiaorg/celestia-node/share"
)

var _ EdsFile = (*closeOnceFile)(nil)

var errFileClosed = errors.New("file closed")

type closeOnceFile struct {
	f        EdsFile
	size     int
	datahash share.DataHash
	closed   atomic.Bool
}

func WithClosedOnce(f EdsFile) EdsFile {
	return &closeOnceFile{
		f:        f,
		size:     f.Size(),
		datahash: f.DataHash(),
	}
}

func (c *closeOnceFile) Close() error {
	if !c.closed.Swap(true) {
		err := c.f.Close()
		// release reference to the file to allow GC to collect all resources associated with it
		c.f = nil
		return err
	}
	return nil
}

func (c *closeOnceFile) Reader() (io.Reader, error) {
	if c.closed.Load() {
		return nil, errFileClosed
	}
	return c.f.Reader()
}

func (c *closeOnceFile) Size() int {
	return c.size
}

func (c *closeOnceFile) DataHash() share.DataHash {
	return c.datahash
}

func (c *closeOnceFile) Share(ctx context.Context, x, y int) (*share.Sample, error) {
	if c.closed.Load() {
		return nil, errFileClosed
	}
	return c.f.Share(ctx, x, y)
}

func (c *closeOnceFile) AxisHalf(ctx context.Context, axisType rsmt2d.Axis, axisIdx int) (AxisHalf, error) {
	if c.closed.Load() {
		return AxisHalf{}, errFileClosed
	}
	return c.f.AxisHalf(ctx, axisType, axisIdx)
}

func (c *closeOnceFile) Data(ctx context.Context, namespace share.Namespace, rowIdx int) (share.RowNamespaceData, error) {
	if c.closed.Load() {
		return share.RowNamespaceData{}, errFileClosed
	}
	return c.f.Data(ctx, namespace, rowIdx)
}

func (c *closeOnceFile) EDS(ctx context.Context) (*rsmt2d.ExtendedDataSquare, error) {
	if c.closed.Load() {
		return nil, errFileClosed
	}
	return c.f.EDS(ctx)
}
