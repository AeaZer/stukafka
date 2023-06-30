package producer

import (
	"github.com/gin-gonic/gin"
	"github.com/gokafka/service/kafka"
	"net/http"
)

type addParams struct {
	Contents []string `json:"contents"`

	// Topic   string `json:"topic"`
}

func Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		p, err := newAddParams(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		contents := make([]string, 0, len(p.Contents))
		for i := range p.Contents {
			contents = append(contents, p.Contents[i])
		}
		err = kafka.WriteMessages(contents)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "kafka service disappear")
		}
		c.JSON(http.StatusCreated, "kafka message success created")
	}
}

func newAddParams(c *gin.Context) (*addParams, error) {
	var p addParams
	err := c.BindJSON(&p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
