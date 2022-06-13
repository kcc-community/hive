# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /ishikari-distribute-block-reward
WORKDIR /ishikari-distribute-block-reward
RUN go build .

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /ishikari-distribute-block-reward/ishikari-distribute-block-reward /ishikari-distribute-block-reward
ENTRYPOINT ["./ishikari-distribute-block-reward"]
