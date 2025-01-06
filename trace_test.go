package delayedbatchprocessor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor/processortest"
)

func TestTracesProcessorNoErrors(t *testing.T) {
	f := NewFactory()
	lte, err := f.CreateTracesProcessor(context.Background(), processortest.NewNopCreateSettings(), f.CreateDefaultConfig(), nil)
	require.NotNil(t, lte)
	assert.NoError(t, err)

	assert.NoError(t, lte.ConsumeTraces(context.Background(), ptrace.NewTraces()))

	assert.NoError(t, lte.Shutdown(context.Background()))
}
