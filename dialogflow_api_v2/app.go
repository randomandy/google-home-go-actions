package main

import (
	"fmt"
	"log"
	"net/http"

	// "github.com/davecgh/go-spew/spew"
	"github.com/Depado/articles/code/dialogflow/cocktail"
	"github.com/gin-gonic/gin"
	df "github.com/leboncoin/dialogflow-go-webhook"
	"github.com/sirupsen/logrus"
)

type BasicCard struct {
	Title         string     `json:"title,omitempty"`         // Optional. The title of the card.
	Subtitle      string     `json:"subtitle,omitempty"`      // Optional. The subtitle of the card.
	FormattedText string     `json:"formattedText,omitempty"` // Required, unless image is present. The body text of the card.
	Image         *CardImage `json:"image,omitempty"`         // Optional. The image for the card.
	// Buttons       []CardButton `json:"buttons,omitempty"`       // Optional. The collection of card buttons.
}

func (bc BasicCard) GetKey() string {
	return "basicCard"
}

type CardImage struct {
	ImageURI          string `json:"imageUri,omitempty"`
	AccessibilityText string `json:accessibility_text,omitempty`
}

type searchParams struct {
	Alcohol   string `json:"alcohol"`
	DrinkType string `json:"drink-type"`
	Name      string `json:"name"`
}

func webhook(c *gin.Context) {

	var err error
	var dfr *df.Request

	if err = c.BindJSON(&dfr); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// spew.Dump(dfr)

	switch dfr.QueryResult.Action {
	case "search":
		log.Println("Search action detected")
		c.JSON(http.StatusOK, gin.H{})
	case "random":
		log.Println("Random action detected")
		random(c, dfr)
	default:
		log.Println("Unknown action")
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{})
}

func search(c *gin.Context, dfr *df.Request) {
	var err error
	var p searchParams

	if err = dfr.GetParams(&p); err != nil {
		logrus.WithError(err).Error("Couldn't get parameters")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	spew.Dump(p)

	c.JSON(http.StatusOK, gin.H{})
}

func random(c *gin.Context, dfr *df.Request) {
	var err error
	var d *cocktail.FullDrink

	if d, err = cocktail.C.GetRandomDrink(); err != nil {
		logrus.WithError(err).Error("Coudln't get random drink")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fmt.Println(d)

	out := fmt.Sprintf("I found that cocktail : %s", d.StrDrink)
	dff := &df.Fulfillment{
		FulfillmentMessages: df.Messages{
			{
				RichMessage: df.Text{
					Text: []string{out},
				},
			},
			// df.ForGoogle(cardFromDrink(d)),
			df.ForGoogle(df.SingleSimpleResponse(out, out)),
		},
	}

	c.JSON(http.StatusOK, dff)
}

func cardFromDrink(d *cocktail.FullDrink) BasicCard {
	card := BasicCard{
		Title:         d.StrDrink,
		FormattedText: d.StrInstructions,
		Image: &CardImage{
			ImageURI:          d.StrDrinkThumb,
			AccessibilityText: "cocktail",
		},
	}
	return card
}

func main() {
	r := gin.Default()
	r.POST("/", webhook)
	if err := r.Run("127.0.0.1:8008"); err != nil {
		panic(err)
	}
}
