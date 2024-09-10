FROM golang:1.22.5-alpine3.20 as build-dev

WORKDIR /go/src/app

COPY --link . .

RUN apk add --no-cache upx tzdata || \
        go env -w GO111MODULE=on && \
        go env -w GOPROXY=https://goproxy.cn,direct && \
        go mod tidy && \
        go build -ldflags="-s -w" -o dockerproxy . && \
        [ -e /usr/bin/upx ] && upx dockerproxy || echo

FROM alpine

WORKDIR /src

COPY --link --from=build-dev /go/src/app/dockerproxy ./dockerproxy
COPY --from=build-dev /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-dev /usr/share/zoneinfo /usr/share/zoneinfo
COPY .env.example .env
COPY templates ./templates
COPY static ./static

ENV TZ=Asia/Shanghai
EXPOSE 9101

CMD ["./dockerproxy"]