# Integration Connector Agent Custom Processor Template <!-- omit from toc -->

This repository contains a custom processor template for the Integration Connector Agent. It is designed to be used as a starting point for creating custom processors that can handle specific integration tasks.

## Table of Contents <!-- omit from toc -->

- [About Integration Connector Agent](#about-integration-connector-agent)
- [Getting Started](#getting-started)
- [Writing Your Custom Processor](#writing-your-custom-processor)
  - [Core Interface](#core-interface)
  - [Implementation Steps](#implementation-steps)
- [Important Configuration Points](#important-configuration-points)
  - [Module Path in Dockerfile](#module-path-in-dockerfile)
  - [Temporary Directory Access](#temporary-directory-access)
  - [Logging](#logging)
- [Deployment](#deployment)

## About Integration Connector Agent

The Integration Connector Agent is a powerful tool for handling data integration workflows. If you have special requirements that cannot be met by the standard processors, you can create custom processors using this template.

Custom processors are written in **Go** and are packaged as RPC plugins that communicate with the Integration Connector Agent. This repository template will build a Docker image that includes the Integration Connector Agent with your custom-built processor already integrated and ready to use.

## Getting Started

To get started with this template, follow these steps:

1. **Clone the Repository**: Clone this repository to your local machine
2. **Implement the Processor Logic**: Modify the `processor.go` file to implement your custom processing logic.
3. **Build the Processor**: Use the provided Makefile to build your custom processor.
4. **Test the Processor**: Ensure that your processor works as expected by running tests.
5. **Deploy the Processor**: Once tested, deploy your processor in your Project environment.

This template will build a fully functioning integration connector agent image with your custom processor included.

## Writing Your Custom Processor

Your custom processor must implement the `InitializableProcessor` interface defined by the [Integration Connector Agent](https://github.com/mia-platform/integration-connector-agent/blob/main/entities/processor.go#L26). The template provides a basic structure that you need to customize:

### Core Interface

The `CustomProcessor` struct must implement two main methods:

```go
type CustomProcessor struct {
    logger rpcprocessor.Logger
    config []byte
}

// Process handles the main data transformation logic
func (g *CustomProcessor) Process(input entities.PipelineEvent) (entities.PipelineEvent, error) {
    // Your custom processing logic here
    output := input.Clone()
    
    // Transform the data according to your requirements
    // Example: output.WithData([]byte(`{"data":"processed by CustomProcessor"}`))
    
    g.logger.Trace("CustomProcessor successfully processed the input event")
    return output, nil
}

// Init is called when the processor is initialized with configuration
func (g *CustomProcessor) Init(config []byte) error {
    g.config = config
    
    // Initialize your processor with the provided configuration
    // Parse config, set up connections, etc.
    
    g.logger.Info("CustomProcessor initialized")
    return nil
}
```

### Implementation Steps

1. **Modify the `Process` method**: This is where your main data transformation logic goes. The method receives a `PipelineEvent` and should return a transformed `PipelineEvent`.

2. **Customize the `Init` method**: Use this method to initialize your processor with any configuration data, establish database connections, or set up other resources, these configuration are provided by you in the processor `initOptions` written in the configuration file.

## Important Configuration Points

### Module Path in Dockerfile

The RPC plugin processor configuration requires a module path to be provided. Ensure it matches the path you used in your Dockerfile while building the image.

This template uses `/var/run/processor` as target path for the binary, so make sure to use it in your configuration. You can always change it to your needs, just make sure you update both the Dockerfile and the configuration accordingly.

### Temporary Directory Access

Your custom processor will need access to `/tmp` for temporary file operations. Make sure your deployment configuration includes an empty directory mount for this path:

Make sure that an EmptyDir volume is mounted to `/tmp` in your configuration. You can use the [EmptyDir feature in Console](https://docs.mia-platform.eu/docs/development_suite/api-console/api-design/services#empty-dirs) to manage it.

### Logging

The processor comes with a built-in logger that supports different log levels:

```go
g.logger.Trace("Detailed debugging information")
g.logger.Debug("Debug information")
g.logger.Info("General information")
g.logger.Warn("Warning messages")
g.logger.Error("Error messages")
```

## Deployment

Once you have built and tested your custom processor:

1. **Push the Docker image** to your container registry
2. **Configure the Integration Connector Agent** to use your custom processor
3. **Deploy** in your Project environment with proper volume mounts for `/tmp`
4. **Monitor** the logs to ensure your processor is working correctly

The resulting image contains both the Integration Connector Agent and your custom processor, ready to handle your specific integration requirements.
