package common

import (
	"fmt"
	"github.com/bitmeme-taxi/bitmemed/domain/dagconfig"
	"os"
	"sync/atomic"
	"syscall"
	"testing"
)

// RunBitmemeForTesting runs bitmemed for testing purposes
func RunKaspadForTesting(t *testing.T, testName string, rpcAddress string) func() {
	appDir, err := TempDir(testName)
	if err != nil {
		t.Fatalf("TempDir: %s", err)
	}

	kaspadRunCommand, err := StartCmd("BITMEMED",
		"bitmemed",
		NetworkCliArgumentFromNetParams(&dagconfig.DevnetParams),
		"--appdir", appDir,
		"--rpclisten", rpcAddress,
		"--loglevel", "debug",
	)
	if err != nil {
		t.Fatalf("StartCmd: %s", err)
	}
	t.Logf("Bitmemed started with --appdir=%s", appDir)

	isShutdown := uint64(0)
	go func() {
		err := kaspadRunCommand.Wait()
		if err != nil {
			if atomic.LoadUint64(&isShutdown) == 0 {
				panic(fmt.Sprintf("Bitmemed closed unexpectedly: %s. See logs at: %s", err, appDir))
			}
		}
	}()

	return func() {
		err := kaspadRunCommand.Process.Signal(syscall.SIGTERM)
		if err != nil {
			t.Fatalf("Signal: %s", err)
		}
		err = os.RemoveAll(appDir)
		if err != nil {
			t.Fatalf("RemoveAll: %s", err)
		}
		atomic.StoreUint64(&isShutdown, 1)
		t.Logf("Bitmemed stopped")
	}
}
