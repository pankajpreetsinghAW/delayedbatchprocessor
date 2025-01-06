package delayedbatchprocessor

import (
	"context"
	"time"

	"log/slog"

	"go.opentelemetry.io/collector/pdata/pmetric"
)

func processMetrics(
	ctx context.Context,
	pm pmetric.Metrics,
) (mt pmetric.Metrics, err error) {
	slog.Info("**************** Going to sleep..")
	time.Sleep(1 * time.Minute)
	slog.Info("**************** coming out of sleep..")
	return mt, nil

}
