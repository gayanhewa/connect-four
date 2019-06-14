FROM golang:alpine AS builder
WORKDIR /go/src/github.com/gayanhewa/connect-four
COPY . /go/src/github.com/gayanhewa/connect-four
RUN cd /go/src/github.com/gayanhewa/connect-four && go build -o c4

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/github.com/gayanhewa/connect-four/c4 /app/c4
ENTRYPOINT ./c4