FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ai-balancer ./cmd/api


FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/ai-balancer .
COPY configs ./configs

EXPOSE 9080

CMD ["./ai-balancer"]