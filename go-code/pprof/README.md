# Pprof

A Profile is a collection of stack traces showing the call sequences that led to instances of a particular event, such as allocation. Package can create and maintain their own profiles; the most common use is for tracking resources that must be explicitly closed, such as files or network connections.

Each Profile has a unique name. A few profiles are predefined:

|Profile|Description|
|-|-|
|goroutine|stack traces of all current goroutines|
|heap|a sampling of all heap allocations|
|threadcreate|stack traces that led to the creation of new OS threads|
|block|stack traces that led to blocking on synchronization primitives|
|mutex|stack traces of holders of contended mutexes|

Prfiling is a program optimization technique. "To profile a program", means to collect detailed statistics about how a program runs. Those statistics can be CPU usage, memory allocation, time spent on a program routines, number of function calls ... etc.

## Getting a heap profile with pprof

```go
```

## alloc_space vs inuse_space

go tool pprof has the option to show you either allocation counts of in use memory.

## CPU time

The CPU time represents the Central Processing Unit's time to execute the set of instructions defined in your program. The more your program is complex and makes intensive calculations, the more CPU time you need.

We can split CPU time into two subcategories:

- CPU user time
- CPU system time

## Kernel

In this section, we used the term "kernel". Kernel refers to the central component of an os. The kernel manages the system resources. It also manages the different hardware components. When we are doing a system call, we use the kernel facilities. For instance, opening a file in Go will trigger a system call that the kernel will handle.

## Diganostic

Diagnostics solutions can be categorized into the following groups:

- Profiling: Profiling tools analyze the complexity and costs of a Go program such as its memory usage and frequently called functions to identify the expensive sections of a Go program.

- Tracing: Tracing is a way to instrument code to analyze latency throughput the lifecycle of a call or user request. Traces provide an overview of how much latency each component contributes to the orverall latency in a system. Traces can span multiple Go processes.

- Debugging: Debugging allows us to pause a Go program and examine its execution. Program state and flow can be verified with debugging.

- Runtime statistics and events

