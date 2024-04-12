package file

import (
	"context"
	"fmt"
	mrand "math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/celestiaorg/nmt"
	"github.com/celestiaorg/rsmt2d"

	"github.com/celestiaorg/celestia-node/share"
	"github.com/celestiaorg/celestia-node/share/testing/edstest"
	"github.com/celestiaorg/celestia-node/share/testing/sharetest"
)

type createFile func(eds *rsmt2d.ExtendedDataSquare) EdsFile

func testFileShare(t *testing.T, createFile createFile, odsSize int) {
	eds := edstest.RandEDS(t, odsSize)
	fl := createFile(eds)

	dah, err := share.NewRoot(eds)
	require.NoError(t, err)

	width := int(eds.Width())
	t.Run("single thread", func(t *testing.T) {
		for x := 0; x < width; x++ {
			for y := 0; y < width; y++ {
				testShare(t, fl, eds, dah, x, y)
			}
		}
	})

	t.Run("parallel", func(t *testing.T) {
		wg := sync.WaitGroup{}
		for y := 0; y < width; y++ {
			for x := 0; x < width; x++ {
				wg.Add(1)
				go func(x, y int) {
					defer wg.Done()
					testShare(t, fl, eds, dah, x, y)
				}(x, y)
			}
		}
		wg.Wait()
	})
}

func testShare(t *testing.T,
	fl EdsFile,
	eds *rsmt2d.ExtendedDataSquare,
	dah *share.Root,
	x, y int) {
	shr, err := fl.Share(context.TODO(), x, y)
	require.NoError(t, err)

	ok := shr.VerifyInclusion(dah, x, y)
	require.True(t, ok)
}

func testFileData(t *testing.T, createFile createFile, size int) {
	t.Run("included", func(t *testing.T) {
		// generate EDS with random data and some Shares with the same namespace
		namespace := sharetest.RandV0Namespace()
		amount := mrand.Intn(size*size-1) + 1
		eds, dah := edstest.RandEDSWithNamespace(t, namespace, amount, size)
		f := createFile(eds)
		testData(t, f, namespace, dah)
	})

	t.Run("not included", func(t *testing.T) {
		// generate EDS with random data and some Shares with the same namespace
		eds := edstest.RandEDS(t, size)
		dah, err := share.NewRoot(eds)
		require.NoError(t, err)

		maxNs := nmt.MaxNamespace(dah.RowRoots[(len(dah.RowRoots))/2-1], share.NamespaceSize)
		targetNs, err := share.Namespace(maxNs).AddInt(-1)
		require.NoError(t, err)

		f := createFile(eds)
		testData(t, f, targetNs, dah)
	})
}

func testData(t *testing.T, f EdsFile, namespace share.Namespace, dah *share.Root) {
	for i, root := range dah.RowRoots {
		if !namespace.IsOutsideRange(root, root) {
			nd, err := f.Data(context.Background(), namespace, i)
			require.NoError(t, err)
			ok := nd.Verify(root, namespace)
			require.True(t, ok)
		}
	}
}

func testFileAxisHalf(t *testing.T, createFile createFile, odsSize int) {
	eds := edstest.RandEDS(t, odsSize)
	fl := createFile(eds)

	t.Run("single thread", func(t *testing.T) {
		for _, axisType := range []rsmt2d.Axis{rsmt2d.Col, rsmt2d.Row} {
			for i := 0; i < int(eds.Width()); i++ {
				half, err := fl.AxisHalf(context.Background(), axisType, i)
				require.NoError(t, err)
				require.Len(t, half.Shares, odsSize)

				var expected []share.Share
				if half.IsParity {
					expected = getAxis(eds, axisType, i)[odsSize:]
				} else {
					expected = getAxis(eds, axisType, i)[:odsSize]
				}

				require.Equal(t, expected, half.Shares)
			}
		}
	})

	t.Run("parallel", func(t *testing.T) {
		wg := sync.WaitGroup{}
		for _, axisType := range []rsmt2d.Axis{rsmt2d.Col, rsmt2d.Row} {
			for i := 0; i < int(eds.Width()); i++ {
				wg.Add(1)
				go func(axisType rsmt2d.Axis, idx int) {
					defer wg.Done()
					half, err := fl.AxisHalf(context.Background(), axisType, idx)
					require.NoError(t, err)
					require.Len(t, half.Shares, odsSize)

					var expected []share.Share
					if half.IsParity {
						expected = getAxis(eds, axisType, idx)[odsSize:]
					} else {
						expected = getAxis(eds, axisType, idx)[:odsSize]
					}

					require.Equal(t, expected, half.Shares)
				}(axisType, i)
			}
		}
		wg.Wait()
	})
}

func testFileEds(t *testing.T, createFile createFile, size int) {
	eds := edstest.RandEDS(t, size)
	fl := createFile(eds)

	eds2, err := fl.EDS(context.Background())
	require.NoError(t, err)
	require.True(t, eds.Equals(eds2))
}

func testFileReader(t *testing.T, createFile createFile, odsSize int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	eds := edstest.RandEDS(t, odsSize)
	f := createFile(eds)

	// verify that the reader represented by file can be read from
	// multiple times, without exhausting the underlying reader.
	wg := sync.WaitGroup{}
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			testReader(t, ctx, f, eds)
		}()
	}
	wg.Wait()
}

func testReader(t *testing.T, ctx context.Context, f EdsFile, eds *rsmt2d.ExtendedDataSquare) {
	reader, err := f.Reader()
	require.NoError(t, err)

	streamed, err := ReadEds(ctx, reader, f.Size())
	require.NoError(t, err)
	require.True(t, eds.Equals(streamed))
}

func benchGetAxisFromFile(b *testing.B, newFile func(size int) EdsFile, minSize, maxSize int) {
	for size := minSize; size <= maxSize; size *= 2 {
		f := newFile(size)

		// loop over all possible axis types and quadrants
		for _, axisType := range []rsmt2d.Axis{rsmt2d.Row, rsmt2d.Col} {
			for _, squareHalf := range []int{0, 1} {
				name := fmt.Sprintf("Size:%v/Axis:%s/squareHalf:%s", size, axisType, strconv.Itoa(squareHalf))
				b.Run(name, func(b *testing.B) {
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						_, err := f.AxisHalf(context.TODO(), axisType, f.Size()/2*(squareHalf))
						require.NoError(b, err)
					}
				})
			}
		}
	}
}

func benchGetShareFromFile(b *testing.B, newFile func(size int) EdsFile, minSize, maxSize int) {
	for size := minSize; size <= maxSize; size *= 2 {
		f := newFile(size)

		// loop over all possible axis types and quadrants
		for _, q := range quadrants {
			name := fmt.Sprintf("Size:%v/quadrant:%s", size, q)
			b.Run(name, func(b *testing.B) {
				x, y := q.coordinates(f.Size())
				// warm up cache
				_, err := f.Share(context.TODO(), x, y)
				require.NoError(b, err)

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_, err := f.Share(context.TODO(), x, y)
					require.NoError(b, err)
				}
			})
		}

	}
}

type quadrant int

var (
	quadrants = []quadrant{1, 2, 3, 4}
)

func (q quadrant) String() string {
	return strconv.Itoa(int(q))
}

func (q quadrant) coordinates(edsSize int) (x, y int) {
	x = edsSize/2*(int(q-1)%2) + 1
	y = edsSize/2*(int(q-1)/2) + 1
	return
}
