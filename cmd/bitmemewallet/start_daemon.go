package main

import "github.com/bitmeme-taxi/bitmemed/cmd/bitmemewallet/daemon/server"

func startDaemon(conf *startDaemonConfig) error {
	return server.Start(conf.NetParams(), conf.Listen, conf.RPCServer, conf.KeysFile, conf.Profile, conf.Timeout)
}
