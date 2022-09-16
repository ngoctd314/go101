# Opentelemetry

In order to make a system observable, it must be instrumented. That is, the code must emit traces, metrics and logs. The instrumented data must then be sent to an Observability back-end.

## What can OpenTelemetry do for me?

- A single, vendor-agnostic instrumentation library per language with support for both automatic and manual instrumentation
- A single vendor-neutral collector binary that can be deployed in a variety of ways
- An end-to-end implementation to generate, emit, collect process and export telemetry data

## Signals

### Traces

To understand how tracing in OpenTelemetry works, let's look at a list of components that will play a part in instrumenting our code:

- Trace
- Trace Provider
- Trace Exporter
- Trace Context

#### Tracer Provider

A Tracer Provider is a factory for Tracers. In most applications, a Tracer Provider is initialized once and its lifecycle matches the application's lifecycle. Tracer Provider initialization also includes Resource and Exporter initialization.

#### Tracer

A Tracer creates spans containing more information about what is happening for a given operation, such as a request in a service. Tracers are created from Tracer Providers

#### Trace Exporters
Trace Exporters send traces to a consumer.

#### Trace Context
Trace Context is metadata about trace spans that provides correlation between spans across service and process boundaries.

#### Context Propagation
Context Propagation is the core concept that enables Distributed Tracing. With Context Propagation, Spans can be correlated with each other and assembled into a trace, regardless of where Spans are generated. 

A **Context** is an object contains the information for the sending and receiving to correlate one span with another and associate it with the trace overrall  

**Propagation** is the mechanism that moves Context between services and processes.

We identify Span Context using four major components: a traceID, and spanID, Trace Flags and Trace State.

**traceID** - A unique 16-byte array to identify the trace that a span is associated with
**spanID** - Hex-encoded 8-byte array to identify the current span
**Trace Flags** - Provides more details about the trace, such as if it is sampled
**Trace State** - Provides more vendor-specific information for tracing across multiple distributed systems

### Spans in OTel

A Span represents a unit of work or operation. Spans are the building blocks of Traces. In OpenTelemetry, they include the following information:

- Name
- Start and End Timestamps
- Span Context
- Attributes
- Events
- Links
- Status

#### Span Context
Span Context is an immutable object on every span that contains the following:
- The Trace ID representing the trace that the span is a part of
- The Span's Span ID

### Logs

### Metrics

### Baggage

Baggage refers to contextual information that's passed between spans

Imagine you want to have a CustomerId attribute on every span in your trace, which involves multiple services; however, CustomerId is only available in one specific service. To accomplish your goal, you can use OpenTelemetry Baggage to propagate this value across your system.

In OpenTelemetry, "Baggage" refers to contextual information that's passed between spans. It's a k-v store that resides within a trace context, making values available to any span created within that trace.

#### Why does OTel Baggage exist?

OpenTelemetry is cross-platform and cross-framework. Baggage makes it such that the context values live in the same place, have the same format, and follow the same pattern.

#### What should OTel Baggage be used for?

OTel Baggage should be used for non-sensitive data that you're okay with potentially exposing to third parties.

## Components

## Instrumenting

## Data Collection
