package rpchandlers

import (
	"github.com/bitmeme-taxi/bitmemed/app/appmessage"
	"github.com/bitmeme-taxi/bitmemed/app/rpc/rpccontext"
	"github.com/bitmeme-taxi/bitmemed/infrastructure/network/netadapter/router"
)

// HandleGetHeaders handles the respectively named RPC command
func HandleGetHeaders(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetHeadersResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
