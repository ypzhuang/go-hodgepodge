FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git tzdata

RUN adduser -D -g '' appuser

ADD . /go/src/github.com/ypzhuang/goiothub
WORKDIR /go/src/github.com/ypzhuang/goiothub
RUN go get -d -v  github.com/eclipse/paho.mqtt.golang

RUN go build -o /go/bin/goiothub
RUN ls -al /go/bin

FROM alpine
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/bin/goiothub /go/bin/goiothub

USER appuser

ENV TZ=Asia/Shanghai
EXPOSE 8000
ENTRYPOINT ["/go/bin/goiothub"]


