# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.14.5-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh 
    

# Add Maintainer Info
LABEL maintainer="Eudes Vicente <eudes.vss@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/eudes16/go-sentinel

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

RUN env GO111MODULE=on go get github.com/cortesi/modd/cmd/modd
