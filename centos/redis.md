# centos7 中如何安装redis

```bash
yum install epel-release yum-utils
yum install http://rpms.remirepo.net/enterprise/remi-release-7.rpm
yum-config-manager --enable remi
yum install redis
systemctl start redis
systemctl enable redis
```

- 配置文件存储位置
    - /etc/redis.conf

- 再加上一行设置密码： requirepass xxxxxx

- bind 127.0.0.1 [你的IP] 
    - 这样其他的机器就可以通过本机的ip访问redis