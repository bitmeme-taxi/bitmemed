package mempool

import (
	"github.com/bitmeme-taxi/bitmemed/domain/consensus/model/externalapi"
	"github.com/bitmeme-taxi/bitmemed/domain/miningmanager/mempool/model"
	"github.com/bitmeme-taxi/bitmemed/infrastructure/logger"
)

func (mp *mempool) revalidateHighPriorityTransactions() ([]*externalapi.DomainTransaction, error) {


	onEnd := logger.LogAndMeasureExecutionTime(log, "revalidateHighPriorityTransactions")
	defer onEnd()



	validTransactions := []*externalapi.DomainTransaction{}
	
	
		isValid, err := mp.revalidateTransaction(transaction)
		if err != nil {
			return nil, err
		}
		if !isValid {
			continue
		}

		validTransactions = append(validTransactions, transaction.Transaction().Clone())
	}

	return validTransactions, nil
}

func (mp *mempool) revalidateTransaction(transaction *model.MempoolTransaction) (isValid bool, err error) {
	clearInputs(transaction)

	_, missingParents, err := mp.fillInputsAndGetMissingParents(transaction.Transaction())
	if err != nil {
		return false, err
	}
	if len(missingParents) > 0 {
		log.Debugf("Removing transaction %s, it failed revalidation", transaction.TransactionID())
		err := mp.removeTransaction(transaction.TransactionID(), true)
		if err != nil {
			return false, err
		}
		return false, nil
	}

	return true, nil
}

func clearInputs(transaction *model.MempoolTransaction) {
	for _, input := range transaction.Transaction().Inputs {
		input.UTXOEntry = nil
	}
}
