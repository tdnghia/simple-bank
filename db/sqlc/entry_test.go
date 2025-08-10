package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tdnghia/simple-bank/util"
)

func createRandomAccountAndEntry(t *testing.T) (Account, Entry) {
	account := CreateRandomAccount(t)

	entry := createRandomEntry(t, account)

	return account, entry
}

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.Nil(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomAccountAndEntry(t)
}

func TestGetEntry(t *testing.T) {
	_, entry := createRandomAccountAndEntry(t)

	actualEntry, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.Nil(t, err)
	require.NotEmpty(t, actualEntry)
	require.Equal(t, actualEntry.ID, entry.ID)
	require.Equal(t, actualEntry.AccountID, entry.AccountID)
	require.Equal(t, actualEntry.Amount, entry.Amount)
	require.WithinDuration(t, actualEntry.CreatedAt.Time, entry.CreatedAt.Time, time.Second)
}

func TestListEntries(t *testing.T) {
	account := CreateRandomAccount(t)

	for range 10 {
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
