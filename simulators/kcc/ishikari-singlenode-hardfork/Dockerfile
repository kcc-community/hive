# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /ishikari-singlenode-hardfork
WORKDIR /ishikari-singlenode-hardfork
RUN go build .

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /ishikari-singlenode-hardfork/ishikari-singlenode-hardfork /ishikari-singlenode-hardfork
ENTRYPOINT ["./ishikari-singlenode-hardfork"]
