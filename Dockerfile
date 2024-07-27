FROM golang:1.21-alpine AS builder

WORKDIR /build
RUN apk add --no-cache make

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN make build

FROM alpine:3.18
RUN apk add --no-cache ca-certificates

LABEL maintainer="Akhand Patel  (https://akhand.me)"
LABEL description="API for SRE Bootcamp(https://one2n.io/sre-bootcamp)"

COPY --from=builder /build/bin/nerdstore /nerdstore
COPY --from=builder /build/*.env /

# Expose port 8080
EXPOSE 8080

# Command to run when starting the container.
ENTRYPOINT ["/nerdstore"]