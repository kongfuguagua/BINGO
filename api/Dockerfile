FROM registry.cn-hangzhou.aliyuncs.com/eryajf/golang:1.23.4-alpine3.21 AS builder
# 添加标签，标识这是构建阶段
LABEL stage=gobuilder

# 禁用CGO，确保生成的二进制文件是静态链接的
ENV CGO_ENABLED=0
# 设置Go模块代理，使用国内镜像加速下载
ENV GOPROXY=https://goproxy.cn,direct
# 将Alpine的软件源替换为阿里云镜像，提高下载速度
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 更新包索引并安装时区数据包
RUN apk update --no-cache && apk add --no-cache tzdata

# 设置工作目录为/build
WORKDIR /build

# 复制Go模块文件到容器中
ADD go.mod .
ADD go.sum .
# 下载Go模块依赖
RUN go mod download
# 复制所有源代码到容器中
COPY . .
# 编译Go程序，-ldflags="-s -w"用于减小二进制文件大小
# -s: 去除符号表和调试信息
# -w: 去除DWARF调试信息
RUN go build -ldflags="-s -w" -o /app/dl api/dl.go


# 使用轻量级的busybox镜像作为最终运行环境
FROM swr.cn-north-4.myhuaweicloud.com/ddn-k8s/docker.io/library/busybox:latest

# 设置工作目录为/app
WORKDIR /app
# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/dl /app/dl
# 从构建阶段复制配置文件目录
COPY --from=builder /build/api/etc /app/etc

# 设置容器启动时执行的命令
# 运行dl程序，并指定配置文件路径
CMD ["./dl", "-f", "etc/dl-api.yaml"]


