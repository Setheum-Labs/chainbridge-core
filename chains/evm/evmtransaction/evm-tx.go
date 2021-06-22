package evmtransaction

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type TX struct {
	*types.Transaction
}

// RawWithSignature mostly copies WithSignature interface of type.Transaction from go-ethereum,
// but return raw byte representation of transaction to be compatible and interchangeable between different go-ethereum forks
// WithSignature returns a new transaction with the given signature.
// This signature needs to be in the [R || S || V] format where V is 0 or 1.
func (a *TX) RawWithSignature(key *ecdsa.PrivateKey, chainID *big.Int) ([]byte, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, err
	}
	tx, err := opts.Signer(crypto.PubkeyToAddress(key.PublicKey), a.Transaction)
	if err != nil {
		return nil, err
	}
	rawTX, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return nil, err
	}
	log.Debug().Str("hash", a.Transaction.Hash().Hex()).Msg("HAAAAASH")
	return rawTX, nil
}

func NewTransaction(nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) *TX {
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)
	return &TX{tx}
}

func (a *TX) Hash() common.Hash {
	return a.Transaction.Hash()
}
