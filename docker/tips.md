# 使用docker的一些tips
- 使用docker安装centos
```bash
# 如果不加 /usr/sbin/init 启动命令的话，不会启动systemd进程的
docker run -d --name xxx  –privileged=true centos:7 /usr/sbin/init
```