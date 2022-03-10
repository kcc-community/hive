# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /issue-9
WORKDIR /issue-9
RUN go build .

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /issue-9/issue-9 /issue-9
ENTRYPOINT ["./issue-9"]
