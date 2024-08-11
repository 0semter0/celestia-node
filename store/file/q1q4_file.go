package file

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/celestiaorg/rsmt2d"

	"github.com/celestiaorg/celestia-node/share"
	eds "github.com/celestiaorg/celestia-node/share/new_eds"
	"github.com/celestiaorg/celestia-node/share/shwap"
)

var _ eds.AccessorStreamer = (*Q1Q4File)(nil)

const q1q4FileExtension = ".q4"

// Q1Q4File is an Accessor that contains the first and fourth quadrants of an extended data
// square. It extends the ODSFile with the ability to read the fourth quadrant of the square.
// Reading from the fourth quadrant allows to serve samples from Q2 and Q4 quadrants of the square,
// without the need to read entire Q1.
type Q1Q4File struct {
	ods  *ODSFile
	file *os.File
}

func OpenQ1Q4File(path string) (*Q1Q4File, error) {
	ods, err := OpenODSFile(path)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path + q1q4FileExtension)
	if err != nil {
		return nil, err
	}

	return &Q1Q4File{
		ods:  ods,
		file: f,
	}, nil
}

func CreateQ1Q4File(path string, roots *share.AxisRoots, eds *rsmt2d.ExtendedDataSquare) (*Q1Q4File, error) {
	type res struct {
		ods *ODSFile
		err error
	}
	resCh := make(chan res)
	go func() {
		// creating the file in parallel reduces ~27% of time taken
		ods, err := CreateODSFile(path, roots, eds)
		resCh <- res{ods: ods, err: err}
	}()

	mod := os.O_RDWR | os.O_CREATE | os.O_EXCL // ensure we fail if already exist
	f, err := os.OpenFile(path+q1q4FileExtension, mod, 0o666)
	if err != nil {
		return nil, fmt.Errorf("creating Q4 file: %w", err)
	}

	// buffering gives us ~4x speed up
	buf := bufio.NewWriterSize(f, writeBufferSize)

	err = writeQ4(buf, eds)
	if err != nil {
		return nil, fmt.Errorf("writing Q4: %w", err)
	}

	err = buf.Flush()
	if err != nil {
		return nil, fmt.Errorf("flushing Q4: %w", err)
	}

	err = f.Sync()
	if err != nil {
		return nil, fmt.Errorf("syncing Q4 file: %w", err)
	}

	r := <-resCh
	if r.err != nil {
		return nil, r.err
	}

	return &Q1Q4File{
		ods:  r.ods,
		file: f,
	}, nil
}

func CreateEmptyQ1Q4File(path string) error {
	f, err := os.Create(path + odsFileExtension)
	if err != nil {
		return fmt.Errorf("creating empty ODS file: %w", err)
	}
	if err = f.Close(); err != nil {
		return fmt.Errorf("closing empty ODS file: %w", err)
	}

	f, err = os.Create(path + q1q4FileExtension)
	if err != nil {
		return fmt.Errorf("creating empty Q4 file: %w", err)
	}
	if err = f.Close(); err != nil {
		return fmt.Errorf("closing empty Q4 file: %w", err)
	}
	return nil
}

func RemoveQ1Q4(path string) error {
	err := os.Remove(path + odsFileExtension)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("removing ODS %s: %w", path+odsFileExtension, err)
	}

	err = os.Remove(path + q1q4FileExtension)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("removing Q4 %s: %w", path+q1q4FileExtension, err)
	}
	return nil
}

func LinkQ1Q4File(filepath, linkpath string) error {
	err := os.Link(filepath+odsFileExtension, linkpath+odsFileExtension)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return fmt.Errorf("creating ODS hardlink: %w", err)
	}

	err = os.Link(filepath+q1q4FileExtension, linkpath+q1q4FileExtension)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return fmt.Errorf("creating Q4 hardlink: %w", err)
	}

	return nil
}

func Q1Q4Exists(path string) (bool, error) {
	_, errOds := os.Stat(path + odsFileExtension)
	if errOds != nil && !errors.Is(errOds, os.ErrNotExist) {
		return false, errOds
	}
	_, errQ4 := os.Stat(path + q1q4FileExtension)
	if errQ4 != nil && !errors.Is(errQ4, os.ErrNotExist) {
		return false, errQ4
	}
	return errOds == nil && errQ4 == nil, nil
}

