package delayedbatchprocessor

import (
	"context"
	"log/slog"
	"time"

	"go.opentelemetry.io/collector/pdata/plog"
)

func processLogs(
	ctx context.Context,
	lg plog.Logs,
) (pl plog.Logs, err error) {

	slog.Info("**************** Going to sleep..")
	time.Sleep(1 * time.Minute)
	slog.Info("**************** coming out of sleep..")
	return pl, nil

}
