# shortener
WIP: Shortener API written in GO using Redis, gRPC and optional RabbitMQ to send events

## Requirements

* go
* Redis (disable keys eviction when RAM is full if you want to keep old keys!)
* (optional) RabbitMQ

## Running

### HTTP server

HTTP server only makes redirections. Build `main.go` file in http_server folder and run with needed environment variables.

### gRPC Server

gRPC api is used to manage urls. Build `main.go` file in grpc_server folder and run with needed environment variables.

