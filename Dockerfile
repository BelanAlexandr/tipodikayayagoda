FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/server/main.go


FROM alpine:latest

WORKDIR /app


COPY --from=builder /app/app .
COPY --from=builder /app/internal/templates ./internal/templates
COPY --from=builder /app/migrations ./migrations


RUN mkdir -p static/uploads

EXPOSE 8080

CMD ["./app"]