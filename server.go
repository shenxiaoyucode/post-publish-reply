package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"pro/cotroller"
	"pro/repository"
	"strconv"
)

func main() {
	if err := Init("./data/"); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := cotroller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})

	r.POST("/community/page/publish/topic", func(context *gin.Context) {

		title := context.PostForm("title")
		content := context.PostForm("content")

		data, err2 := cotroller.PublishTopic(title, content)
		if err2 != nil {
			return
		}

		context.JSON(200, data)
	})

	r.POST("/community/page/publish/post", func(context *gin.Context) {
		pid, _ := strconv.ParseInt(context.PostForm("pid"), 10, 64)
		content := context.PostForm("content")

		data, err2 := cotroller.PublishPost(pid, content)
		if err2 != nil {
			return
		}

		context.JSON(200, data)
	})

	err := r.Run()
	if err != nil {
		return
	}
}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		return err
	}
	return nil
}
