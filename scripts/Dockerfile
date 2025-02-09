# Dockerfile
FROM golang:1.22-bullseye AS builder
WORKDIR /app
COPY . .
RUN make build
FROM ubuntu:22.04
RUN apt install libc6 fuse
COPY --from=builder /app/bin/system-usage-collector /system-usage-collector #修改对应的文件
CMD ["/system-usage-collector"] #修改对应的文件