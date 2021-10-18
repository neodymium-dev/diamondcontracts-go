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
	facets *ensDiamondFacets
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

	facets, err := newEnsDiamondFacets(baseDiamond)
	if err != nil {
		return nil, err
	}

	return &ENSDiamond{
		Diamond: baseDiamond,
		facets:  facets,
	}, nil
}

func (d *ENSDiamond) AddressResolver() *addressresolver.Facet {
	return d.facets.addressResolver
}

func (d *ENSDiamond) NameResolver() *nameresolver.Facet {
	return d.facets.nameResolver
}

func (d *ENSDiamond) UtilsFacet() *utils.Facet {
	return d.facets.utilsFacet
}

type ensDiamondFacets struct {
	addressResolver *addressresolver.Facet
	nameResolver    *nameresolver.Facet
	utilsFacet      *utils.Facet
}

func newEnsDiamondFacets(d *diamond.Diamond) (*ensDiamondFacets, error) {
	facets := &ensDiamondFacets{}

	var (
		err       error
		baseFacet = facet.NewFacet(d)
	)

	facets.addressResolver, err = addressresolver.NewFacet(baseFacet)
	if err != nil {
		return nil, err
	}

	facets.nameResolver, err = nameresolver.NewFacet(baseFacet)
	if err != nil {
		return nil, err
	}

	facets.utilsFacet, err = utils.NewFacet(baseFacet)
	if err != nil {
		return nil, err
	}

	return facets, nil
}
