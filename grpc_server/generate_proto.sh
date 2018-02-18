#!/bin/bash

protoc -I proto/ proto/shortener.proto --go_out=plugins=grpc:proto
