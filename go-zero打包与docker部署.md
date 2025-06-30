# Go-zero镜像打包与Docker部署

## 镜像打包

将服务部署到服务器上，需要将项目服务打包成镜像，这就需要一个dockerfile文件来执行docker build

Dockerfile举例（根据goctl docker指令生成并添加配置目录，不一定稳定，需要进行测试，根据你的需求修改）：

```dockerfile
# 第一阶段：构建阶段
# 使用官方Go镜像作为构建环境，alpine版本体积更小
FROM golang:alpine AS builder

# 添加标签，标识这是构建阶段
LABEL stage=gobuilder

# 禁用CGO，确保生成的二进制文件是静态链接的
ENV CGO_ENABLED 0
# 设置Go模块代理，使用国内镜像加速下载
ENV GOPROXY https://goproxy.cn,direct
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


# 第二阶段：运行阶段
# 使用轻量级的busybox镜像作为最终运行环境
FROM busybox:latest

# 从构建阶段复制SSL证书，用于HTTPS连接
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# 从构建阶段复制时区信息
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
# 设置时区环境变量为上海时区
ENV TZ Asia/Shanghai

# 设置工作目录为/app
WORKDIR /app
# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/dl /app/dl
# 从构建阶段复制配置文件目录
COPY --from=builder /build/api/etc /app/etc

# 设置容器启动时执行的命令
# 运行dl程序，并指定配置文件路径
CMD ["./dl", "-f", "etc/dl-api.yaml"]

```

在Dockerfile所在目录执行docker build -t <image_name> .

可以在windows或者linux执行，只要安装了docker即可，build构建完成后该镜像存储在本地，比如如果是windows本机进行的构建，需要将镜像推送到云端的镜像仓库，才能在服务器上进行镜像的拉取，所以推荐直接在服务器上构建

构建完成后，得到了镜像，可以使用docker images查询是否构建成功，接下来可以使用docker进行服务的部署

## Docker服务部署

该步骤要在服务器上执行，确保服务器的网络环境能够拉取需要使用镜像

可以通过docker-compose的方式进行部署，根据你的bingo项目，好像加入了redis缓存？，选择是否要部署redis服务，由于bingo还需要etcd服务进行服务注册，所以该项目需要部署前面打包的api服务，redis服务，etcd服务

docker-compose.yml举例：根据add_db分支的bingo项目生成，需要你根据实际的服务需求进行修改

```yaml
# 指定Docker Compose文件格式版本
version: '3'

# 定义服务列表
services:
  # Redis缓存服务
  redis:
    image: redis:7-alpine #redis镜像，根据你使用的修改
    # 设置容器名称为bingo-redis
    container_name: bingo-redis
    # 端口映射配置
    ports:
      # 将主机的6379端口映射到容器的6379端口
      - "6379:6379"
    # 网络配置
    networks:
      # 加入自定义网络bingo-network
      - bingo-network

  # etcd服务发现和配置管理服务
  etcd:
    image: quay.io/coreos/etcd:v3.5.9 #etcd镜像，使用CoreOS官方etcd v3.5.9镜像，根据需要修改
    # 设置容器名称为bingo-etcd
    container_name: bingo-etcd
    # 环境变量配置
    environment:
      # 设置etcd节点名称
      - ETCD_NAME=etcd0
      # 设置etcd数据存储目录
      - ETCD_DATA_DIR=/etcd-data
      # 设置客户端监听地址，监听所有网络接口的2379端口
      - ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379
      # 设置客户端访问地址，其他服务通过etcd:2379访问
      - ETCD_ADVERTISE_CLIENT_URLS=http://etcd:2379
      # 设置节点间通信监听地址
      - ETCD_LISTEN_PEER_URLS=http://0.0.0.0:2380
      # 设置节点间通信广播地址
      - ETCD_INITIAL_ADVERTISE_PEER_URLS=http://etcd:2380
      # 设置初始集群配置
      - ETCD_INITIAL_CLUSTER=etcd0=http://etcd:2380
      # 设置集群令牌，用于区分不同集群
      - ETCD_INITIAL_CLUSTER_TOKEN=etcd-cluster
      # 设置集群初始状态为新建
      - ETCD_INITIAL_CLUSTER_STATE=new
    # 端口映射配置
    ports:
      # 映射客户端访问端口
      - "2379:2379"
      # 映射节点间通信端口
      - "2380:2380"
    # 数据卷挂载配置
    volumes:
      # 将etcd_data卷挂载到容器的/etcd-data目录，实现数据持久化
      - etcd_data:/etcd-data
    # 网络配置
    networks:
      # 加入自定义网络bingo-network
      - bingo-network

  # API服务
  dl-api:
    image: dl-api:latest #你构建的api镜像，使用本地构建的dl-api镜像，根据你的镜像名称和标签修改
    # 设置容器名称为bingo-dl-api
    container_name: bingo-dl-api
    # 端口映射配置
    ports:
      # 将主机的8888端口映射到容器的8888端口
      - "8888:8888"
    # 服务依赖配置，确保redis和etcd先启动
    depends_on:
      # 依赖redis服务
      - redis
      # 依赖etcd服务
      - etcd
    # 环境变量配置
    environment:
      # 设置时区为上海时区
      - TZ=Asia/Shanghai
    # 网络配置
    networks:
      # 加入自定义网络bingo-network
      - bingo-network
    # 重启策略：除非手动停止，否则总是重启
    restart: unless-stopped

# 数据卷定义
volumes:
  # etcd数据卷，用于持久化etcd数据
  etcd_data:

# 网络定义
networks:
  # 自定义网络bingo-network
  bingo-network:
    # 使用bridge驱动，创建桥接网络
    driver: bridge
```

使用指令：docker-compose up <-d> 进行服务的部署

结束时使用docker-compose down 停止启动的容器、网络等，重置运行的环境



部署之后，**访问服务器的IP:服务对应的端口号**   即可使用该服务对外暴露的接口



除了docker-compose之外，还可以用docker直接部署服务，比如bingo项目中的etcd服务：

docker run  -it --rm -p 2379:2379   -p 2380:2380   --name etcd-gcr-v3.5.16   gcr.io/etcd-development/etcd:v3.5.16   /usr/local/bin/etcd   --name s1   --data-dir /etcd-data   --listen-client-urls http://0.0.0.0:2379   --advertise-client-urls http://0.0.0.0:2379   --listen-peer-urls http://0.0.0.0:2380   --initial-advertise-peer-urls http://0.0.0.0:2380   --initial-cluster s1=http://0.0.0.0:2380   --initial-cluster-token tkn   --initial-cluster-state new   --log-level info   --logger zap   --log-outputs stderr

对于其他服务也可以使用该方式进行部署