// writeQ4 writes the frth quadrant of the square to the writer. iIt writes the quadrant in row-major
// order
func writeQ4(w io.Writer, eds *rsmt2d.ExtendedDataSquare) error {
	half := eds.Width() / 2
	for i := range half {
		for j := range half {
			shr := eds.GetCell(i+half, j+half) // TODO: Avoid copying inside GetCell
			_, err := w.Write(shr)
			if err != nil {
				return fmt.Errorf("writing share: %w", err)
			}
		}
	}
	return nil
}

func (f *Q1Q4File) Size(ctx context.Context) int {
	return f.ods.Size(ctx)
}

func (f *Q1Q4File) DataHash(ctx context.Context) (share.DataHash, error) {
	return f.ods.DataHash(ctx)
}

func (f *Q1Q4File) AxisRoots(ctx context.Context) (*share.AxisRoots, error) {
	return f.ods.AxisRoots(ctx)
}

func (f *Q1Q4File) Sample(ctx context.Context, rowIdx, colIdx int) (shwap.Sample, error) {
	// use native AxisHalf implementation, to read axis from Q4 quandrant when possible
	half, err := f.AxisHalf(ctx, rsmt2d.Row, rowIdx)
	if err != nil {
		return shwap.Sample{}, fmt.Errorf("reading axis: %w", err)
	}
	shares, err := half.Extended()
	if err != nil {
		return shwap.Sample{}, fmt.Errorf("extending shares: %w", err)
	}
	return shwap.SampleFromShares(shares, rsmt2d.Row, rowIdx, colIdx)
}

func (f *Q1Q4File) AxisHalf(ctx context.Context, axisType rsmt2d.Axis, axisIdx int) (eds.AxisHalf, error) {
	size := f.Size(ctx) // TODO(@Wondertan): Should return error.
	if axisIdx < size/2 {
		half, err := f.ods.AxisHalf(ctx, axisType, axisIdx)
		if err != nil {
			return eds.AxisHalf{}, fmt.Errorf("reading axis half: %w", err)
		}
		return half, nil
	}

	return f.readAxisHalf(ctx, axisType, axisIdx)
}

func (f *Q1Q4File) RowNamespaceData(ctx context.Context,
	namespace share.Namespace,
	rowIdx int,
) (shwap.RowNamespaceData, error) {
	half, err := f.AxisHalf(ctx, rsmt2d.Row, rowIdx)
	if err != nil {
		return shwap.RowNamespaceData{}, fmt.Errorf("reading axis: %w", err)
	}
	shares, err := half.Extended()
	if err != nil {
		return shwap.RowNamespaceData{}, fmt.Errorf("extending shares: %w", err)
	}
	return shwap.RowNamespaceDataFromShares(shares, namespace, rowIdx)
}

func (f *Q1Q4File) Shares(ctx context.Context) ([]share.Share, error) {
	return f.ods.Shares(ctx)
}

func (f *Q1Q4File) Reader() (io.Reader, error) {
	return f.ods.Reader()
}

func (f *Q1Q4File) Close() error {
	return f.ods.Close()
}

func (f *Q1Q4File) readAxisHalf(ctx context.Context, axisType rsmt2d.Axis, axisIdx int) (eds.AxisHalf, error) {
	size := f.Size(ctx)
	q4AxisIdx := axisIdx - size/2
	if q4AxisIdx < 0 {
		return eds.AxisHalf{}, fmt.Errorf("invalid axis index for Q4: %d", axisIdx)
	}

	axisHalf, err := readAxisHalf(f.file, axisType, f.ods.ShareSize(), size, 0, q4AxisIdx)
	if err != nil {
		return eds.AxisHalf{}, fmt.Errorf("reading axis half: %w", err)
	}

	return eds.AxisHalf{
		Shares:   axisHalf,
		IsParity: true,
	}, nil
}
