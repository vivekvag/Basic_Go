package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vivekvag/Basic_Go/db/util"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, account)
	require.NotEmpty(t, arg.Owner)
	require.NotEmpty(t, arg.Balance)
	require.NotEmpty(t, arg.Currency)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

}
