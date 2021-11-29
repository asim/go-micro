# Go Micro [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/go-micro.dev/v4?tab=doc)

Go Micro is a framework for distributed systems development.

## Overview

Go Micro provides the core requirements for distributed systems development including RPC and Event driven communication. 
The **Micro** philosophy is sane defaults with a pluggable architecture. We provide defaults to get you started quickly 
but everything can be easily swapped out. 

## Features

Go Micro abstracts away the details of distributed systems. Here are the main features.

- **Authentication** - Auth is built in as a first class citizen. Authentication and authorization enable secure 
zero trust networking by providing every service an identity and certificates. This additionally includes rule 
based access control.

- **Dynamic Config** - Load and hot reload dynamic config from anywhere. The config interface provides a way to load application 
level config from any source such as env vars, file, etcd. You can merge the sources and even define fallbacks.

- **Data Storage** - A simple data store interface to read, write and delete records. It includes support for memory, file and 
CockroachDB by default. State and persistence becomes a core requirement beyond prototyping and Micro looks to build that into the framework.

- **Service Discovery** - Automatic service registration and name resolution. Service discovery is at the core of micro service 
development. When service A needs to speak to service B it needs the location of that service. The default discovery mechanism is 
multicast DNS (mdns), a zeroconf system.

- **Load Balancing** - Client side load balancing built on service discovery. Once we have the addresses of any number of instances 
of a service we now need a way to decide which node to route to. We use random hashed load balancing to provide even distribution 
across the services and retry a different node if there's a problem. 

- **Message Encoding** - Dynamic message encoding based on content-type. The client and server will use codecs along with content-type 
to seamlessly encode and decode Go types for you. Any variety of messages could be encoded and sent from different clients. The client 
and server handle this by default. This includes protobuf and json by default.

- **RPC Client/Server** - RPC based request/response with support for bidirectional streaming. We provide an abstraction for synchronous 
communication. A request made to a service will be automatically resolved, load balanced, dialled and streamed.

- **Async Messaging** - PubSub is built in as a first class citizen for asynchronous communication and event driven architectures. 
Event notifications are a core pattern in micro service development. The default messaging system is a HTTP event message broker.

- **Event Streaming** - PubSub is great for async notifications but for more advanced use cases event streaming is preferred. Offering 
persistent storage, consuming from offsets and acking. Go Micro includes support for NATS Jetstream and Redis streams.

- **Synchronization** - Distributed systems are often built in an eventually consistent manner. Support for distributed locking and 
leadership are built in as a Sync interface. When using an eventually consistent database or scheduling use the Sync interface.

- **Pluggable Interfaces** - Go Micro makes use of Go interfaces for each distributed system abstraction. Because of this these interfaces 
are pluggable and allows Go Micro to be runtime agnostic. You can plugin any underlying technology.

## Getting Started

To make use of Go Micro

```go
import "go-micro.dev/v4"

// create a new service
service := micro.NewService(
micro.Name("helloworld"),
)

// initialise flags
service.Init(
    BeforeStart(func() error {
        // Processing operations prior to service start.
    }),
    AfterStart(func() error {
        // Handling operations after service start.
    }),
    micro.BeforeStop(func() error {
        // Handling operations before service stoppage.
    }),
    micro.AfterStop(func() error {
        // Handling operations after a service stop.
    })
)

// start the service
service.Run()
```

For more startup configuration items please see [options](./options.go)

See the [examples](https://github.com/micro/go-micro/tree/master/examples) for detailed information on usage.

## Command Line Interface

See [cmd/micro](https://github.com/asim/go-micro/tree/master/cmd/micro) for the command line interface.

## Code Generation

See [cmd/protoc-gen-micro](https://github.com/micro/go-micro/tree/master/cmd/protoc-gen-micro) for protobuf code generation.

## Example Usage

See [examples](https://github.com/micro/go-micro/tree/master/examples) directory for usage examples.

## Plugins

See [plugins](https://github.com/micro/go-micro/tree/master/plugins) directory for all the plugins.

## Services

See [services](https://github.com/micro/go-micro/tree/master/services) directory for third party services.

## License

Go Micro is Apache 2.0 licensed.
