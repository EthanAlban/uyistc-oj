package test

import (
	"testing"
	"unioj/utils/future/rpc"
)

func TestRpcServer(t *testing.T) {
	rpc.Server()
}

func TestRpcClient(t *testing.T) {
	rpc.Client()
}
