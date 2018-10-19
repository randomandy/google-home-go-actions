package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func handleWebhook(c *gin.Context) {
	var err error

	wr := dialogflow.WebhookRequest{}
	if err = jsonpb.Unmarshal(c.Request.Body, &wr); err != nil {
		logrus.WithError(err).Error("Couldn't Unmarshal request to jsonpb")
		c.Status(http.StatusBadRequest)
		return
	}

	// contexts := wr.QueryResult.OutputContexts
	// contexts = wr.GetQueryResult().GetOutputContexts()

}

func main() {
	r := gin.Default()
	r.POST("/", handleWebhook)
	if err := r.Run("127.0.0.1:8008"); err != nil {
		panic(err)
	}
}
