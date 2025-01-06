package delayedbatchprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

const (
	// typeStr is the type of the processor
	typeStr = "delayedbatchprocessor"
)

// NewFactory creates a processor factory
func NewFactory() processor.Factory {
	return processor.NewFactory(
		component.MustNewType(typeStr),
		createDefaultConfig,
		// Uncomment the processor type that you would like, change the second parameter as you like
		// component.StabilityLevelUndefined
		// component.StabilityLevelUnmaintained
		// component.StabilityLevelDeprecated
		// component.StabilityLevelDevelopment
		// component.StabilityLevelAlpha
		// component.StabilityLevelBeta
		// component.StabilityLevelStable
		processor.WithTraces(createTracesProcessor, component.StabilityLevelBeta),

		processor.WithLogs(createLogsProcessor, component.StabilityLevelAlpha),

		processor.WithMetrics(createMetricsProcessor, component.StabilityLevelBeta),
	)
}

func createDefaultConfig() component.Config {

	return &config{}
}

// createTracesProcesor creates a trace processor based on this config.
func createTracesProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {

	return processorhelper.NewTraces(
		ctx,
		set,
		cfg,
		nextConsumer,
		processTraces,
		//	The parameters below are optional. Uncomment any as you need.
		//	processorhelper.WithStart(start component.StartFunc),
		//processorhelper.WithShutdown(shutdown component.ShutdownFunc),
		//processorhelper.WithCapabilities(capabilities consumer.Capabilities)
	)

}
func createLogsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (processor.Logs, error) {

	return processorhelper.NewLogs(
		ctx,
		set,
		cfg,
		nextConsumer,
		processLogs,
		//	The parameters below are optional. Uncomment any as you need.
		//	processorhelper.WithStart(start component.StartFunc),
		//processorhelper.WithShutdown(shutdown component.ShutdownFunc),
		//processorhelper.WithCapabilities(capabilities consumer.Capabilities)
	)

}

// createMetricsProcessor creates a metrics processor based on this config.
func createMetricsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (processor.Metrics, error) {

	return processorhelper.NewMetrics(
		ctx,
		set,
		cfg,
		nextConsumer,
		processMetrics,
		//	The parameters below are optional. Uncomment any as you need.
		//	processorhelper.WithStart(start component.StartFunc),
		//processorhelper.WithShutdown(shutdown component.ShutdownFunc),
		//processorhelper.WithCapabilities(capabilities consumer.Capabilities)
	)

}
