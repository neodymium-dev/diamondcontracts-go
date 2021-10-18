package diamondloupe

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/neodymium-dev/diamondcontracts-go/facets/diamondloupe/diamondloupeabi"
	"github.com/neodymium-dev/diamondcontracts-go/facets/facet"
)

type Facet struct {
	*facet.Facet
	loupe *diamondloupeabi.IDiamondLoupe
}

func NewFacet(f *facet.Facet) (*Facet, error) {
	loupe, err := diamondloupeabi.NewIDiamondLoupe(f.ConstructorParams())
	if err != nil {
		return nil, err
	}

	return &Facet{
		Facet: f,
		loupe: loupe,
	}, nil
}

func (f *Facet) FacetAddress(functionSelector [4]byte) (common.Address, error) {
	return f.loupe.FacetAddress(f.Diamond.CallOpts(), functionSelector)
}

func (f *Facet) FacetAddresses() ([]common.Address, error) {
	return f.loupe.FacetAddresses(f.Diamond.CallOpts())
}

func (f *Facet) FacetFunctionSelectors(facetAddress common.Address) ([][4]byte, error) {
	return f.loupe.FacetFunctionSelectors(f.Diamond.CallOpts(), facetAddress)
}

func (f *Facet) Facets() ([]diamondloupeabi.IDiamondLoupeFacet, error) {
	return f.loupe.Facets(f.Diamond.CallOpts())
}

//
// type IDiamondLoupeFacet interface {
// 	FacetAddress() common.Address
// 	FunctionSelectors() [][4]byte
// }
