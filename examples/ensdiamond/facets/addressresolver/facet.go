package addressresolver

import (
	"github.com/neodymium-dev/diamondcontracts-go/examples/ensdiamond/facets/addressresolver/abi"
	"github.com/neodymium-dev/diamondcontracts-go/facets/facet"
)

type Facet struct {
	*facet.Facet
	abi *abi.AddressResolverFacet
}

func NewFacet(f *facet.Facet) (*Facet, error) {
	resolverAbi, err := abi.NewAddressResolverFacet(f.ConstructorParams())
	if err != nil {
		return nil, err
	}

	return &Facet{
		Facet: f,
		abi:   resolverAbi,
	}, nil
}
