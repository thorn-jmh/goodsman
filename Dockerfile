FROM golang:1.18 AS builder

ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,direct \
    GO111MODULE=on \
    WORKDIR=/tmp/src/ \
    CGO_ENABLED=0

RUN mkdir -p $WORKDIR

COPY . $WORKDIR

RUN cd $WORKDIR && go mod download all

RUN cd $WORKDIR && go build -o /goodsman

FROM alpine:3.15.2

COPY --from=builder /goodsman /goodsman

CMD ["/goodsman"]