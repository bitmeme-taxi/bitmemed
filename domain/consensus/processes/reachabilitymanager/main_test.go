package reachabilitymanager_test

import (
	"os"
	"testing"

	"github.com/bitmeme-taxi/bitmemed/infrastructure/logger"
)

const logLevel = logger.LevelWarn

func TestMain(m *testing.M) {
	logger.SetLogLevels(logLevel)
	logger.InitLogStdout(logLevel)
	os.Exit(m.Run())
}
