FROM golang:1.25rc2-alpine3.22 AS builder

WORKDIR /src/portfolio

COPY go.mod go.sum ./
COPY . .
RUN go build -o portfolio ./cmd/api

FROM alpine:3.20 as binary

COPY --from=builder /src/portfolio/portfolio .

EXPOSE 3303
CMD ["./portfolio"]
