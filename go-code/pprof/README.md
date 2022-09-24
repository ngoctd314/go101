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

## Getting a heap profile with pprof

```go
```

## alloc_space vs inuse_space

go tool pprof has the option to show you either allocation counts of in use memory.