# Java 示例项目

这是一个基本的Java项目，使用Maven作为构建工具。

## 项目结构

```
java-project/
├── src/
│   ├── main/java/com/example/app/    # 主要源代码
│   └── test/java/com/example/app/    # 测试代码
└── pom.xml                           # Maven配置文件
```

## 环境要求

- JDK 11或更高版本
- Maven 3.6或更高版本

## 构建和运行

### 编译项目

```bash
cd java-project
mvn compile
```

### 运行测试

```bash
mvn test
```

### 打包项目

```bash
mvn package
```

### 运行应用程序

```bash
java -jar target/app-1.0-SNAPSHOT.jar
```

或者使用Maven：

```bash
mvn exec:java -Dexec.mainClass="com.example.app.App"
```

## 功能

这个示例项目包含一个简单的Java应用程序，它打印"你好，世界！"并提供一个简单的加法方法。 