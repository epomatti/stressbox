FROM golang:alpine AS builder

# Required for AWS services certificates
RUN apk update && apk add ca-certificates

WORKDIR /build
COPY . .
RUN go mod download
RUN go build .

FROM golang:alpine
RUN apk update && apk add curl
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
WORKDIR /app
COPY --from=builder /build/main .
ENTRYPOINT ["/app/main"]