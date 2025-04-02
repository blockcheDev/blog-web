FROM golang:1.24 AS builder
WORKDIR /app/server
COPY ./server/ ./
# CGO_ENABLED=0 静态链接
RUN go env -w  GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && CGO_ENABLED=0 GOOS=linux go build -o server main.go


FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/server/server ./
COPY --from=builder /app/server/config/config.yaml ./config/

EXPOSE 8080

ENTRYPOINT ./server