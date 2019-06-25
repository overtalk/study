# Redis 数据的导入与导出

- 数据导入导出的工具：Redis-port
    - https://github.com/CodisLabs/redis-port

- dump 将redis中的数据导入rdb文件（本地Redis）
    - redis-port dump -f 127.0.0.1:6379 -o xxx(输出文件)
- restore 将rdb文件数据导入redis中（本地Redis）
    - redis-port restore -i xxx（输入文件） -t 127.0.0.1:6379 