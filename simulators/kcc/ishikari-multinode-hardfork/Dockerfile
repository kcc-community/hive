# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /ishikari-multinode-hardfork
WORKDIR /ishikari-multinode-hardfork
RUN go build .

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /ishikari-multinode-hardfork/ishikari-multinode-hardfork /ishikari-multinode-hardfork
ENTRYPOINT ["./ishikari-multinode-hardfork"]
