# 文件上传以及下载

- 下载服务器内容（rsync）
    - rsync -zP 用户名@服务器地址:服务器里的文件路径  本地存放的文件路径
        - z : 进行压缩
        - P : 保持文件权限
        - eg : rsync -zP root@sausage-aliyun-rew:/data/rdb/06_25_snapshot/rdb-0 /Users/qinhan/Downloads/06-25-snapshot/

- 上传到服务器
    - scp 本地文件路径 用户名@服务器地址:服务器里的文件路径
	    - eg : scp /Users/qinhan/Downloads/06-25-snapshot/06_25_cleaned root@sausage-aliyun-rew:/data/rdb/06_25_snapshot/