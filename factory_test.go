package delayedbatchprocessor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/processor/processortest"
)

func TestCreateDefaultConfig(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, componenttest.CheckConfigStruct(cfg))
}
func TestCreateTracesProcessor(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()

	te, err := factory.CreateTracesProcessor(context.Background(), processortest.NewNopCreateSettings(), cfg, nil)
	assert.NoError(t, err)
	assert.NotNil(t, te)
}
func TestCreateLogsProcessor(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()

	te, err := factory.CreateLogsProcessor(context.Background(), processortest.NewNopCreateSettings(), cfg, nil)
	assert.NoError(t, err)
	assert.NotNil(t, te)
}
func TestCreateMetricsProcessor(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()

	me, err := factory.CreateMetricsProcessor(context.Background(), processortest.NewNopCreateSettings(), cfg, nil)
	assert.NoError(t, err)
	assert.NotNil(t, me)
}