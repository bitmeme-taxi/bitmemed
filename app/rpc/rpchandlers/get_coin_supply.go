package rpchandlers

import (
	"github.com/bitmeme-taxi/bitmemed/app/appmessage"
	"github.com/bitmeme-taxi/bitmemed/app/rpc/rpccontext"
	"github.com/bitmeme-taxi/bitmemed/domain/consensus/utils/constants"
	"github.com/bitmeme-taxi/bitmemed/infrastructure/network/netadapter/router"
)

// HandleGetCoinSupply handles the respectively named RPC command
func HandleGetCoinSupply(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetCoinSupplyResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when bitmemed is run without --utxoindex")
		return errorMessage, nil
	}

	circulatingSompiSupply, err := context.UTXOIndex.GetCirculatingSompiSupply()
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetCoinSupplyResponseMessage(
		constants.MaxSompi,
		circulatingSompiSupply,
	)

	return response, nil
}
