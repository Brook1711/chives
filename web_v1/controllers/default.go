package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DanmuRequest struct {
	Token  string
	Id     string
	Author int64
	Time   float32
	Text   string
	Color  int32
	Type   int8
}

type ReturnData struct {
	Code int
	Data []DanmuRequest
}

type MainController struct {
	beego.Controller
}

type DanmuController struct {
	beego.Controller
}

func (c *DanmuController) Post() {
	var data DanmuRequest
	RequestData := c.Ctx.Input.RequestBody
	//fmt.Println(RequestData)
	json.Unmarshal(RequestData, &data)
	//fmt.Println(data)
	c.Data["json"] = data
	c.ServeJSON()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://0.0.0.0:27017"))
	if err != nil {
		return
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	//err = client.Ping(ctx, readpref.Primary())
	collection := client.Database("beego_test_2").Collection("testing_2")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//str := `{"token":"token-example","id":"vid-example","author":1000,"time":0,"text":"test","color":16777215,"type":0}`
	danmu := DanmuRequest{data.Token, data.Id, data.Author, data.Time, data.Text, data.Color, data.Type}
	insertOne, err := collection.InsertOne(ctx, danmu)
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)

}

func (c *DanmuController) Get() {
	//var (
	//	collection *mongo.Collection
	//	err        error
	//	cursor
	//)

	//c.TplName = "index.tql"
	jsoninfo := c.GetString("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://0.0.0.0:27017"))
	if err != nil {
		return
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
	collection := client.Database("beego_test_2").Collection("testing_2")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	filter := bson.D{{"id", jsoninfo}}
	cursor, err := collection.Find(context.TODO(), filter, options.Find().SetSkip(0),
		options.Find().SetLimit(2))

	if err != nil {
		fmt.Println(err)
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			fmt.Println(err)
		}
	}()

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}

	//var result DanmuRequest
	//err = collection.FindOne(ctx, filter).Decode(&result)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(result)
	//filter := bson.D{{"id",jsoninfo}}
	//var result ReturnData
	//cur, err := collection.Find(ctx, filter)
	//if err != nil { fmt.Println(err)}
	//defer cur.Close(ctx)
	//for cur.Next(ctx) {
	//	var result bson.D
	//	err := cur.Decode(&result)
	//	if err != nil { fmt.Println(err) }
	//	do something with result....
	//}
	//if err := cur.Err(); err != nil {
	//	fmt.Println(err)
	//}

	//err = collection.Find(context.TODO(), filter)
	//.Decode(&result)

	if jsoninfo == "" {
		c.Ctx.WriteString("jsoninfo is empty")
		return
	}
}

func (c *MainController) Get() {
	c.Data["Website"] = "chives.me"
	c.Data["Email"] = "brook1711@bupt.edu.cn"
	c.TplName = "index.html"
}
