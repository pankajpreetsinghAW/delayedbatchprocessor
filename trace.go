package delayedbatchprocessor

import (
	"context"
	"log/slog"
	"time"

	"go.opentelemetry.io/collector/pdata/ptrace"
)

func processTraces(
	ctx context.Context,
	pt ptrace.Traces,
) (tc ptrace.Traces, err error) {
	slog.Info("**************** Going to sleep..")
	time.Sleep(1 * time.Minute)
	slog.Info("**************** coming out of sleep..")
	return tc, nil

}
