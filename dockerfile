FROM golang:1.24.0-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on
RUN go build -o /usr/bin/trans main.go
FROM alpine:3.21
RUN apk add --no-cache sqlite sqlite-dev translate-shell
COPY --from=builder /usr/bin/trans /usr/bin/trans
ENTRYPOINT ["/usr/bin/trans"]
