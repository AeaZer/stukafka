package producer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gokafka/config"
	"github.com/gokafka/controllers/producer"
)

type Command struct{}

func (c Command) Name() string {
	return "producer"
}

func (c Command) Run(conf *config.Config) {
	err := InitService(&conf.Service)
	if err != nil {
		return
	}
	newWebServer(conf.HTTP.Port)
}

func newWebServer(port int) {
	router := gin.New()
	doc := router.Group("/doc")
	{
		doc.POST("/add", producer.Add())
	}
	_ = router.Run(fmt.Sprintf(":%d", port))
}
