FROM golang:1.21-alpine AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build .

FROM golang:1.21-alpine

# curl added only for health check if anyone needs it
RUN apk update && apk add curl
WORKDIR /app
COPY --from=builder /build/main .

ENTRYPOINT ["/app/main"]