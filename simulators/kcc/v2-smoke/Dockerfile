# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /kcc-v2-smoke
WORKDIR /kcc-v2-smoke
RUN go build .

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /kcc-v2-smoke/v2-smoke /kcc-v2-smoke
ENTRYPOINT ["./kcc-v2-smoke"]
