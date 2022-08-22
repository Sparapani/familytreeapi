# FamilyTree: multi stage build

FROM golang:1.19-alpine as builder
# Install go toolchain
RUN apk update && apk add --no-cache curl gcc git libc-dev
# revive
RUN go install github.com/mgechev/revive@latest && \
    revive ./...
# gosec - Golang Security Checker
RUN curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b ${GOPATH}/bin latest && \
    gosec ./...
# go test - Golang test unity and benchmarks
RUN go test ./... -v
# Use a Docker multi-stage build
FROM alpine:3.11