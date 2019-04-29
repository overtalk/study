package mongodb

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBInfo struct {
	Size          uint16
	Addr          []string
	Username      string
	Password      string
	DBName        string
	ReplicaSet    string
	AuthMechanism string
}

type Pool = mongo.Database

func Connect(info DBInfo) (*Pool, error) {
	if info.Size < 1 {
		log.Printf("invalid mysql pool size: %d", info.Size)
		return nil, errors.New("invalid mysql pool size")
	}
	option := options.Client()
	option.SetHosts(info.Addr)
	option.SetAuth(options.Credential{
		PasswordSet:   true,
		Username:      info.Username,
		Password:      info.Password,
		AuthMechanism: info.AuthMechanism,
	})
	option.SetMaxPoolSize(info.Size)
	option.SetReplicaSet(info.ReplicaSet)

	pool, err := mongo.NewClient(option)
	if err != nil {
		log.Printf("failed to new client to mongodb [%v]: %v", info.Addr, err)
		return nil, err
	}
	err = pool.Connect(context.Background())
	if err != nil {
		log.Printf("failed to connect to mongodb [%v]: %v", info.Addr, err)
		return nil, err
	}
	database := pool.Database(info.DBName)
	return database, nil
}
