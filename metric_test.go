package delayedbatchprocessor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/processor/processortest"
)

func TestMetricsProcessorNoErrors(t *testing.T) {
	f := NewFactory()
	lme, err := f.CreateMetricsProcessor(context.Background(), processortest.NewNopCreateSettings(), f.CreateDefaultConfig(), nil)
	require.NotNil(t, lme)
	assert.NoError(t, err)

	assert.NoError(t, lme.ConsumeMetrics(context.Background(), pmetric.NewMetrics()))

	assert.NoError(t, lme.Shutdown(context.Background()))
}
