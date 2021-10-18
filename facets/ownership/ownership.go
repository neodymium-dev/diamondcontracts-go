package ownership

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/neodymium-dev/diamondcontracts-go/facets/facet"
	"github.com/neodymium-dev/diamondcontracts-go/facets/ownership/ownershipabi"
)

type Facet struct {
	*facet.Facet
	abi *ownershipabi.DiamondOwnership
}

func NewFacet(f *facet.Facet) (*Facet, error) {
	var err error

	ownershipFacet, err := ownershipabi.NewDiamondOwnership(f.ConstructorParams())
	if err != nil {
		return nil, err
	}

	return &Facet{
		Facet: f,
		abi:   ownershipFacet,
	}, nil
}

func (f *Facet) Owner() (common.Address, error) {
	return f.abi.Owner(f.Diamond.CallOpts())
}
