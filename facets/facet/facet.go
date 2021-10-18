package facet

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// diamond is a hacky way to pass a *diamond.Diamond without actually passing the concrete type,
// so as to prevent import cycle issues.
type diamond interface {
	Address() common.Address
	RPCClient() *ethclient.Client
	CallOpts() *bind.CallOpts
	TransactOpts() (*bind.TransactOpts, error)
}

// Facet was implemented to fix weird import cycle issues.
// It acts as a base class for all facets, providing access to various
// parts of a parent diamond.Diamond via the unexported "diamond" interface.
// Yeah, it's as hacky as it sounds.
type Facet struct {
	Diamond diamond
}

// NewFacet returns a *Facet, and should be passed something which quacks like a *diamond.Diamond.
func NewFacet(f diamond) *Facet {
	return &Facet{f}
}

// ConstructorParams returns the Facet's diamond address and ethclient.Client,
// which are commonly used for abigen constructors.
func (f *Facet) ConstructorParams() (common.Address, *ethclient.Client) {
	return f.Diamond.Address(), f.Diamond.RPCClient()
}

// CallOpts returns a non-pointer bind.CallOpts, in case that's necessary for some reason...
func (f *Facet) CallOpts() bind.CallOpts {
	opts := f.Diamond.CallOpts()
	if opts != nil {
		return *opts
	}

	return bind.CallOpts{}
}
