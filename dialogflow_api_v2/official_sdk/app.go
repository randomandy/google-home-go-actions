package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/sirupsen/logrus"
	"google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"log"
	"net/http"
)

func handleWebhook(c *gin.Context) {
	var err error

	wr := dialogflow.WebhookRequest{}

	var unmar jsonpb.Unmarshaler
	// Workaround for `unknown field "version" in dialogflow.OriginalDetectIntentRequest`
	unmar.AllowUnknownFields = true

	if err = unmar.Unmarshal(c.Request.Body, &wr); err != nil {
		logrus.WithError(err).Error("Couldn't Unmarshal request to jsonpb")
		c.Status(http.StatusBadRequest)
		return
	}

	switch wr.QueryResult.Action {
	case "random":
		random(c, wr)
	case "search":
		search(c, wr)
	default:
		log.Println("Unknown action")
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

func search(c *gin.Context, wr dialogflow.WebhookRequest) {
	log.Println("Search action detected")

	//textMessage := dialogflow.Intent_Message_Text{
	//	Text: []string{"foo", "bar"},
	//}

	//basicCard := dialogflow.Intent_Message_BasicCard{
	//	Title: "My Title",
	//	Subtitle: "My Subtitle",
	//	Image: &dialogflow.Intent_Message_Image{
	//		ImageUri: "http://...",
	//		AccessibilityText: "My Accessibility Text",
	//	},
	//	Buttons: []*dialogflow.Intent_Message_BasicCard_Button{
	//		{
	//			Title: "My Button",
	//			OpenUriAction: &dialogflow.Intent_Message_BasicCard_Button_OpenUriAction{
	//				Uri: "http://...",
	//			},
	//		},
	//	},
	//}

	fullfillment := dialogflow.WebhookResponse{
		FulfillmentText: "foobar",
		//FulfillmentMessages: []*dialogflow.Intent_Message{&basicCard},
	}

	c.JSON(http.StatusOK, fullfillment)
}

func random(c *gin.Context, wr dialogflow.WebhookRequest)  {
	log.Println("Random action detected")

	//textMessage := dialogflow.Intent_Message_Text{
	//	Text: []string{"foo", "bar"},
	//}

	//basicCard := dialogflow.Intent_Message_BasicCard{
	//	Title: "My Title",
	//	Subtitle: "My Subtitle",
	//	Image: &dialogflow.Intent_Message_Image{
	//		ImageUri: "http://media.petsathome.com/wcsstore/pah-cas01//c/small_petL.png",
	//		AccessibilityText: "My Accessibility Text",
	//	},
	//	Buttons: []*dialogflow.Intent_Message_BasicCard_Button{
	//		{
	//			Title: "My Button",
	//			OpenUriAction: &dialogflow.Intent_Message_BasicCard_Button_OpenUriAction{
	//				Uri: "http://www.google.com",
	//			},
	//		},
	//	},
	//}

	fullfillment := dialogflow.WebhookResponse{
		FulfillmentText: "foobar",
		FulfillmentMessages: []*dialogflow.Intent_Message{
			//{
			//	Message: &dialogflow.Intent_Message_BasicCard_{
			//		BasicCard: &basicCard,
			//	},
			//},
			//{
			//	Message: &dialogflow.Intent_Message_Text_{
			//		Text: &textMessage,
			//	},
			//},
			{
				Message: &dialogflow.Intent_Message_SimpleResponses_{
					SimpleResponses: &dialogflow.Intent_Message_SimpleResponses{
						SimpleResponses: []*dialogflow.Intent_Message_SimpleResponse{
							{
								TextToSpeech: "Foobar",
								DisplayText: "Barfoo",
							},
						},
					},
				},
			},
		},
	}

	c.PureJSON(http.StatusOK, fullfillment)
	//c.JSON(http.StatusOK, fullfillment)
}