# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /v2-multinode-hardfork
WORKDIR /v2-multinode-hardfork
RUN go build .

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /v2-multinode-hardfork/v2-multinode-hardfork /v2-multinode-hardfork
ENTRYPOINT ["./v2-multinode-hardfork"]
