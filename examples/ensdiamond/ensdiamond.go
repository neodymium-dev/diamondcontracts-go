package ensdiamond

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/neodymium-dev/diamondcontracts-go/account"
	"github.com/neodymium-dev/diamondcontracts-go/diamond"
	"github.com/neodymium-dev/diamondcontracts-go/examples/ensdiamond/facets/addressresolver"
	"github.com/neodymium-dev/diamondcontracts-go/examples/ensdiamond/facets/nameresolver"
	"github.com/neodymium-dev/diamondcontracts-go/examples/ensdiamond/facets/utils"
	"github.com/neodymium-dev/diamondcontracts-go/facets/facet"
)

type ENSDiamond struct {
	*diamond.Diamond
	addressResolver *addressresolver.Facet
	nameResolver    *nameresolver.Facet
	utilsFacet      *utils.Facet
}

func NewENSDiamond(
	address common.Address,
	chainID *big.Int,
	rpcClient *ethclient.Client,
	acc *account.Account,
) (*ENSDiamond, error) {

	baseDiamond, err := diamond.NewDiamond(address, chainID, rpcClient, acc)
	if err != nil {
		return nil, err
	}

	addressResolverFacet, err := addressresolver.NewFacet(facet.NewFacet(baseDiamond))
	if err != nil {
		return nil, err
	}

	nameResolverFacet, err := nameresolver.NewFacet(facet.NewFacet(baseDiamond))
	if err != nil {
		return nil, err
	}

	utilsFacet, err := utils.NewFacet(facet.NewFacet(baseDiamond))
	if err != nil {
		return nil, err
	}

	return &ENSDiamond{
		Diamond:         baseDiamond,
		addressResolver: addressResolverFacet,
		nameResolver:    nameResolverFacet,
		utilsFacet:      utilsFacet,
	}, nil
}

func (d *ENSDiamond) AddressResolver() *addressresolver.Facet {
	return d.addressResolver
}

func (d *ENSDiamond) NameResolver() *nameresolver.Facet {
	return d.nameResolver
}

func (d *ENSDiamond) UtilsFacet() *utils.Facet {
	return d.utilsFacet
}
