package shwap

import (
	"bytes"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/celestiaorg/celestia-node/share"
	"github.com/celestiaorg/celestia-node/share/eds/edstest"
	"github.com/celestiaorg/celestia-node/share/sharetest"
)

func TestNamespacedRowFromShares(t *testing.T) {
	const odsSize = 8

	minNamespace, err := share.NewBlobNamespaceV0(slices.Concat(bytes.Repeat([]byte{0}, 8), []byte{1, 0}))
	require.NoError(t, err)
	err = minNamespace.ValidateForData()
	require.NoError(t, err)

	for namespacedAmount := 1; namespacedAmount < odsSize; namespacedAmount++ {
		shares := sharetest.RandSharesWithNamespace(t, minNamespace, namespacedAmount, odsSize)
		parity, err := share.DefaultRSMT2DCodec().Encode(shares)
		require.NoError(t, err)
		extended := slices.Concat(shares, parity)

		nr, err := NamespacedRowFromShares(extended, minNamespace, 0)
		require.NoError(t, err)
		require.Equal(t, namespacedAmount, len(nr.Shares))
	}
}

func TestNamespacedRowFromSharesNonIncluded(t *testing.T) {
	// TODO: this will fail until absence proof support is added
	t.Skip()

	const odsSize = 8
	// Test absent namespace
	shares := sharetest.RandShares(t, odsSize)
	absentNs, err := share.GetNamespace(shares[0]).AddInt(1)
	require.NoError(t, err)

	parity, err := share.DefaultRSMT2DCodec().Encode(shares)
	require.NoError(t, err)
	extended := slices.Concat(shares, parity)

	nr, err := NamespacedRowFromShares(extended, absentNs, 0)
	require.NoError(t, err)
	require.Len(t, nr.Shares, 0)
	require.True(t, nr.Proof.IsOfAbsence())
}

func TestNamespacedSharesFromEDS(t *testing.T) {
	const odsSize = 8
	sharesAmount := odsSize * odsSize
	namespace := sharetest.RandV0Namespace()
	for amount := 1; amount < sharesAmount; amount++ {
		eds, root := edstest.RandEDSWithNamespace(t, namespace, amount, odsSize)
		nd, err := NewNamespacedSharesFromEDS(eds, namespace)
		require.NoError(t, err)
		require.True(t, len(nd) > 0)
		require.Len(t, nd.Flatten(), amount)

		err = nd.Verify(root, namespace)
		require.NoError(t, err)
	}
}

func TestValidateNamespacedRow(t *testing.T) {
	const odsSize = 8
	sharesAmount := odsSize * odsSize
	namespace := sharetest.RandV0Namespace()
	for amount := 1; amount < sharesAmount; amount++ {
		eds, root := edstest.RandEDSWithNamespace(t, namespace, amount, odsSize)
		nd, err := NewNamespacedSharesFromEDS(eds, namespace)
		require.NoError(t, err)
		require.True(t, len(nd) > 0)

		_, from, to := share.FilterRootByNamespace(root, namespace)
		require.Len(t, nd, to-from)
		idx := from
		for _, row := range nd {
			err = row.Validate(root, namespace, idx)
			require.NoError(t, err)
			idx++
		}
	}
}

func TestNamespacedRowProtoEncoding(t *testing.T) {
	const odsSize = 8
	namespace := sharetest.RandV0Namespace()
	eds, _ := edstest.RandEDSWithNamespace(t, namespace, odsSize, odsSize)
	nd, err := NewNamespacedSharesFromEDS(eds, namespace)
	require.NoError(t, err)
	require.True(t, len(nd) > 0)

	expected := nd[0]
	pb := expected.ToProto()
	ndOut := NamespacedRowFromProto(pb)
	require.Equal(t, expected, ndOut)
}
