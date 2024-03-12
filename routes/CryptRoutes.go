package routes

import "github.com/gin-gonic/gin"

type StructKey struct {
	Guid   string `json:"guid"`
	Typ    string `json:"typ" binding:"required"`
	Desc   string `json:"desc"`
	Value1 string `json:"value1" binding:"required"`
	Value2 string `json:"value2"`
	Value3 string `json:"value3"`
	Value4 string `json:"value4"`
}

func SetKey(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
