# BINGO

本项目是zrpc跨语言的demo实验，希望结合go-zero脚手架的优异特性和python在深度学习应用中的丰富社区。我们在这个项目中使用api调用zrpc-client，然后使用grpc实现跨语言的调用，python实现的grpc客户端可以实现模型训练和推理。由于这可能将是我们后续的只是开始或者基石，所以我命名其为bingo,寓意begin + go = bingo.  

<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />

<p align="center">
  <a href="https://github.com/kongfuguagua/BINGO/blob/main/BINGO.png">
    <img src="BINGO.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">BINGO</h3>
  <p align="center">
    RAG架构
    <br />
    <a href="https://github.com/kongfuguagua/BINGO"><strong>探索本项目的文档 »</strong></a>
    <br />
    <br />
    <a href="https://github.com/kongfuguagua/BINGO">查看Demo</a>
    ·
    <a href="https://github.com/kongfuguagua/BINGO/issues">报告Bug</a>
    ·
    <a href="https://github.com/kongfuguagua/BINGO/issues">提出新特性</a>
  </p>

</p>


 本篇README.md面向开发者
 
## 目录

- [上手指南](#上手指南)
  - [开发前的配置要求](#开发前的配置要求)
  - [安装步骤](#安装步骤)
- [文件目录说明](#文件目录说明)
- [Demo](#Demo)
- [使用到的框架](#使用到的框架)
- [版本控制](#版本控制)
- [作者](#作者)
- [鸣谢](#鸣谢)

### 上手指南

本项目架构是 users->api->rpc Client(go)->rpc Server(python)->model
<p align="center">
  <a href="https://github.com/kongfuguagua/BINGO/">
    <img src="device.png">
  </a>
  </p>
</p>

快速开始！


1启动etcd作为注册表
```
docker run  -it --rm -p 2379:2379   -p 2380:2380   --name etcd-gcr-v3.5.16   gcr.io/etcd-development/etcd:v3.5.16   /usr/local/bin/etcd   --name s1   --data-dir /etcd-data   --listen-client-urls http://0.0.0.0:2379   --advertise-client-urls http://0.0.0.0:2379   --listen-peer-urls http://0.0.0.0:2380   --initial-advertise-peer-urls http://0.0.0.0:2380   --initial-cluster s1=http://0.0.0.0:2380   --initial-cluster-token tkn   --initial-cluster-state new   --log-level info   --logger zap   --log-outputs stderr
```
2启动rpc服务器，程序会自动注册，并启动监听
```
cd rpc/pyserver
python server.py
```
3启动api网关
```
cd api
go run dl.go -f etc/dl-api.yaml
```
4测试demo
```
curl --location --request POST 'http://localhost:8888/dl/api/create' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--header 'Accept: */*' \
--header 'Host: localhost:8888' \
--header 'Connection: keep-alive' \
--data-raw '{
    "namespace": "test",
    "dlname": "mnist",
    "spec": {
        "dlmodel": {
            "dlmodelname": "my-cnn",
            "dlmodelpath": "/sbin",
            "dlmodelstatus": true,
            "dlinput": "3*24*24",
            "dloutput": "4*24"
        },
        "dldataobj": {
            "database": "mnist",
            "name": "mnist",
            "status": false,
            "total": 23,
            "type": "pic"
        }
    }
}'
```
预期结果
```
{
    "dlinfo": {
        "metadata": {
            "id": "test",
            "namespace": "test",
            "dlname": "test"
        },
        "spec": {
            "dlmodel": {
                "dlmodelname": "my-cnn",
                "dlmodelpath": "",
                "dlmodelstatus": false,
                "dlinput": "",
                "dloutput": ""
            },
            "dldataobj": {}
        }
    }
}
```
###### 开发前的配置要求

1. python3.12
2. 参考requirements
3. golang1.22
4. 参考mod.go

###### **安装步骤**

1. Clone the repo

```sh
git clone https://github.com/kongfuguagua/BINGO.git
```

2. Build the images (pass)

```sh
cd ApplicationLibrary/xxx
docker build -t --platform=linux/arm64 image_name:image_tag .
```

3. RUN the application

```sh
cd xxx
go run main.go -f etc/xxx.yaml
```

### 文件目录说明
eg:

```
BINGO 
--api 是项目对外暴露的api和逻辑实现
--db  是项目数据库内容，coming！
--rpc 是项目的rpc定义和实现
```



### Demo 

一个深度学习应用调用流程如下：

<p align="center">
  <a href="https://github.com/kongfuguagua/BINGO/blob/main/energy.png">
    <img src="energy.png">
  </a>
  <a href="https://github.com/kongfuguagua/BINGO/blob/main/diagram.png">
    <img src="diagram.png">
  </a>
  </p>
</p>

涉及模块均可以在xxx找到，主要包括自编码器、微调大模型数据集等



### 使用到的框架

- [Domain-driven design](https://en.wikipedia.org/wiki/Domain-driven_design)
- [go-zero](https://go-zero.dev/docs/concepts/overview)
- [grpc-python](https://grpc.org.cn/docs/languages/python/basics/)
coming!

#### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request


### 作者

a25505597703@gmail.com  


### 版权说明

该项目签署了MIT 授权许可，详情请参阅 [LICENSE.txt](https://github.com/kongfuguagua/BINGO/blob/master/LICENSE.txt)

### 鸣谢


- [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
- [Img Shields](https://shields.io)
- [Choose an Open Source License](https://choosealicense.com)
- [GitHub Pages](https://pages.github.com)
- [Animate.css](https://daneden.github.io/animate.css)

<!-- links -->
[your-project-path]:kongfuguagua/BINGO
[contributors-shield]: https://img.shields.io/github/contributors/kongfuguagua/BINGO?style=flat-square
[contributors-url]: https://github.com/kongfuguagua/BINGO/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/kongfuguagua/BINGO?style=flat-square
[forks-url]: https://github.com/kongfuguagua/BINGO/network/members
[stars-shield]: https://img.shields.io/github/stars/kongfuguagua/BINGO?style=flat-square
[stars-url]: https://github.com/kongfuguagua/BINGO/stargazers
[issues-shield]: https://img.shields.io/github/issues/kongfuguagua/BINGO?style=flat-square
[issues-url]: https://img.shields.io/github/issues/kongfuguagua/BINGO
[license-shield]: https://img.shields.io/github/license/kongfuguagua/BINGO?style=flat-square
[license-url]: https://github.com/kongfuguagua/BINGO/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/kongfuguagua