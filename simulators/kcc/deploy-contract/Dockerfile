# Build the simulator.
FROM golang:1-alpine AS builder
RUN apk --no-cache add gcc musl-dev linux-headers
ADD . /deploy-contract
WORKDIR /deploy-contract
RUN go build .

# Build the runner container.
FROM alpine:latest
ADD . /
COPY --from=builder /deploy-contract/deploy-contract /deploy-contract
ENTRYPOINT ["./deploy-contract"]
