package app

import (
	// "cloud.google.com/go/dialogflow/apiv2"
	"encoding/json"
	// "fmt"
	"google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"net/http"
)

func init() {
	http.HandleFunc("/", handle)
}

var (
	userMsg = "The air quality index in your city is foobar."
	errMsg  = "Sorry, I was unable to get air quality index in your place."
)

func handle(w http.ResponseWriter, r *http.Request) {
	// dfReq := DialogFlowRequest{}
	dfReq := dialogflow.WebhookRequest{}
	dfErr := json.NewDecoder(r.Body).Decode(&dfReq)

	if dfErr == nil && dfReq.QueryResult.Action == "location" {
		handleLocation(w, r)
		return
	}

	// if dfErr == nil && dfReq.Result.Action == "get" {
	// 	handleGet(w, r, dfReq)
	// 	return
	// }

	// resp := Music{
	//     Genre: struct {
	//         Country string
	//         Rock    string
	//     }{
	//         Country: "Taylor Swift",
	//         Rock:    "Aimee",
	//     },
	// }

	mySimpleResponse := SimpleResponse{
		TextToSpeech: "Foo.",
	}

	myItem := Item{
		SimpleResponse: mySimpleResponse,
	}

	myItems := []Item{myItem}

	myRichResponse := RichResponse{
		Items: myItems,
	}

	myGoogle := Google{
		RichResponse: myRichResponse,
	}

	myPayload := Payload{
		Google: myGoogle,
	}

	json.NewEncoder(w).Encode(WebhookResponse{
		Payload: myPayload,
	})

	// returnAPIErrorMessage(w)
}

func handleLocation(w http.ResponseWriter, r *http.Request) {

	mySimpleResponse := SimpleResponse{
		TextToSpeech: "Foo.",
	}

	myItem := Item{
		SimpleResponse: mySimpleResponse,
	}

	myItems := []Item{myItem}

	myRichResponse := RichResponse{
		Items: myItems,
	}

	myGoogle := Google{
		RichResponse: myRichResponse,
	}

	myPayload := Payload{
		Google: myGoogle,
	}

	json.NewEncoder(w).Encode(WebhookResponse{
		Payload: myPayload,
	})
}

// func handleLocation(w http.ResponseWriter, r *http.Request, dfReq DialogFlowRequest) {
// 	json.NewEncoder(w).Encode(DialogFlowLocationResponse{
// 		Speech: "PLACEHOLDER_FOR_PERMISSION",
// 		Data: DialogFlowResponseData{
// 			Google: DialogFlowResponseGoogle{
// 				ExpectUserResponse: true,
// 				IsSsml:             false,
// 				SystemIntent: DialogFlowResponseSystemIntent{
// 					Intent: "actions.intent.PERMISSION",
// 					Data: DialogFlowResponseSystemIntentData{
// 						Type:        "type.googleapis.com/google.actions.v2.PermissionValueSpec",
// 						OptContext:  "To get city for air quality check",
// 						Permissions: []string{"DEVICE_PRECISE_LOCATION"},
// 					},
// 				},
// 			},
// 		},
// 	})
// }

// func handleGet(w http.ResponseWriter, r *http.Request, dfReq DialogFlowRequest) {
// 	lat := dfReq.OriginalRequest.Data.Device.Location.Coordinates.Lat
// 	long := dfReq.OriginalRequest.Data.Device.Location.Coordinates.Long
// 	if lat == 0 || long == 0 {
// 		returnAPIErrorMessage(w)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(DialogFlowResponse{
// 		Speech: fmt.Sprintf(userMsg),
// 	})
// }

// func returnAPIErrorMessage(w http.ResponseWriter) {
// 	json.NewEncoder(w).Encode(DialogFlowResponse{
// 		Speech: errMsg,
// 	})
// }
