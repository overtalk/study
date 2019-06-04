package json

import (
	"encoding/json"
)

// GetJSON 拿到bytes数据
func GetJSON() ([]byte, error) {
	confJSON := make(map[string]interface{})
	confJSON["pvp_server_addrs"] = []string{"pvp服务器地址1", "pvp服务器地址2"}
	confJSON["match_server_addrs"] = []string{"match服务器地址"}
	confJSON["consul_addr"] = "consul地址"
	// gm
	gmJSON := make(map[string]interface{})
	gmJSON["url"] = "后台配置url"
	gmJSON["username"] = "后台配置用户名"
	gmJSON["token"] = "后台配置token"
	confJSON["gm_repo"] = gmJSON
	// redis
	redisJSON := make(map[string]interface{})
	redisJSON["addrs"] = []string{"Redis地址1", "Redis地址2"}
	redisJSON["password"] = "Redis密码"
	redisJSON["maxconn"] = 30
	confJSON["redis"] = redisJSON
	// mysql
	mysqlJSON := make(map[string]interface{})
	mysqlJSON["addr"] = "MySQL地址"
	mysqlJSON["username"] = "MySQL用户名"
	mysqlJSON["password"] = "MySQL密码"
	mysqlJSON["database"] = "MySQL数据库"
	mysqlJSON["sharding_size"] = 30
	mysqlJSON["maxconn"] = 30
	mysqlJSON["load_all_player_on_startup"] = true
	confJSON["mysql"] = mysqlJSON
	// mongo
	mongoJSON := make(map[string]interface{})
	mongoJSON["addrs"] = []string{"MONGODB地址1", "MONGODB地址2"}
	mongoJSON["username"] = "MONGODB用户名"
	mongoJSON["password"] = "MONGODB密码"
	mongoJSON["database"] = "MONGODB数据库"
	mongoJSON["maxconn"] = 30
	mongoJSON["replica_set"] = "MONGODB的REPLICASET"
	mongoJSON["auth_mechanism"] = "MONGODB加密方式"
	confJSON["mongodb"] = mongoJSON
	// smtps
	smtpJSON := make(map[string]interface{})
	smtpJSON["enable"] = true
	smtpJSON["addr"] = "SMTP地址"
	smtpJSON["host"] = "SMTP的HOST"
	smtpJSON["username"] = "SMTP用户名"
	smtpJSON["password"] = "SMTP密码"
	smtpJSON["to"] = []string{"SMTP发送对象"}
	smtpJSON["from"] = "SMTP发送者"
	smtpJSON["content_type"] = "SMTP类型"
	smtpJSON["message_template"] = "SMTP模版"
	confJSON["smtp"] = smtpJSON

	return json.Marshal(confJSON)
}
