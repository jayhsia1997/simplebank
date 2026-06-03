package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jayhsia1997/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	arg := CreateTransferParams{
		FromAccountID: pgtype.Int8{Int64: account1.ID, Valid: true},
		ToAccountID:   pgtype.Int8{Int64: account2.ID, Valid: true},
		Amount:        utils.RandomInt(1, 1000),
	}

	transfer, err := testStore.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	transfer2, err := testStore.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		arg := CreateTransferParams{
			FromAccountID: pgtype.Int8{Int64: account1.ID, Valid: true},
			ToAccountID:   pgtype.Int8{Int64: account2.ID, Valid: true},
			Amount:        utils.RandomInt(1, 1000),
		}
		_, err := testStore.CreateTransfer(context.Background(), arg)
		require.NoError(t, err)
	}

	arg := ListTransfersParams{
		FromAccountID: pgtype.Int8{Int64: account1.ID, Valid: true},
		ToAccountID:   pgtype.Int8{Int64: account2.ID, Valid: true},
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testStore.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
