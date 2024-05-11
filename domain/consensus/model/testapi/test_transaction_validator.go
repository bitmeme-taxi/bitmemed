package testapi

import (
	"github.com/bitmeme-taxi/bitmemed/domain/consensus/model"
	"github.com/bitmeme-taxi/bitmemed/domain/consensus/utils/txscript"
)

// TestTransactionValidator adds to the main TransactionValidator methods required by tests
type TestTransactionValidator interface {
	model.TransactionValidator
	SigCache() *txscript.SigCache
	SetSigCache(sigCache *txscript.SigCache)
}
