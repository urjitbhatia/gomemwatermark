# gomemwatermark
A High-memory watermark implementation for Go

## Why

Ideally, we'd like to build applications that can safely run without crashing due to OOMs.

This lib was initially written as part of ChronoMQ - https://github.com/chronomq/chronomq/tree/master/internal/monitor and extracted out so that others can use it in their projects.

The idea behind this implementation comes from looking at RabbitMQ implements a `high-memory watermark` which lets it safely apply backpressue to its publishers.


## The Big Picture/When to use this

This is ideal for applications where you can measure the objects your application stores in memory easily and you run more or less stateful loads.

The prime example is [ChronoMQ](https://github.com/chronomq/chronomq/tree/master/internal/monitor), which is where this idea germinated.
ChronoMQ is an in-memory, high-throughput queue that orders jobs globaly by time. It'd be bad idea for a service like this to crash.
By adding a safe way to apply back-pressure to the job producers, the service is able to maximize its memory utilization while staying safe.
When the service hits the high-watermark, it blocks producers, hoping for consumers to pick up the pace and clear some space.
As long as the produce/consume workloads are more or less balanced against the given memory, this lets the service operators deploy a service that won't OOM.

Another example can be implementing an http service that does some long-running operations on some objects in memory, letting the server
reject new requests if it is in danger or running out of memory.

This lib is also suitable for batch data processing use cases - anywhere that has a the ability to apply backpressure.

## Usage:

For ChronoMQ usage, check this implementation: https://github.com/chronomq/chronomq/blob/ac6a617562cfd24c4fc1f4ac381e13d8c6c24484/pkg/protocol/rpcserver.go#L32
