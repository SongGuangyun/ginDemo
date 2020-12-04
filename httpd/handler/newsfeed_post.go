package handler


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ginDemo/platform/newsfeed"
	"ginDemo/httpd/response"
)

type newsfeedPostRequest struct {
    Title string `form:"title"` 
    Post string `form:"post"`
}
func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
		return func(c *gin.Context) {
			newsfeedPost := newsfeedPostRequest{}
			if err:=c.ShouldBind(&newsfeedPost);err!=nil {
				c.JSON(200,gin.H{
						"错误！":err.Error(),
				})
			}
			fmt.Println(feed)
			feed.Add(newsfeed.Item{
				Title:newsfeedPost.Title,
				Post:newsfeedPost.Post,
			})
			response.OkWithData(newsfeedPost,c)

    }
}