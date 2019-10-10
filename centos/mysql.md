# centos7 中如何安装mysql

```bash
sudo rpm -ivh https://dev.mysql.com/get/mysql57-community-release-el7-11.noarch.rpm

sudo yum -y install mysql-community-server

sudo systemctl start mysqld

sudo systemctl enable mysqld

```

# 获取密码
```bash
cat /var/log/mysqld.log | grep -i 'temporary password'
```

# 使用上面的密码登陆进去之后设置密码
```bash
# 修改validate_password_policy参数的值
set global validate_password_policy=0; 

# 再修改密码的长度
set global validate_password_length=1;

# 更新密码
ALTER USER 'root'@'localhost' IDENTIFIED BY 'root123';
```