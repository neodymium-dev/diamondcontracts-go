package utils

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/neodymium-dev/diamondcontracts-go/examples/ensdiamond/facets/utils/abi"
	"github.com/neodymium-dev/diamondcontracts-go/facets/facet"
)

type Facet struct {
	*facet.Facet
	abi *abi.UtilsFacet
}

func NewFacet(f *facet.Facet) (*Facet, error) {
	resolverAbi, err := abi.NewUtilsFacet(f.ConstructorParams())
	if err != nil {
		return nil, err
	}

	return &Facet{
		Facet: f,
		abi:   resolverAbi,
	}, nil
}

func (f *Facet) NodeHash(address common.Address) ([32]byte, error) {
	return f.abi.Node(f.Diamond.CallOpts(), address)
}
