# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /smoke-v2-hardfork
WORKDIR /smoke-v2-hardfork
RUN go build .

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /smoke-v2-hardfork/smoke-v2-hardfork /smoke-v2-hardfork
ENTRYPOINT ["./smoke-v2-hardfork"]
