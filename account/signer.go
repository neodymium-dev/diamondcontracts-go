package account

import (
	"crypto/ecdsa"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Signer struct {
	address    common.Address
	publicKey  ecdsa.PublicKey
	privateKey *ecdsa.PrivateKey
}

func (s Signer) toMarshalable() marshalableSigner {
	m := marshalableSigner{}
	if s.privateKey != nil {
		m.Address = s.address.String()
		m.PrivateKey = crypto.FromECDSA(s.privateKey)
	}

	return m
}

type marshalableSigner struct {
	Address    string `json:"address" yaml:"address"`
	PrivateKey []byte `json:"private_key" yaml:"private_key"`
}

func (s Signer) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.toMarshalable())
}

func (s *Signer) UnmarshalJSON(data []byte) error {
	var raw marshalableSigner

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if raw.Address != "" {
		s.address = common.HexToAddress(raw.Address)
	}

	if raw.PrivateKey != nil && len(raw.PrivateKey) > 0 {
		privkey, err := crypto.ToECDSA(raw.PrivateKey)
		if err != nil {
			return err

		}

		s.privateKey = privkey
		s.publicKey = s.privateKey.PublicKey
	}

	return nil
}

func NewSigner() *Signer {
	return &Signer{}
}

func NewSignerFromPrivateKey(privateKey *ecdsa.PrivateKey) *Signer {
	return &Signer{
		address:    crypto.PubkeyToAddress(privateKey.PublicKey),
		publicKey:  privateKey.PublicKey,
		privateKey: privateKey,
	}
}

func NewSignerFromPrivateKeyHex(privateKey string) (*Signer, error) {
	panic("not implemented")
}
