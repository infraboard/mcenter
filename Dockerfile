FROM golang:1.18 as builder
COPY ./ /code
ENV GOPROXY="https://goproxy.cn,direct"

WORKDIR /code
RUN make build

FROM busybox
WORKDIR /go/bin
EXPOSE 80
COPY --from=builder /code/dist/mcenter /go/bin/
CMD ["/go/bin/mcenter"]