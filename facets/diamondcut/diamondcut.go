package diamondcut

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/neodymium-dev/diamondcontracts-go/facets/diamondcut/diamondcutabi"
	"github.com/neodymium-dev/diamondcontracts-go/facets/facet"
)

// Facet is something which should, somewhere down the line, quack like an IDiamondCutFacet.
type Facet struct {
	*facet.Facet
	cut *diamondcutabi.IDiamondCut
}

// NewFacet returns an initialized DiamondCutFacet.
func NewFacet(f *facet.Facet) (*Facet, error) {
	cut, err := diamondcutabi.NewIDiamondCut(f.ConstructorParams())
	if err != nil {
		return nil, err
	}

	return &Facet{
		Facet: f,
		cut:   cut,
	}, nil
}

// DiamondCut is a wrapper around the IDiamondCutFacet.DiamondCut function call.
func (f *Facet) DiamondCut(
	opts *bind.TransactOpts,
	facetCuts []diamondcutabi.IDiamondCutFacetCut,
	initContract common.Address,
	initCalldata []byte,
) (*types.Transaction, error) {
	if opts == nil {
		defaultOpts, err := f.Diamond.TransactOpts()
		if err != nil {
			return nil, err
		}

		opts = defaultOpts
	}

	return f.cut.DiamondCut(opts, facetCuts, initContract, initCalldata)
}
