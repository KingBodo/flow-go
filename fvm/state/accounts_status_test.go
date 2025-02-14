package state_test

import (
	"bytes"
	"testing"

	"github.com/onflow/atree"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/fvm/state"
)

func TestAccountStatus(t *testing.T) {

	s := state.NewAccountStatus()
	require.False(t, s.IsAccountFrozen())

	t.Run("test frozen flag set/reset", func(t *testing.T) {
		s.SetFrozenFlag(true)
		require.True(t, s.IsAccountFrozen())

		s.SetFrozenFlag(false)
		require.False(t, s.IsAccountFrozen())
	})

	t.Run("test setting values", func(t *testing.T) {
		// set some values for side effect checks
		s.SetFrozenFlag(true)

		index := atree.StorageIndex{1, 2, 3, 4, 5, 6, 7, 8}
		s.SetStorageIndex(index)
		s.SetPublicKeyCount(34)
		s.SetStorageUsed(56)

		require.Equal(t, uint64(56), s.StorageUsed())
		returnedIndex := s.StorageIndex()
		require.True(t, bytes.Equal(index[:], returnedIndex[:]))
		require.Equal(t, uint64(34), s.PublicKeyCount())

		// check no side effect on flags
		require.True(t, s.IsAccountFrozen())
	})

	t.Run("test serialization", func(t *testing.T) {
		b := append([]byte(nil), s.ToBytes()...)
		clone, err := state.AccountStatusFromBytes(b)
		require.NoError(t, err)
		require.Equal(t, s.IsAccountFrozen(), clone.IsAccountFrozen())
		require.Equal(t, s.StorageIndex(), clone.StorageIndex())
		require.Equal(t, s.PublicKeyCount(), clone.PublicKeyCount())
		require.Equal(t, s.StorageUsed(), clone.StorageUsed())

		// invalid size bytes
		_, err = state.AccountStatusFromBytes([]byte{1, 2})
		require.Error(t, err)
	})
}
