package nameresolver

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/neodymium-dev/diamondcontracts-go/examples/ensdiamond/facets/nameresolver/abi"
	"github.com/neodymium-dev/diamondcontracts-go/facets/facet"
)

type Facet struct {
	*facet.Facet
	abi *abi.NameResolverFacet
}

func NewFacet(f *facet.Facet) (*Facet, error) {
	resolverAbi, err := abi.NewNameResolverFacet(f.ConstructorParams())
	if err != nil {
		return nil, err
	}

	return &Facet{
		Facet: f,
		abi:   resolverAbi,
	}, nil
}

func (f *Facet) Name(node [32]byte) (string, error) {
	return f.abi.Name(&bind.CallOpts{}, node)
}
