# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /app
WORKDIR /app
RUN go build -v -o app 

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /app/app /app
ENTRYPOINT ["./app"]
