# Einride Interview: Backend Software Engineer

# Introduction

Welcome to your technical interview at Einride! We're excited to pair
program together with you.

Before the pair programming session, please complete this small
preparation exercise. During the pair programming, you will be the
[driver][pair-programming], and we will develop the code further together.

[pair-programming]: https://en.wikipedia.org/wiki/Pair_programming

# Background

This preparation exercise is designed to be representative of backend
development at Einride.

We use the [Go][go] programming language for our backend services, we use
[gRPC][grpc] for our APIs, and we use [protobuf][protobuf] for our domain
model.

[go]: https://golang.org
[grpc]: https://grpc.io/
[protobuf]: https://developers.google.com/protocol-buffers

To help us design consistent and easy-to-use APIs, we follow the
[AIP][aip] (API Improvement Proposals) design system for [gRPC][grpc]
APIs.

[aip]: https://google.aip.dev/

By following the [AIP][aip] design system, we are able to make use of
[Open Source API tools and libraries from Google][googleapis], such as an
[automatically generated CLI tool][gapic-generator-go] for talking to our
backend services, which we will make use of in this interview.

[googleapis]: https://github.com/googleapis
[gapic-generator-go]: https://github.com/googleapis/gapic-generator-go

# Pre-requisites

- A Linux or Mac development environment. If you only have access to a
  Windows development environment, we recommend using [WSL][wsl] or a
  [VirtualBox][virtualbox] with [Ubuntu 20.04][ubuntu].

  (At Einride, most developers use the latest Ubuntu LTS for their
  workstations).

- [Go >= 1.15](https://golang.org/dl/).

- A Go IDE. Einride provides all developers with [JetBrains
  Goland][goland] licenses, and there is a free 30-day trial that you can
  use during the interview.

[wsl]: https://docs.microsoft.com/en-us/windows/wsl/install-win10
[virtualbox]: https://www.virtualbox.org
[goland]: https://www.jetbrains.com/go/
[ubuntu]: https://releases.ubuntu.com/20.04/

# Task

In this preparation exercise, you will implement a small
[TodoMVC][todomvc]-style backend service using some of Einride's key
technologies (namely [Go][go], [gRPC][grpc], [protobuf][protobuf] and
[AIP][aip]).

[todomvc]: http://todomvc.com/

To help you get started, and to fast-track you through the boilerplate
project setup, we have provided you with a starting template service.

We have also provided you with the [full gRPC
API](./proto/src/einride/todo/v1/todo_service.proto) to implement,
including a [Makefile](./Makefile) with build scripts that handle all the
code generation for you.

This should let you get started right away on implementing the following
API methods:

- [CreateTodo][create]
- [GetTodo][get]
- [UpdateTodo][update]
- [DeleteTodo][delete]
- [ListTodos][list]

[get]: https://google.aip.dev/131
[create]: https://google.aip.dev/133
[update]: https://google.aip.dev/134
[delete]: https://google.aip.dev/135
[list]: https://google.aip.dev/132

Make sure to glance through the AIP design guide for the standard CRUD
methods to provide an implementation that is consistent with the
specification. The introductory part on [resource-oriented
design][resource-oriented-design] and the part on [resource
names][resource-names] may come in handy.

[resource-oriented-design]: https://google.aip.dev/121
[resource-names]: https://google.aip.dev/122

You are free to implement persistent storage however you want. In-memory
is fine.

As you implement the methods, you may notice that there is also a unit
test suite that needs updating.

When you have finished implementing the methods, you can use the helper
targets in the [Makefile](./Makefile) to start the service:

```bash
$ make run-server
[run-server] starting...
listening on localhost:8080...
```

We have provided a minimal [integration test
script](./scripts/integration-test.bash):

```bash
$ make integration-test
[integration-test] running integration tests...
```

The integration test script shows how to use the auto-generated CLI to
communicate with your service. This may sometimes come in handy for
debugging.

# Next steps

When you have completed the preparation exercise, you are all set for the
pair programming, where we will keep working on your code together, adding
more amazing features!

Make sure to familiarize yourself with the whole code base, including the
build scripts, as we might be adding more endpoints to the API.

# What we are looking for

The purpose of this interview is twofold:

- Give you a chance to try the technologies we use.

- Give us a chance to program together with you and have a conversation
  about building things with code.

We want to get a feel for how you think, talk and work as a software
engineer. How you approach a problem, how you reason about trade-offs, how
you communicate ideas, how you work with tools, and how you give and
receive feedback.

We place as much importance on the discussions during the interview as the
quality of the code you bring with you.

# Hope to see you soon!

We're looking forward to pair programming with you!
