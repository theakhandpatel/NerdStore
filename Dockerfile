FROM golang:1.20-alpine AS builder

WORKDIR /build
RUN apk add --no-cache make

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download
RUN go get -d -v ./...
RUN go install -v ./...

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN make build

FROM scratch

LABEL maintainer="Akhand Patel  (https://akhand.me)"
LABEL description="API for SRE Bootcamp(https://one2n.io/sre-bootcamp)"


COPY --from=builder ["/build/*.env", "/"]

# Command to run when starting the container.
ENTRYPOINT ["/nerdstore"]