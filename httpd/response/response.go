package response

import (
	"net/http"
	"github.com/gin-gonic/gin"

)
type Response struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	Msg string  `json:"msg"`

}

func Result(code int,data interface{},msg string,c *gin.Context) {
    c.JSON(http.StatusOK,Response{
        code,
        data,
        msg,
    })
}

func Ok(c *gin.Context) {
    Result(0,map[string]interface{}{},"操作成功",c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(0, data, "操作成功", c)
}