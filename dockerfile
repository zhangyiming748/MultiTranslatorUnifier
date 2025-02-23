FROM golang:1.23.6-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on
#RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o /usr/bin/trans main.go
# RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#http://mirrors4.tuna.tsinghua.edu.cn/alpine#g' /etc/apk/repositories
RUN apk add translate-shell
CMD ["/usr/bin/trans"]
#docker run -it --rm --name alpine -v C:\Users\zen\Github\MultiTranslatorUnifier:/data golang:1.23.6-alpine3.21 ash