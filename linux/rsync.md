# rsync 的用法

## OPTION参数
- z : 进行压缩
- P : 保持文件权限

- 下载服务器内容
    - rsync [OPTION]... [USER@]HOST:SRC DEST
        - eg : rsync -zP ./1.txt root@sausage-intranet-12:/root/

- 同步文件到服务器
    - rsync [OPTION]... SRC [USER@]HOST:DEST
        - eg : rsync -zP root@sausage-aliyun-rew:/data/rdb/06_25_snapshot/rdb-0 /Users/qinhan/Downloads/06-25-snapshot/