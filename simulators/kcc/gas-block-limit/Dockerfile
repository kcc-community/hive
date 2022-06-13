# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /gas-block-limit
WORKDIR /gas-block-limit
RUN go build .

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /gas-block-limit/gas-block-limit /gas-block-limit
ENTRYPOINT ["./gas-block-limit"]
