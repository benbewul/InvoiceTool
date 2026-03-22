FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY cmd ./cmd
COPY web ./web

RUN go mod tidy
RUN go build -o app ./cmd/server

FROM alpine:3.20
WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/web ./web

EXPOSE 8080
CMD ["./app"]
