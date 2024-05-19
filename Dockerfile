FROM registry.cn-hangzhou.aliyuncs.com/godev/golang:1.22 AS builder

LABEL stage=gobuilder

WORKDIR /src
COPY go.mod .
COPY go.sum .

ENV CGO_ENABLED=0 \
GOOS=linux \
GOARCH=amd64 \
GOPROXY=https://goproxy.cn,direct

# 下载依赖
RUN go mod download
# 执行构建
COPY . .
RUN make build

FROM registry.cn-hangzhou.aliyuncs.com/godev/alpine:latest
WORKDIR /app
EXPOSE 8010
COPY --from=builder /src/dist/mcenter-api /app/mcenter-api
COPY --from=builder /src/etc /app/etc

# 默认配置, HTTPS
ENV APP_NAME=mcenter \
APP_DOMAIN=console.mdev.group \
APP_SECURITY=true \
HTTP_HOST=127.0.0.1 \
HTTP_PORT=8010 \
GRPC_HOST=127.0.0.1 \
GRPC_PORT=18010 \
MONGO_ENDPOINTS=127.0.0.1:27017 \
MONGO_DATABASE=mcenter


CMD ["./mcenter-api", "start", "-t", "env"]