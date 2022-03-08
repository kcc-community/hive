# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /v2-distribute-block-reward
WORKDIR /v2-distribute-block-reward
RUN go build .

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /v2-distribute-block-reward/v2-distribute-block-reward /v2-distribute-block-reward
ENTRYPOINT ["./v2-distribute-block-reward"]
