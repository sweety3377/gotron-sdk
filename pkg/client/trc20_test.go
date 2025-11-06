package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sweety3377/gotron-sdk/pkg/client"
)

func TestTRC20_Balance(t *testing.T) {
	trc20Contract := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" // USDT
	address := "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9"

	conn := client.NewGrpcClient("grpc.trongrid.io:50051")
	err := conn.Start(client.GRPCInsecure())
	require.Nil(t, err)

	balance, err := conn.TRC20ContractBalance(address, trc20Contract)
	require.Nil(t, err)
	assert.Greater(t, balance.Int64(), int64(0))
}
