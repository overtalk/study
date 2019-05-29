package mongodb

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConf : config of mongodb
type MongoConf struct {
	Size          uint16
	Addr          []string
	Username      string
	Password      string
	DBName        string
	ReplicaSet    string
	AuthMechanism string
}

// Connect : connect mongo db
func (m *MongoConf) Connect() (*mongo.Database, error) {
	if m.Size < 1 {
		log.Printf("invalid mysql pool size: %d", m.Size)
		return nil, errors.New("invalid mysql pool size")
	}
	option := options.Client()
	option.SetHosts(m.Addr)
	option.SetAuth(options.Credential{
		PasswordSet:   true,
		Username:      m.Username,
		Password:      m.Password,
		AuthMechanism: m.AuthMechanism,
	})
	option.SetMaxPoolSize(m.Size)
	option.SetReplicaSet(m.ReplicaSet)

	pool, err := mongo.NewClient(option)
	if err != nil {
		log.Printf("failed to new client to mongodb [%v]: %v", m.Addr, err)
		return nil, err
	}
	err = pool.Connect(context.Background())
	if err != nil {
		log.Printf("failed to connect to mongodb [%v]: %v", m.Addr, err)
		return nil, err
	}
	return pool.Database(m.DBName), nil
}
