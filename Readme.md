# Train Ticket Service

This project implements a gRPC service for purchasing train tickets, querying receipts, managing seat assignments, and modifying or removing user bookings.

## Overview

The service allows clients to perform the following operations:

- Purchase a train ticket from London to France.
- Retrieve the details of the purchased ticket.
- View which users are seated in a particular section of the train.
- Modify the seat assignment for a user.
- Remove a user from the train booking system.

## Getting Started

### Prerequisites

- Go 1.21 or later
- Protocol Buffer Compiler (protoc)
- gRPC Go plugins for the Protocol Compiler

## Running the service

```
go mod vendor
go run cmd/main.go
```

## Testing the service

### Unit testing
```
cd service
go test -cover .
```

### Integration testing
#### Start running server
```
go run cmd/main.go
```

#### Start running integration client (In separate terminal)
```
go run client/main.go
```


