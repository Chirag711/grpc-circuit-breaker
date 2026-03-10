# gRPC Circuit Breaker Example (Go)

This project demonstrates how to implement the **Circuit Breaker pattern** in a gRPC-based microservice using Go.
The system simulates a payment service that sometimes fails, and a client that uses a circuit breaker to prevent repeated requests to a failing service.

The circuit breaker is implemented using the Go library **github.com/sony/gobreaker**.

---

## Overview

In distributed systems, repeatedly calling a failing service can cause system overload and cascading failures.
A circuit breaker detects repeated failures and temporarily stops sending requests to the service.

Flow:

Client → Payment Service (gRPC)

If failures exceed a threshold:

Client → Circuit Breaker → Requests blocked temporarily

This protects the system until the service becomes healthy again.

---

## Project Structure

```
grpc-circuit-breaker
│
├── go.mod
│
├── proto
│   └── payment.proto
│
├── pb
│   ├── payment.pb.go
│   └── payment_grpc.pb.go
│
├── server
│   └── main.go
│
└── client
    └── main.go
```

---

## Prerequisites

Before running the project, install the following:

* Go 1.20 or later
* Protocol Buffers compiler (protoc)
* gRPC Go plugins

Install required plugins:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

---

## Install Dependencies

Initialize the module:

```
go mod init grpc-circuit-breaker
```

Install required libraries:

```
go get google.golang.org/grpc
go get github.com/sony/gobreaker
```

---

## Proto Definition

File: `proto/payment.proto`

```
syntax = "proto3";

package payment;

option go_package = "grpc-circuit-breaker/pb";

service PaymentService {
  rpc ProcessPayment (PaymentRequest) returns (PaymentResponse);
}

message PaymentRequest {
  string orderId = 1;
  double amount = 2;
}

message PaymentResponse {
  string status = 1;
}
```

Generate Go code:

```
protoc --go_out=. --go-grpc_out=. proto/payment.proto
```

---

## Server

The server simulates a payment system that randomly fails to process requests.

File: `server/main.go`

The server randomly returns an error to simulate an unstable external service.

Example behavior:

* Some requests succeed
* Some requests fail
* Client must handle failures

---

## Client

The client calls the payment service and wraps the request with a **circuit breaker**.

File: `client/main.go`

The circuit breaker monitors failures and stops requests when the failure threshold is exceeded.

Behavior:

* First few failures are allowed
* If failures continue, the circuit opens
* Requests are temporarily blocked
* After a timeout, the circuit attempts recovery

---

## Running the Application

### Start the Server

```
go run ./server
```

Expected output:

```
Payment Server running...
```

---

### Run the Client

```
go run ./client
```

Example output:

```
Payment Status: Payment Successful
Circuit Breaker Triggered: bank service unavailable
Circuit Breaker Triggered: circuit breaker is open
Circuit Breaker Triggered: circuit breaker is open
Payment Status: Payment Successful
```

---

## Circuit Breaker States

Closed
All requests pass normally.

Open
Requests are blocked because the service has failed too many times.

Half-Open
A few requests are allowed to test if the service has recovered.

State transitions:

Closed → Open → Half-Open → Closed

---

## Concepts Demonstrated

* gRPC client-server communication
* Circuit breaker pattern
* Failure simulation
* Resilient microservice communication
* Fault tolerance

---

## Use Cases

Circuit breakers are commonly used in:

* Payment gateways
* External API integrations
* Microservice communication
* Database access layers
* Banking systems

---

## Possible Improvements

This project can be extended with:

* Timeout handling
* Retry mechanisms
* Load balancing
* Logging interceptors
* Metrics and monitoring
* Distributed tracing

