package ensdiamond_test

import (
	"bytes"
	"context"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/neodymium-dev/diamondcontracts-go/examples/ensdiamond"
)

const rpcUri string = "https://polygon-mainnet.g.alchemy.com/v2/O6nbQONUKZ-V4B_4111Xt7Dg0vm_bQEm"

var (
	rpc            *ethclient.Client
	rpcOnce        sync.Once
	chainId        = big.NewInt(137)
	diamondAddress = common.HexToAddress("0x1409EF959bdc9BeF9Fd4f14517FD44Bbe6601852")
)

var (
	addr         = common.HexToAddress("0x982693778347b2A1817d6B0D2812BE1d52964266")
	addrNodeHash = common.HexToHash("0x45d2f8ebe7df4e75367f83f75a1ff864061098a4649aa8b91d60ef1adfe79c23")
)

func TestUtilsFacet_NodeHash(t *testing.T) {
	c := rpcClient(t)

	sparkly, err := ensdiamond.NewENSDiamond(diamondAddress, chainId, c, nil)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := sparkly.UtilsFacet().NodeHash(addr)
	if err != nil {
		t.Error(err)
		return
	}

	resBytes := make([]byte, 32)
	copy(resBytes[:], res[:])

	resHash := common.BytesToHash(resBytes)

	if !bytes.Equal(resHash.Bytes(), addrNodeHash.Bytes()) {
		t.Errorf("expected %[1]s, got %[2]s", addrNodeHash.String(), resHash.String())
	}
}

func TestNameResolverFacet_Name(t *testing.T) {
	c := rpcClient(t)

	sparkly, err := ensdiamond.NewENSDiamond(diamondAddress, chainId, c, nil)
	if err != nil {
		t.Error(err)
		return
	}

	res, err := sparkly.NameResolver().Name(bytesToBytes32(addrNodeHash.Bytes()))
	if err != nil {
		t.Error(err)
		return
	}

	if len(res) != 0 {
		t.Error("actually got a result wtf")
		return
	}
}

func bytesToBytes32(b []byte) [32]byte {
	b32 := new([32]byte)

	copy((*b32)[:], b[:])

	return *b32
}

func rpcClient(t *testing.T) *ethclient.Client {
	rpcOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		c, err := ethclient.DialContext(ctx, rpcUri)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		rpc = c
	})

	return rpc
}
