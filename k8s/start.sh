#! /bin/bash

clear
while true
do
    echo " Start Tool"
    echo "════════════════════════════"
    echo " [1] start pvp server"
    echo " [2] start match server"
    echo " [3] clean docker images"
    echo "════════════════════════════"
    echo 
    
    read -p "Select: " option
    case "$option" in
    		1)  clear
                docker build -t k8s_test_pvp:v1 -f ./docker/pvp/Dockerfile .
                docker run -p 8081:8081 -d --name="pvp" k8s_test_pvp
                exit 1;;
    
            2)	clear
                docker build -t k8s_test_game:v1 -f ./docker/game/Dockerfile .
                docker run -p 8080:8080 -d --name="game" k8s_test_game
                exit 1;;
    
            3)	clear
                docker kill game pvp 
                docker rm game pvp 
                docker rmi k8s_test_game k8s_test_pvp 
                exit 1;;

        	*)  clear
                echo "invalid inputs[ " $option " ]" 
                echo ;;
	esac
done