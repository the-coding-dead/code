FROM golang:1.15-alpine3.12 as build

RUN apk add --no-cache gcc \
                       git \
                       libc-dev

WORKDIR /shield

COPY go.mod .
COPY go.sum .

RUN go mod download
