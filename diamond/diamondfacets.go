package diamond

import (
	"github.com/neodymium-dev/diamondcontracts-go/facets/diamondcut"
	"github.com/neodymium-dev/diamondcontracts-go/facets/diamondloupe"
	"github.com/neodymium-dev/diamondcontracts-go/facets/facet"
	"github.com/neodymium-dev/diamondcontracts-go/facets/ownership"
)

type diamondFacets struct {
	diamondCut   *diamondcut.Facet
	diamondLoupe *diamondloupe.Facet
	ownership    *ownership.Facet
}

func newDiamondFacets(d *Diamond) (*diamondFacets, error) {
	facets := new(diamondFacets)

	var (
		err       error
		baseFacet = facet.NewFacet(d)
	)

	facets.diamondCut, err = diamondcut.NewFacet(baseFacet)
	if err != nil {
		return nil, err
	}

	facets.diamondLoupe, err = diamondloupe.NewFacet(baseFacet)
	if err != nil {
		return nil, err
	}

	facets.ownership, err = ownership.NewFacet(baseFacet)
	if err != nil {
		return nil, err
	}

	return facets, nil
}
