# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /new-txpool
WORKDIR /new-txpool
RUN go build .

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /new-txpool/new-txpool /new-txpool
ENTRYPOINT ["./new-txpool"]
