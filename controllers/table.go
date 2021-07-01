package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetTables(c *gin.Context) {
	data := make(map[string]interface{})
	// TODO
	c.JSON(200, data)
}

func GetTable(c *gin.Context) {
	id := c.Param("id")
	// TODO
	c.JSON(200, id)
}
