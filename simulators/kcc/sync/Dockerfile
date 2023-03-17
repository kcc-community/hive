# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /sync
WORKDIR /sync
RUN go build .

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /sync/sync /sync
ENTRYPOINT ["./sync"]
