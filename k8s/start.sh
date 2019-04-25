#! /bin/bash

tag=$1

if [ $tag == "game" ]
then
    docker build -t "k8s_test_game" -f ./docker/game/Dockerfile .
    docker run -p 8080:8080 -d --name="game" k8s_test_game
    exit 1
fi

if [ $tag == "pvp" ]
then
    docker build -t "k8s_test_pvp" -f ./docker/pvp/Dockerfile .
    docker run -p 8081:8081 -d --name="pvp" k8s_test_pvp
    exit 1
fi

if [ $tag == "clean" ]
then
    docker kill game pvp 
    docker rm game pvp 
    docker rmi k8s_test_game k8s_test_pvp 
    exit 1
fi

echo "非法参数: [ " $tag " ]"
echo "  args : 
        - pvp   : pvp服务器
        - game  : game服务器
        - clean : 清除"