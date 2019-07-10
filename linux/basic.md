# linux 基础命令


- ls : 列出目录
- cd : 切换目录
- pwd : 显示目前的目录
- mkdir : 创建一个新的目录
- rmdir : 删除一个空的目录
- cp : 复制文件或目录
- rm : 移除文件或目录
- mv : 移动文件与目录，或修改文件与目录的名称

## 例子
- 创建多层目录
    - -p : 帮助你直接将所需要的目录(包含上一级目录)递归创建起来！
        - mkdir -p test1/test2/test3/test4
- 删除多层空目录
    - rmdir -p test1/test2/test3/test4
- 复制目录
    - -r : 递归持续复制，用於目录的复制行为；(常用)
        - cp -r  /Users/qinhan/Downloads/07-09-snapshot ./test
- 移动目录、文件 / 或者用于改名字
    - -f ：force 强制的意思，如果目标文件已经存在，不会询问而直接覆盖；
        -  mv -f  ~/Downloads/2.txt 2.txt
    - 修改名称
        - mv mvtest mvtest2