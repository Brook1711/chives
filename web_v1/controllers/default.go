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
	Vid     string
	Author int64
	Time   float32
	Text   string
	Color  int32
	Type   int8
}

type Content struct {
	Time float64
	Type int32
	Color int32
	Author int64
	Text string
}

type ReturnDatas struct {
	Code int
	Data []interface{}
	//[[1,3,4,.5],]
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
	danmu := DanmuRequest{data.Token, data.Vid, data.Author, data.Time, data.Text, data.Color, data.Type}
	insertOne, err := collection.InsertOne(ctx, danmu)
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println("Inserted a Single Document: ", insertOne.InsertedID)

}

func (c *DanmuController) Get() {
	c.TplName = "index.html"
	//c.Ctx.WriteString("6666")
	jsoninfo := c.GetString("vid")

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
	var danmu_and_code ReturnDatas

	danmu_and_code.Code = 0

	for _, result := range results {
		fmt.Println(result)
		fmt.Println(result[1].Value)
		//fmt.Println(result[0])
		//var one_data Content
		//for i:=0;i<5;i++{
		//	one_data = result[0].time, result[0].dplayertype, result[0].color, result[0].author, result[0].text
		//}
		//one_data = make(string,)
		//one_data.append()
		one_data := make([]interface{}, 0)
		one_data = append(one_data, result[4].Value.(float64), result[7].Value.(int32),
			result[6].Value.(int32), result[3].Value.(int64), result[5].Value.(string))
		fmt.Println(one_data)
		//one_data.Time = result[4].Value.(float64)
		//one_data.Type = result[7].Value.(int32)
		//one_data.Color = result[6].Value.(int32)
		//one_data.Author = result[3].Value.(int64)
		//one_data.Text = result[5].Value.(string)

		danmu_and_code.Data = append(danmu_and_code.Data, one_data)
	}
	//c.Ctx.ResponseWriter(danmu_and_code)
	//handle results
	c.Data["json"] = danmu_and_code
	c.ServeJSON()




	//  results[0]

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
