# 使用 golang alpine 作为构建镜像
FROM golang:1.24.0-alpine3.21 AS builder
# 设置工作目录为 /app
WORKDIR /app
# 安装编译所需的依赖
RUN apk add --no-cache gcc musl-dev build-base
# 将当前目录的所有文件复制到容器的工作目录
COPY . .
# 启用 Go modules
RUN go env -w GO111MODULE=on
# 编译 Go 程序，启用 CGO
RUN CGO_ENABLED=1 go build -o /usr/bin/trans main.go
# 使用 alpine 作为运行镜像
FROM alpine:3.21
# 安装必要的运行时依赖：sqlite、sqlite-dev、translate-shell 和 基础运行时库
RUN apk add --no-cache sqlite sqlite-dev translate-shell libc6-compat build-base
# 从构建阶段复制编译好的程序到运行镜像
COPY --from=builder /usr/bin/trans /usr/bin/trans
# 设置容器启动命令
ENTRYPOINT ["/usr/bin/trans"]
