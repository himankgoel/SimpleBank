package db

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner: "Jethiya",
		Currency: "INR",
		Balance: 10000,
	}
	result, err := testQueries.CreateAccount(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	assert.Equal(t, arg.Owner, result.Owner)
	assert.Equal(t, arg.Balance, result.Balance)
	assert.Equal(t, arg.Currency, result.Currency)
	assert.NotZero(t, result.ID)

	testQueries.DeleteAccount(context.Background(), result.ID)
}
