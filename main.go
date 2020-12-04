package main


import (
	"github.com/gin-gonic/gin"
	"ginDemo/httpd/handler"
	"ginDemo/platform/newsfeed"

)



func main() {
	feed := newsfeed.New()
	r := gin.Default()
	r.GET("/newfeed",handler.NewsfeedGet(feed));
	r.POST("/newfeed",handler.NewsfeedPost(feed));
	r.Run()
}