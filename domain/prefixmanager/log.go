package prefixmanager

import (
	"github.com/bitmeme-taxi/bitmemed/infrastructure/logger"
	"github.com/bitmeme-taxi/bitmemed/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
