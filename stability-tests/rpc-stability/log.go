package main

import (
	"github.com/bitmeme-taxi/bitmemed/infrastructure/logger"
	"github.com/bitmeme-taxi/bitmemed/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("JSTT")
	spawn      = panics.GoroutineWrapperFunc(log)
)
