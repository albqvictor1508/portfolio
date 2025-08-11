FROM golang:1.25rc2-alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/api/main.go

FROM alpine:3.20

RUN apk add --no-cache netcat-openbsd

COPY --from=builder /app/server /app/server
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /go/bin/goose /app/goose
COPY start.sh /app/start.sh

RUN chmod +x /app/start.sh

EXPOSE 3333

ENTRYPOINT ["/app/start.sh"]
