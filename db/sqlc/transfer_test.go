package db

import (
	"time"
	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
	"testing"
	"context"
)

func createRandomTransfer(t *testing.T, fromAccount Account, toAccount Account) Transfer{
	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID: toAccount.ID,
		Amount: util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, fromAccount.ID)
	require.Equal(t, arg.ToAccountID, toAccount.ID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	createRandomTransfer(t, fromAccount, toAccount)
}

func TestGetTransfer(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, fromAccount, toAccount)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, fromAccount, toAccount)
	}

	arg := ListTransfersParams{
		FromAccountID: fromAccount.ID,
		ToAccountID: toAccount.ID,
		Limit: 5,
		Offset: 5,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Len(t, transfers, 5)
	
	for _, transfer := range(transfers) {
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
		require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	}
}

