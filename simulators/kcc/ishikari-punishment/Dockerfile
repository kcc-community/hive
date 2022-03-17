# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /ishikari-punishment
WORKDIR /ishikari-punishment
RUN go build .

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /ishikari-punishment/ishikari-punishment /ishikari-punishment
ENTRYPOINT ["./ishikari-punishment"]
