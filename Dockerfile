FROM golang:1.23.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o sx-evo-debug cmd/main.go

FROM alpine:3.21
RUN apk update && apk add --no-cache \
    bash \
    can-utils \
    eudev \
    linux-headers \
    net-tools \
    sudo

WORKDIR /app

COPY --from=builder /app/sx-evo-debug .

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

CMD ["./entrypoint.sh"]
