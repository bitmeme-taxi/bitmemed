package main

import (
	"path/filepath"

	"github.com/jessevdk/go-flags"
	"github.com/bitmeme-taxi/bitmemed/infrastructure/config"
	"github.com/bitmeme-taxi/bitmemed/infrastructure/logger"
	"github.com/bitmeme-taxi/bitmemed/stability-tests/common"
	"github.com/bitmeme-taxi/bitmemed/stability-tests/common/rpc"
)

const (
	defaultLogFilename    = "rpc_idle_clients.log"
	defaultErrLogFilename = "rpc_idle_clients_err.log"
)

var (
	// Default configuration options
	defaultLogFile    = filepath.Join(common.DefaultAppDir, defaultLogFilename)
	defaultErrLogFile = filepath.Join(common.DefaultAppDir, defaultErrLogFilename)
)

type configFlags struct {
	rpc.Config
	config.NetworkFlags
	NumClients uint32 `long:"numclients" short:"n" description:"Number of RPC clients to open"`
	Profile    string `long:"profile" description:"Enable HTTP profiling on given port -- NOTE port must be between 1024 and 65536"`
}

var cfg *configFlags

func activeConfig() *configFlags {
	return cfg
}

func parseConfig() error {
	cfg = &configFlags{}

	parser := flags.NewParser(cfg, flags.PrintErrors|flags.HelpFlag)
	_, err := parser.Parse()

	if err != nil {
		return err
	}

	err = cfg.ResolveNetwork(parser)
	if err != nil {
		return err
	}

	err = rpc.ValidateRPCConfig(&cfg.Config)
	if err != nil {
		return err
	}
	log.SetLevel(logger.LevelInfo)
	common.InitBackend(backendLog, defaultLogFile, defaultErrLogFile)
	return nil
}
