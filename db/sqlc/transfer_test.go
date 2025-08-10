package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tdnghia/simple-bank/util"
)

func createRandomFromToAccounts(t *testing.T) (Account, Account) {
	fromAccount := CreateRandomAccount(t)
	toAccount := CreateRandomAccount(t)

	return fromAccount, toAccount
}

func createRandomTransfers(t *testing.T, n int) (Account, Account, []Transfer) {
	fromAccount, toAccount := createRandomFromToAccounts(t)

	var transfers []Transfer

	for range n {
		arg := CreateTransferParams{
			FromAccountID: fromAccount.ID,
			ToAccountID:   toAccount.ID,
			Amount:        util.RandomMoney(),
		}

		transfer, err := testQueries.CreateTransfer(context.Background(), arg)

		require.Nil(t, err)
		require.NotEmpty(t, transfer)

		require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
		require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
		require.Equal(t, arg.Amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		transfers = append(transfers, transfer)
	}

	require.Len(t, transfers, n)

	return fromAccount, toAccount, transfers
}

func TestCreateTransfer(t *testing.T) {
	_, _, transfers := createRandomTransfers(t, 1)

	require.Len(t, transfers, 1)
}

func TestGetTransfer(t *testing.T) {
	_, _, transfers := createRandomTransfers(t, 1)

	expectedTransfer := transfers[0]

	actualTransfer, err := testQueries.GetTransfer(context.Background(), expectedTransfer.ID)

	require.Nil(t, err)
	require.NotEmpty(t, actualTransfer)

	require.Equal(t, actualTransfer.ID, expectedTransfer.ID)
	require.Equal(t, actualTransfer.FromAccountID, expectedTransfer.FromAccountID)
	require.Equal(t, actualTransfer.ToAccountID, expectedTransfer.ToAccountID)
	require.WithinDuration(t, actualTransfer.CreatedAt.Time, expectedTransfer.CreatedAt.Time, time.Second)
}

func TestListTransfers(t *testing.T) {
	fromAccount, toAccount, _ := createRandomTransfers(t, 10)

	arg := ListTransfersParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.Nil(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
