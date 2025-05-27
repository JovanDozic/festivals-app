FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/app .

RUN mkdir -p ./cmd/api/assets
COPY --from=builder /app/cmd/api/assets ./assets/

EXPOSE 4000

CMD ["./app"]