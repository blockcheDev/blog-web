FROM golang:1.24 AS builder
WORKDIR /usr/local/server
COPY ./server/ ./
# CGO_ENABLED=0 静态链接
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o server main.go


FROM alpine:3.21
WORKDIR /usr/local
COPY --from=builder /usr/local/server/server ./
COPY --from=builder /usr/local/server/config/config.yaml ./config/

EXPOSE 8080

ENTRYPOINT ./server