FROM golang:1.25rc2-alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/api/main.go

FROM alpine:3.20

COPY --from=builder /app/server /server
COPY .env ./.env

EXPOSE 3333

ENTRYPOINT ["/server"]