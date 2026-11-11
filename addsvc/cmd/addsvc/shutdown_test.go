package main

import (
	"bytes"
	"strings"
	"syscall"
	"testing"

	"github.com/go-kit/kit/log"
	"os"
)

func TestWaitForInterruptLogsShutdown(t *testing.T) {
	var buf bytes.Buffer
	c := make(chan os.Signal, 1)
	c <- syscall.SIGTERM

	err := waitForInterrupt(log.NewLogfmtLogger(&buf), c, make(chan struct{}))
	if err == nil || !strings.Contains(err.Error(), "terminated") {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf.String(), "graceful shutdown started") {
		t.Fatalf("missing shutdown log: %q", buf.String())
	}
}
