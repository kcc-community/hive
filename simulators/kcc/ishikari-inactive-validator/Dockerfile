# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /ishikari-inactive-validator
WORKDIR /ishikari-inactive-validator
RUN go build .

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /ishikari-inactive-validator/ishikari-inactive-validator /ishikari-inactive-validator
ENTRYPOINT ["./ishikari-inactive-validator"]
