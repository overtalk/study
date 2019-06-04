package mongodb_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/qinhan-shu/study/mongodb"
)

func getMongoDB() (*mongo.Database, error) {
	conf := &mongodb.MongoConf{
		Size:          10,
		Addr:          []string{"localhost:27017"},
		Username:      "test",
		Password:      "test",
		DBName:        "test",
		AuthMechanism: "SCRAM-SHA-1",
	}

	return conf.Connect()
}

func TestGet(t *testing.T) {
	db, err := getMongoDB()
	if err != nil {
		t.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	where := make(map[string]interface{})
	where["_id"] = "qinhan"
	results := db.Collection("test").FindOne(ctx, where)
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

func TestSet(t *testing.T) {
	db, err := getMongoDB()
	if err != nil {
		t.Error(err)
		return
	}

	doc := "user"
	data := make(map[string]interface{})
	data["_id"] = "qinhan"
	data["age"] = 22
	data["school"] = "SHU"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := db.Collection(doc).InsertOne(ctx, data)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("InsertedID = ", result.InsertedID)
}

func TestCount(t *testing.T) {
	db, err := getMongoDB()
	if err != nil {
		t.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	where := make(map[string]interface{})
	countNum, err := db.Collection("test").CountDocuments(ctx, where)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("num = ", countNum)
}

func TestSetMany(t *testing.T) {
	db, err := getMongoDB()
	if err != nil {
		t.Error(err)
		return
	}

	var toStore []interface{}
	for i := 0; i < 30; i++ {
		temp := make(map[string]interface{})
		temp["_id"] = fmt.Sprintf("index_%d", i)
		temp["sort"] = i
		toStore = append(toStore, temp)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := db.Collection("gettest").InsertMany(ctx, toStore)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("InsertedIDs = ", result.InsertedIDs)
}

func TestGetBySort(t *testing.T) {
	db, err := getMongoDB()
	if err != nil {
		t.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	where := make(map[string]interface{})
	value := make(map[string]interface{})
	value["$gt"] = 10
	// $gt 大于
	// $ne 不等于
	where["sort"] = value

	// 查询哪些字段
	col := make(map[string]interface{})
	col["_id"] = 1

	// 排序
	sort := make(map[string]interface{})
	sort["sort"] = -1

	cur, err := db.Collection("gettest").Find(ctx, where, &options.FindOptions{Projection: col, Sort: sort})
	if err != nil {
		t.Error(err)
		return
	}

	defer cur.Close(nil)
	var data []map[string]interface{}
	for cur.Next(nil) {
		newData := new(map[string]interface{})
		if err := cur.Decode(newData); err != nil {
			t.Error(err)
			return
		}
		data = append(data, *newData)
	}
	if err := cur.Err(); err != nil {
		t.Error(err)
		return
	}

	for _, v := range data {
		t.Log("data = ", v)
	}
}

func TestGetOR(t *testing.T) {
	db, err := getMongoDB()
	if err != nil {
		t.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	where := make(map[string]interface{})
	conditions1 := make(map[string]interface{})
	conditions2 := make(map[string]interface{})
	conditions1["a"] = 1
	conditions2["b"] = 2
	// $gt 大于
	// $ne 不等于
	where["$or"] = []map[string]interface{}{conditions1, conditions2}

	cur, err := db.Collection("test").Find(ctx, where)
	if err != nil {
		t.Error(err)
		return
	}

	defer cur.Close(nil)
	var data []map[string]interface{}
	for cur.Next(nil) {
		newData := new(map[string]interface{})
		if err := cur.Decode(newData); err != nil {
			t.Error(err)
			return
		}
		data = append(data, *newData)
	}
	if err := cur.Err(); err != nil {
		t.Error(err)
		return
	}

	for _, v := range data {
		t.Log("data = ", v)
	}
}

func TestGetAnd(t *testing.T) {
	db, err := getMongoDB()
	if err != nil {
		t.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	where := make(map[string]interface{})
	where["a"] = 1
	where["b"] = 1

	cur, err := db.Collection("test").Find(ctx, where)
	if err != nil {
		t.Error(err)
		return
	}

	defer cur.Close(nil)
	var data []map[string]interface{}
	for cur.Next(nil) {
		newData := new(map[string]interface{})
		if err := cur.Decode(newData); err != nil {
			t.Error(err)
			return
		}
		data = append(data, *newData)
	}
	if err := cur.Err(); err != nil {
		t.Error(err)
		return
	}

	for _, v := range data {
		t.Log("data = ", v)
	}
}
