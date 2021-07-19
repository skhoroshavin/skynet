package http

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Error(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{"err": err.Error()})
	log.Printf(`%s`, err.Error())
}
