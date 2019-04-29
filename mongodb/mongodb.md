# Mongodb

## install
- 通过brew安装即可 : brew install mongoldb
- 配置文件在 /usr/local/etc/mongod.conf 
- 启动mongoDB : 
    - cd /usr/local/etc
    - mongod -f mongodb.conf
- Daemon方式运行mongoDB
    - 我们使用--fork参数可以将mongodb的服务放在后台运行，这样相对比较安全。--fork参数是和--logpath参数一起使用的
    - mongod -dbpath "${dbpath}" -fork -logpath "${logpath}"
        - dbpath, logpath 就是数据库路径以及日志的路径

## golang连接mongodb
- 见代码