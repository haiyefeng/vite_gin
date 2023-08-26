package main

import (
	"fmt"
	"gin-vue/feishu"
	"gin-vue/fund"
	"time"

	"github.com/robfig/cron/v3"
)

// func main() {
// 	r := gin.Default()
// 	config := cors.DefaultConfig()
// 	config.AllowAllOrigins = true // 允许所有来源
// 	r.Use(cors.New(config))
// 	// CORS middleware
// 	// r.Use(func(c *gin.Context) {
// 	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 	// 	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 	// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 	// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

// 	// 	if c.Request.Method == "OPTIONS" {
// 	// 		c.AbortWithStatus(204)
// 	// 		return
// 	// 	}

// 	// 	c.Next()
// 	// })

// 	r.GET("/api/hello", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "hello world",
// 		})
// 	})

//		r.Run()
//	}
func getData() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover: ", err)
		}
	}()
	fmt.Println("start get data")
	codesData, err := fund.GetCodeVal()
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
	}

	for _, data := range *codesData {
		fundMsg := fmt.Sprintf("基金名称: %s | 基金涨幅: %s | 估值时间: %s", data.Name, data.EstimatedRange, data.EstimatedTime)
		feishu.SendData(fundMsg)
	}
}

func main() {
	// job1 := tasks.Job{
	// 	A:    0,
	// 	B:    1,
	// 	C:    "",
	// 	Shut: make(chan int, 1),
	// }
	// // 每分钟执行一次
	// go tasks.StartJob("*/1 * * * *", job1)
	// time.Sleep(time.Minute * 3)
	c := cron.New()
	_, err := c.AddFunc("50 7,9,10,11,13,14,15 * * 1,2,3,4,5", func() {
		go getData()
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}
	c.Start()
	time.Sleep(720 * time.Minute)
	defer c.Stop()
}
