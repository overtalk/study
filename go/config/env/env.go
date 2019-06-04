package env

import (
	"log"
	"os"
)

// SetEnv ： 在环境变量中设置项目配置
func SetEnv() {
	envs := make(map[string]string)
	// server addr
	envs["PVP_SERVER_ADDRS"] = "pvp服务器地址1,pvp服务器地址2"
	envs["MATCH_SERVER_ADDRS"] = "match服务器地址"
	envs["CONSUL_ADDRESS"] = "consul地址"
	// gm
	envs["GM_REPO_URL"] = "后台配置url"
	envs["GM_REPO_USERNAME"] = "后台配置用户名"
	envs["GM_REPO_TOKEN"] = "后台配置token"
	// redis
	envs["REDIS_ADDR"] = "Redis地址1,Redis地址2"
	envs["REDIS_PASS"] = "Redis密码"
	envs["REDIS_MAXCONN"] = "30"
	// mysql
	envs["MYSQL_ADDR"] = "MySQL地址"
	envs["MYSQL_USER"] = "MySQL用户名"
	envs["MYSQL_PASS"] = "MySQL密码"
	envs["MYSQL_DATABASE"] = "MySQL数据库"
	envs["MYSQL_SHARDING_SIZE"] = "30"
	envs["MYSQL_MAXCONN"] = "30"
	envs["MYSQL_LOAD_ALL_PLAYER_ON_STARTUP"] = "true"
	// mongo
	envs["MONGODB_ADDR"] = "MONGODB地址1,MONGODB地址2"
	envs["MONGODB_USERNAME"] = "MONGODB用户名"
	envs["MONGODB_PASSWORD"] = "MONGODB密码"
	envs["MONGODB_DATABASE"] = "MONGODB数据库"
	envs["MONGODB_MAXCONN"] = "30"
	envs["MONGODB_REPLICASET"] = "MONGODB的REPLICASET"
	envs["MONGODB_AUTHMECHANISM"] = "MONGODB加密方式"
	// smtps
	envs["SMTP_ENABLE"] = "true"
	envs["SMTP_ADDR"] = "SMTP地址"
	envs["SMTP_HOST"] = "SMTP的HOST"
	envs["SMTP_USERNAME"] = "SMTP用户名"
	envs["SMTP_PASSWORD"] = "SMTP密码"
	envs["SMTP_MAILTO"] = "SMTP发送对象"
	envs["SMTP_MAILFROM"] = "SMTP发送者"
	envs["SMTP_CONTENT_TYPE"] = "SMTP类型"
	envs["SMTP_MSG_TEMPLATE"] = "SMTP模版"

	for k, v := range envs {
		if err := os.Setenv(k, v); err != nil {
			log.Fatalf("failed to set env pair [%s , %s]", k, v)
		}
	}
}
