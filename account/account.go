package account

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Account struct {
	address     *common.Address
	signer      *Signer
	transactors map[uint64]*bind.TransactOpts
}

func (a Account) toMarshalable() marshalableAccount {
	m := marshalableAccount{}

	if a.address != nil {
		if a.signer != nil {
			m.Address = a.signer.address.String()
		} else {
			m.Address = (*a.address).String()
		}
	}

	if a.signer != nil {
		m.Signer = a.signer.toMarshalable()
	}

	return m
}

type marshalableAccount struct {
	Address string            `json:"address" yaml:"address"`
	Signer  marshalableSigner `json:"signer" yaml:"signer"`
}

func (a Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.toMarshalable())
}

func (a *Account) UnmarshalJSON(data []byte) error {
	var raw marshalableAccount

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if raw.Address != "" {
		addr := common.HexToAddress(raw.Address)
		a.address = &addr
	}

	if raw.Signer.PrivateKey != nil && len(raw.Signer.PrivateKey) > 0 {
		privkey, err := crypto.ToECDSA(raw.Signer.PrivateKey)
		if err != nil {
			return err

		}

		a.signer = NewSignerFromPrivateKey(privkey)
		a.address = &a.signer.address
	}

	return nil
}

func NewAccount() *Account {
	return &Account{}
}

func NewAccountWithAddress(address common.Address) *Account {
	return &Account{
		address: &address,
	}
}

func NewAccountFromSigner(signer *Signer) *Account {
	return &Account{
		address: &signer.address,
		signer:  signer,
	}
}

func NewAccountFromPrivateKey(privateKey *ecdsa.PrivateKey) *Account {
	signer := NewSignerFromPrivateKey(privateKey)

	return &Account{
		address: &signer.address,
		signer:  signer,
	}
}

func NewAccountFromPrivateKeyHex(privateKey string) (acc *Account, err error) {
	signer, err := NewSignerFromPrivateKeyHex(privateKey)
	if err != nil {
		return
	}

	acc = &Account{
		address: &signer.address,
		signer:  signer,
	}

	return
}

func (a Account) Address() (addr common.Address) {
	switch {
	case a.signer != nil:
		addr = a.signer.address
	case a.address != nil && a.signer == nil:
		addr = *a.address
	}

	return
}

func (a *Account) SetAddress(addr common.Address) {
	a.address = &addr
}

func (a *Account) Signer() *Signer {
	return a.signer
}

func (a *Account) SetSigner(signer *Signer) {
	a.signer = signer
}

func (a *Account) SetSignerFromPrivateKey(privateKey string) (err error) {
	a.signer, err = NewSignerFromPrivateKeyHex(privateKey)
	if err != nil {
		return
	}

	a.address = &a.signer.address

	return
}

func (a *Account) TransactOpts(chainID *big.Int) (*bind.TransactOpts, error) {
	return a.buildTransactOpts(chainID)
}

func (a *Account) TransactOptsContext(ctx context.Context, chainID *big.Int) (*bind.TransactOpts, error) {
	opts, err := a.buildTransactOpts(chainID)
	if err != nil {
		return nil, err
	}

	opts.Context = ctx

	return opts, nil
}

func (a *Account) buildTransactOpts(chainID *big.Int) (opts *bind.TransactOpts, err error) {
	if a.signer == nil {
		err = fmt.Errorf("account does not have a valid signer")
		return
	}

	var ok bool

	cid := chainID.Uint64()

	opts, ok = a.transactors[cid]
	if ok {
		return
	}

	opts, err = bind.NewKeyedTransactorWithChainID(a.signer.privateKey, chainID)
	if err != nil {
		return
	}

	a.transactors[cid] = opts

	return
}

func (a Account) CallOpts() *bind.CallOpts {
	return a.buildCallOpts()
}

func (a Account) CallOptsContext(ctx context.Context) *bind.CallOpts {
	opts := a.buildCallOpts()
	opts.Context = ctx

	return opts
}

func (a Account) buildCallOpts() (opts *bind.CallOpts) {
	opts = &bind.CallOpts{}

	if addr := a.Address(); addr.String() != "" {
		opts.From = addr
	}

	return
}
