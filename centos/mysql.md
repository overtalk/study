# centos7 中如何安装mysql

```bash
rpm -ivh https://dev.mysql.com/get/mysql57-community-release-el7-11.noarch.rpm
yum -y install mysql-community-server
systemctl start mysqld
systemctl enable mysqld

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

- 配置文件位置
    - /etc/my.cnf

## mysql 用户管理的常用操作
- 查看所有用户
```bash
select host,user,authentication_string from mysql.user;
```

- 查看用户有哪些权限
```bash
show grants for py;
```

- 创建用户并且授权（需要使用实例级账户登录后操作）
    - 语法如下：
    ```bash
    grant 权限列表 on 数据库 to '用户名'@'访问主机' identified by '密码';

    # 创建新用户，用户名为py，密码为123
    # 操作python数据库的所有对象python.*
    # 访问主机通常使用百分号%表示此账户可以使用任何ip的主机登录访问此数据库
    # 访问主机可以设置成localhost或具体的ip，表示只允许本机或特定主机访问
    grant all privileges on python.* to 'py'@'%' identified by '123';
    ```

- 授权
    - 语法如下：
    ```bash
    grant 权限名称 on 数据库 to 账户1,账户2,... with grant option;
    ```

- 修改密码
```bash
# 需要使用root登陆
update user set authentication_string=password('新密码') where user='用户名';
# 刷新权限
flush privileges
```

- 删除账户
```bash
# 使用root登录
delete from user where user='用户名';
delete from user where user='py';

# 操作结束之后需要刷新权限
flush privileges
```