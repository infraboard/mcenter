FROM registry.cn-hangzhou.aliyuncs.com/godev/golang:1.20 AS builder

LABEL stage=gobuilder

COPY . /src

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GOPROXY https://goproxy.cn,direct
# ENV GOPRIVATE="*.gitlab.com"

WORKDIR /src
RUN make build

FROM alpine
WORKDIR /app
EXPOSE 8010
COPY --from=builder /src/dist/mcenter-api /app/mcenter-api
COPY --from=builder /src/etc /app/etc

CMD ["./mcenter-api", "start", "-t", "env"]