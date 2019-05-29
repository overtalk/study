package mongodb_test

import (
	// "context"
	"testing"
	// "time"

	"github.com/qinhan-shu/study/mongodb"
)

func TestConnect(t *testing.T) {
	conf := &mongodb.MongoConf{
		Size:          10,
		Addr:          []string{"localhost:27017"},
		Username:      "test",
		Password:      "test",
		DBName:        "test",
		AuthMechanism: "SCRAM-SHA-1",
	}

	db, err := conf.Connect()
	if err != nil {
		t.Error(err)
		return
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	where := make(map[string]interface{})
	where["q"] = "q"
	results := db.Collection("aaa").FindOne(nil, where)
	err = results.Err()
	if err != nil {
		t.Error(err)
		return
	}

	newData := new(map[string]interface{})
	if err := results.Decode(newData); err != nil {
		t.Error(err)
		return
	}
	t.Log(newData)

}
