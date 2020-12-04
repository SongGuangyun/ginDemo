package handler


import (
	"ginDemo/platform/newsfeed"
	"github.com/gin-gonic/gin"
	"ginDemo/httpd/response"

)
func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc{
		return func(c *gin.Context) {
				response.OkWithData(feed.GetAll(),c)
    }
}