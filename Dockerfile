FROM golang:1.18 AS builder

RUN mkdir -p /gp/src

WORKDIR /go/src
COPY .