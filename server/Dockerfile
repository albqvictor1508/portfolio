FROM golang:1.25rc2-alpine3.22 AS builder

WORKDIR /src/portfolio

COPY go.mod go.sum ./
COPY . .

RUN go build -o portfolio ./cmd/api

FROM alpine:3.20

COPY --from=builder /app/server /app/server
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /go/bin/goose /app/goose

EXPOSE 3303

CMD ["./portfolio"]
