package japi

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRouter(t *testing.T) {
	r := NewRouter(`test`)
	r.GET("/hello", hello)
	r.Run(":8080")
}

func hello(c *gin.Context) {
	println("hello")
	c.JSON(200, gin.H{
		"1": 1,
		"2": 2,
		"3": 3,
	})
}
