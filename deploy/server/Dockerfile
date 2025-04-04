FROM golang:1.24 AS builder
WORKDIR /app/server
COPY ./server/ ./
# CGO_ENABLED=0 静态链接
RUN go env -w  GOPROXY=https://goproxy.cn,direct \
    && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go


FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/server/server ./
COPY --from=builder /app/server/config/config.yaml ./config/

# 添加时区数据库，设置了TZ后会来这里查询时区信息
RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#https://mirrors.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories \
    && apk update && apk add tzdata

EXPOSE 8080

ENTRYPOINT ./server