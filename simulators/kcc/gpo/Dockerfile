# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /gpo
WORKDIR /gpo
RUN go build .

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /gpo/gpo /gpo
ENTRYPOINT ["./gpo"]
