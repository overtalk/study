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
| docker port 容器名                 |     查看容器的端口映射情况  |


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

### 交互式容器
- docker run IMAGE [COMMAND] [ARG...] 
    - 在新的容器中执行命令
        - docker run -i -t IMAGE /bin/bash
            - eg : docker run -i -t ubuntu /bin/bash
                - -i : --interactive = true (打开标准输入)
                - -t : --tty = true (打开tty终端)

        - 自定义容器名称 : docker run --name=自定义名字 -i -t IMAGE /bin/bash

- docker start [-i] 容器名
    - 重新启动容器(不产生新的容器)

- docker rm 容器名
    - 删除已经停止的容器


### 守护式容器
- docker run -d IMAGE [COMMAND] [ARG...] 
    - -d : 后台方式运行

- docker attach 容器名
    - 附加到运行中的容器 

- docker logs [-f][-t][--tail] 容器名
    - -f : --follows=true   : 追踪日志 
    - -t : -- timestamps=true   : 日志加上时间戳
    - --tail="all"      : tail 的意思是返回结尾出多少数量的日志，如果不指定，返回所有

- docker top 容器名
    - 查看运行中容器内所有的进程

 - docker exec -d [-d][-i][-t] 容器名 [COMMAND] [ARG...] 
    - 在容器内启动新的进程
   
- docker stop 容器名
    - 向容器发送信号，等待容器停止
- docker kill 容器名
    - 直接停止容器


### 容器中的部署
- 容器的端口映射
    - run [-P][-p] (大小写)
        - -P : --publish-all=true 为容器所有暴露的端口进行映射
            - eg : docker run -P -i -t ubuntu /bin/bash
        - -p : --publish=[] 指定映射容器的哪些端口, 有多种格式
            - 只指定容器的端口 （containerPort）
                - 宿主机的端口随机映射
                - eg : docker run -p 80 -i -t ubuntu /bin/bash
            - 同时指定宿主机和容器端口 （hostPort : containerPort）
                - 一一对应的
                - eg : docker run -p 80 -i -t ubuntu /bin/bash
            - ip对应容器端口 （ip : containerPort）
                - eg : docker run -p 0.0.0.0:80 -i -t ubuntu /bin/bash
            - ip+端口对应容器端口 （ip+port : containerPort）
                - eg : docker run -p 0.0.0.0:8080:80 -i -t ubuntu /bin/bash

### 对镜像的操作
- 镜像的两个重要属性
    - repository 镜像的仓库, eg : ubuntu, centos ...
    - tag 标签

- 列出镜像
    - docker images [OPTSIONS] [REPOSITORY]
        - -a 查看所有的镜像
        - -f, --filter=[] 添加过滤条件
        - -q 只查看镜像的唯一id

- 查看镜像的具体信息
    - docker inspect [OPTIONS] CONTAINER/IMAGE (这个命令既支持容器的查看，也支持镜像的查看)
        - -f : --format=""

- 删除镜像
    - docker rmi [OPTIONS] IMAGE
        - -f 强制删除
        - 删除所有镜像 : docker rmi $(docekr images -q)

- 查找镜像
    - docker search [OPTIONS] TERM

- 拉取镜像
    - docker pull [OPTIONS] NAME[:TAG]
        - -a 下载所有匹配NAME的有tag的
        
- 推送镜像
    - docker push NAME[:TAG]


---------------------------------------------------------------

## Dockerflie
- FROM
    - FROM [image]:[tag]
    - dockerfile第一条非注释的指令

- MAINTAINER
    - 作者信息，一般包含所有者以及联系方式

- RUN
    - 镜像构建过程中的指令，包含当前镜像中运行的命令，包含两种模式（shell & exec）
    - RUN [command] (shell 模式)
        - 以 /bin/bash -c command 执行命令
            - eg : RUN echo hello

    - RUN ["executable", "param1", "param2"] (exec 模式)
    - 可以指定其他形式的shell来运行指令
        - eg :  RUN ["/bin/bash", "-c", "echo hello"]

- EXPOSE
    - 指定暴露的容器的端口

- CMD
    - 容器运行的默认命令，与前面的RUN不同的是 : 后者是构建镜像的过程中执行的命令
    - 如果在使用 `docker run` 命令运行容器的时候加上了容器运行时候的指令，则cmd中的指令会被覆盖，不会执行。既：cmd用于指定容器的默认行为
    - CMD指令的两种模式（shell & exec & 特殊模式）
    - CMD command param1 param2 (shell 模式)
        - 以 /bin/bash -c command 执行命令
            - eg : CMD echo hello
    - CMD ["executable", "param1", "param2"] (exec 模式)
        - 可以指定其他形式的shell来运行指令
        - eg :  CMD ["/bin/bash", "-c", "echo hello"]
    - CMD ["param1", "param2"] (只有一些参数)
        - 通常与`ENTERYPOINT`指令配合使用，作为其默认参数

- ENTERYPOINT
    - 与CMD指令类似，也有两种模式（shell & exec），但是它不会被 `docker run` 指令中的指令覆盖
    - 如果需要覆盖，需要在 `docker run` 指令中加上`--enterypoint`选项

- ADD & COPY
    - 将文件/目录复制到使用dockerfile构建的镜像中
    - 用法 : ADD[COPY] [src]...[dest] 
        - src : 可以是本地地址（必须是构建目录的相对地址），也可以是url（不推荐用）
        - dest : 镜像中的绝对路径
    - 区别：
        - ADD 包含类似tar的解压功能，如果只是单纯的复制文件，推荐使用COPY

- VOLUME
    - 用于向基于镜像创建的容器添加卷，可以提供例如共享数据/数据持久化的功能

- WORKDIR
    - 在使用镜像创建容器时，在容器内部设置工作目录，ENTRYPOINT & CMD 的命令都会在这个目录下执行
    - 通常使用绝对路径

- ENV
    - 设置环境变量，与WORKDIR类似，构建 & 运行 过程中都有效

- USER
    - 指定镜像以什么样的用户去运行
    - 默认使用root用户

- ONBUILD
    - 为镜像添加触发器，当这个镜像被当作基础镜像的时候，这个指令就会被执行，这个镜像被构建时，会添加触发器中的指令


---------------------------------------------------------------

## Docker 数据卷

### 容器的数据卷
- 什么是数据卷（Data Volume）
    - 数据卷设计的目的之一就是在于数据的永久化，既数据卷的生命周期与容器的生命周期独立，既Docker不会在容器删除的时候删除其挂载的数据卷
    - docker 数据卷是一个特殊设计的目录，可以绕过联合文件目录（UFS），为一个/多个容器提供访问
    ![avatar](/docker/volume.png)
