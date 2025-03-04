# Java gRPC 服务器

这是一个基于Java的gRPC服务器实现，用于处理深度学习任务的管理。该服务器实现了`dl.proto`中定义的服务接口，并支持与etcd进行服务注册。

## 项目结构

```
javaserver/
├── src/
│   ├── main/
│   │   ├── java/
│   │   │   └── com/example/app/      # 主要源代码
│   │   │       ├── App.java          # 服务器启动类
│   │   │       └── DLFunctionImpl.java # gRPC服务实现
│   │   └── proto/                    # Proto文件目录
│   │       └── dl.proto             # 服务定义文件
└── pom.xml                          # Maven配置文件
```

## 环境要求

- JDK 11或更高版本
- Maven 3.6或更高版本
- etcd服务器（运行在`http://127.0.0.1:2379`）

## 功能特性

- 实现了dl.proto中定义的所有RPC方法：
  - getDLById: 根据ID获取深度学习任务信息
  - createDL: 创建新的深度学习任务
  - initDL: 初始化深度学习任务
- 支持服务注册到etcd
- 自动续约etcd租约
- 优雅关闭处理

## 运行说明

### 1. 安装依赖

确保你的系统已经安装了以下软件：
- JDK 11+
- Maven 3.6+
- etcd服务器

### 2. 启动etcd服务器

在运行Java服务器之前，请确保etcd服务器已经启动：

```bash
# Windows
etcd.exe

# Linux/Mac
etcd
```

### 3. 编译项目

```bash
# 进入项目目录
cd rpc/javaserver

# 清理并编译
mvn clean compile
```

### 4. 打包项目

```bash
# 打包成可执行jar
mvn package
```

### 5. 运行服务器

```bash
# 运行服务器
java -jar target/app-1.0-SNAPSHOT-jar-with-dependencies.jar
```

服务器启动后会：
1. 启动gRPC服务器并监听50051端口
2. 向etcd注册服务（键为"dl.rpc/10"）
3. 自动维护etcd租约（每5分钟续约一次）
4. 在关闭时自动清理资源

## 验证服务是否正常运行

### 1. 检查gRPC服务器

服务器启动后，你应该能看到以下日志：
```
gRPC服务器启动，监听端口: 50051
已注册PickFirst负载均衡器提供者
服务已注册到etcd: dl.rpc/10 -> 127.0.0.1:50051
```

