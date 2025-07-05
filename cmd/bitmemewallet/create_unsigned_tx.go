package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bitmeme-taxi/bitmemed/cmd/gorwallet/daemon/client"
	"github.com/bitmeme-taxi/bitmemed/cmd/gorwallet/daemon/pb"
	"github.com/bitmeme-taxi/bitmemed/cmd/gorwallet/daemon/server"
	"github.com/bitmeme-taxi/bitmemed/cmd/gorwallet/utils"
)

func createUnsignedTransaction(conf *createUnsignedTransactionConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()

	var sendAmountSompi uint64

	if !conf.IsSendAll {
		sendAmountSompi, err = utils.KasToSompi(conf.SendAmount)
		if err != nil {
			return err
		}
	}

	var feePolicy *pb.FeePolicy
	if conf.FeeRate > 0 {
		feePolicy = &pb.FeePolicy{
			FeePolicy: &pb.FeePolicy_ExactFeeRate{
				ExactFeeRate: conf.FeeRate,
			},
		}
	} else if conf.MaxFeeRate > 0 {
		feePolicy = &pb.FeePolicy{
			FeePolicy: &pb.FeePolicy_MaxFeeRate{MaxFeeRate: conf.MaxFeeRate},
		}
	} else if conf.MaxFee > 0 {
		feePolicy = &pb.FeePolicy{
			FeePolicy: &pb.FeePolicy_MaxFee{MaxFee: conf.MaxFee},
		}
	}

	response, err := daemonClient.CreateUnsignedTransactions(ctx, &pb.CreateUnsignedTransactionsRequest{
		From:                     conf.FromAddresses,
		Address:                  conf.ToAddress,
		Amount:                   sendAmountSompi,
		IsSendAll:                conf.IsSendAll,
		UseExistingChangeAddress: conf.UseExistingChangeAddress,
		FeePolicy:                feePolicy,
	})
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stderr, "Created unsigned transaction")
	fmt.Println(server.EncodeTransactionsToHex(response.UnsignedTransactions))

	return nil
}
