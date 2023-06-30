FROM golang:alpine AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build .

FROM golang:alpine

# curl added only for health check
RUN apk update && apk add curl
WORKDIR /app
COPY --from=builder /build/main .

ENTRYPOINT ["/app/main"]