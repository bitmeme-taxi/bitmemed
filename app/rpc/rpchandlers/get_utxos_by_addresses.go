package rpchandlers

import (
	"github.com/bitmeme-taxi/bitmemed/app/appmessage"
	"github.com/bitmeme-taxi/bitmemed/app/rpc/rpccontext"
	"github.com/bitmeme-taxi/bitmemed/domain/consensus/utils/txscript"
	"github.com/bitmeme-taxi/bitmemed/infrastructure/network/netadapter/router"
	"github.com/bitmeme-taxi/bitmemed/util"
)

// HandleGetUTXOsByAddresses handles the respectively named RPC command
func HandleGetUTXOsByAddresses(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetUTXOsByAddressesResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when kaspad is run without --utxoindex")
		return errorMessage, nil
	}

	getUTXOsByAddressesRequest := request.(*appmessage.GetUTXOsByAddressesRequestMessage)

	allEntries := make([]*appmessage.UTXOsByAddressesEntry, 0)
	for _, addressString := range getUTXOsByAddressesRequest.Addresses {
		address, err := util.DecodeAddress(addressString, context.Config.ActiveNetParams.Prefix)
		if err != nil {
			errorMessage := &appmessage.GetUTXOsByAddressesResponseMessage{}
			errorMessage.Error = appmessage.RPCErrorf("Could not decode address '%s': %s", addressString, err)
			return errorMessage, nil
		}
		scriptPublicKey, err := txscript.PayToAddrScript(address)
		if err != nil {
			errorMessage := &appmessage.GetUTXOsByAddressesResponseMessage{}
			errorMessage.Error = appmessage.RPCErrorf("Could not create a scriptPublicKey for address '%s': %s", addressString, err)
			return errorMessage, nil
		}
		utxoOutpointEntryPairs, err := context.UTXOIndex.UTXOs(scriptPublicKey)
		if err != nil {
			return nil, err
		}
		entries := rpccontext.ConvertUTXOOutpointEntryPairsToUTXOsByAddressesEntries(addressString, utxoOutpointEntryPairs)
		allEntries = append(allEntries, entries...)
	}

	response := appmessage.NewGetUTXOsByAddressesResponseMessage(allEntries)
	return response, nil
}
