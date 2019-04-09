# Docker


| docker命令 | 意义 | 
| ----- | ---- |
| docker version                    |        查看docker版本  |
| docker search xxx                 |        从仓库中查找镜像  |
| docker pull xxx/xxx(用户名/镜像名)  |     从仓库中下载镜像  |
| docker run xxx/xxx(用户名/镜像名)   |    启动容器  |
| docker ps （-l , -a...）             |       查看所有容器  |
| docker inspect 容器id              |       查看具体容器  |
| docker commit  容器id              |       从容器创建一个新的镜像  |
| docker push xxx/xxx(用户名/镜像名)  |     提交镜像  |


## 容器技术 vs 虚拟机
- 虚拟机通过中间层将一台或者多台独立的及其运行于物理硬件当中；而容器则是直接运行在操作系统内核之上的用户空间 

- 容器技术的磁盘占有空间少，虚拟机部署应用包含应用和其依赖的库，还需要包含完整的操作系统，容器只需前两者。另外，虚拟机需要模拟硬件的行为，对cpu，内存消耗较大

## docker 简介

- docker 基于linux内核，是一种容器技术（容器虚拟化的方案）
- docker 是一个可以把开发的应用程序自动的部署到容器中的一个开源引擎（Golang）
- docker 在虚拟化的容器执行环境中，增加了一个应用程序部署引擎

## docker组成
- Docker Client  客户端
- Docker Daemon 守护进程
- Docker Image  镜像
- Docker Container 容器
- Docker Registry 仓库

### Docker Client / Docker Daemon （c/s架构）
- 终端中的一些docker命令（eg : docker pull, docker run ...）通过client传给daemon（支持远程），然后返回结果
- 通过客户端访问守护进程，进而操作容器

### Docker Image 
- 容器的基石，容器基于镜像启动和运行
- 一个镜像可以作为另外镜像的底层镜像（父镜像）

### Docker Container 
- 通过镜像启动，是docker的执行单元，一个container可以运行一个/多个用户进程

### Docker Registry 
- 用户保存用户的镜像，分为public和private
- docker 官方提供公开的仓库 Docker Hub 

---------------------------------------------------------------

## Docker 容器的基本操作

### 启动
- docker run IMAGE [COMMAND] [ARG...] 
    - 在新的容器中执行命令
        - 启动交互式容器 : docker run -i -t IMAGE /bin/bash
            - eg : docker run -i -t ubuntu /bin/bash
                - -i : --interactive = true (打开标准输入)
                - -t : --tty = true (打开tty终端)

        - 自定义容器名称 : docker run --name=自定义名字 -i -t IMAGE /bin/bash

- docker start [-i] 容器名
    - 重新启动容器(不产生新的容器)

- docker rm 容器名
    - 删除已经停止的容器
