FROM golang:1.23-alpine as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN go build -o seed-manager ./cmd/service

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/seed-manager .

EXPOSE 8000

CMD ["./seed-manager"]

