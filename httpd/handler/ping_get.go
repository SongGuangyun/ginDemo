package handler

import (
	"github.com/gin-gonic/gin"

)
func PingGet(message string) gin.HandlerFunc{
		return func(c *gin.Context) {
        c.JSON(200,map[string]string {
			"hello":"world",
		})
    }
}