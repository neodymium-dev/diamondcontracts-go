package diamond

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/neodymium-dev/diamondcontracts-go/account"
	"github.com/neodymium-dev/diamondcontracts-go/diamond/diamondabi"
	"github.com/neodymium-dev/diamondcontracts-go/facets/diamondcut"
	"github.com/neodymium-dev/diamondcontracts-go/facets/diamondloupe"
	"github.com/neodymium-dev/diamondcontracts-go/facets/facet"
	"github.com/neodymium-dev/diamondcontracts-go/facets/ownership"
)

type Diamond struct {
	address      common.Address
	diamond      *diamondabi.Diamond
	diamondCut   *diamondcut.Facet
	diamondLoupe *diamondloupe.Facet
	ownership    *ownership.Facet
	chainID      *big.Int
	rpc          *ethclient.Client
	account      *account.Account
}

func NewDiamond(
	address common.Address,
	chainID *big.Int,
	rpcClient *ethclient.Client,
	acc *account.Account,
) (*Diamond, error) {

	diamondAbi, err := diamondabi.NewDiamond(address, rpcClient)
	if err != nil {
		return nil, err
	}

	d := &Diamond{
		address: address,
		diamond: diamondAbi,
		chainID: chainID,
		rpc:     rpcClient,
	}

	d.diamondCut, err = diamondcut.NewFacet(facet.NewFacet(d))
	if err != nil {
		return nil, err
	}

	d.diamondLoupe, err = diamondloupe.NewFacet(facet.NewFacet(d))
	if err != nil {
		return nil, err
	}

	d.ownership, err = ownership.NewFacet(facet.NewFacet(d))
	if err != nil {
		return nil, err
	}

	if acc != nil {
		d.account = acc
	}

	return d, nil
}

func (d *Diamond) Address() common.Address {
	return d.address
}

func (d *Diamond) Account() *account.Account {
	return d.account
}

func (d *Diamond) SetAccount(acc *account.Account) {
	d.account = acc
}

func (d *Diamond) ChainID() *big.Int {
	return d.chainID
}

func (d *Diamond) RPCClient() *ethclient.Client {
	return d.rpc
}

func (d *Diamond) Diamond() *diamondabi.Diamond {
	return d.diamond
}

func (d *Diamond) DiamondCut() *diamondcut.Facet {
	return d.diamondCut
}

func (d *Diamond) DiamondLoupe() *diamondloupe.Facet {
	return d.diamondLoupe
}

func (d *Diamond) CallOpts() *bind.CallOpts {
	if d.account == nil {
		return &bind.CallOpts{}
	}

	return d.account.CallOpts()
}

func (d *Diamond) TransactOpts() (*bind.TransactOpts, error) {
	if d.account == nil || d.account.Signer() == nil {
		return nil, fmt.Errorf("account has no signer to build a *bind.TransactOpts")
	}

	return d.account.TransactOpts(d.ChainID())
}
