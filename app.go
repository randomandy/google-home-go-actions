package app

import (
	"encoding/json"
	"fmt"
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
	dfReq := DialogFlowRequest{}
	dfErr := json.NewDecoder(r.Body).Decode(&dfReq)

	if dfErr == nil && dfReq.Result.Action == "location" {
		handleLocation(w, r, dfReq)
		return
	}

	if dfErr == nil && dfReq.Result.Action == "get" {
		handleGet(w, r, dfReq)
		return
	}

	returnAPIErrorMessage(w)
}

func handleLocation(w http.ResponseWriter, r *http.Request, dfReq DialogFlowRequest) {
	json.NewEncoder(w).Encode(DialogFlowLocationResponse{
		Speech: "PLACEHOLDER_FOR_PERMISSION",
		Data: DialogFlowResponseData{
			Google: DialogFlowResponseGoogle{
				ExpectUserResponse: true,
				IsSsml:             false,
				SystemIntent: DialogFlowResponseSystemIntent{
					Intent: "actions.intent.PERMISSION",
					Data: DialogFlowResponseSystemIntentData{
						Type:        "type.googleapis.com/google.actions.v2.PermissionValueSpec",
						OptContext:  "To get city for air quality check",
						Permissions: []string{"DEVICE_PRECISE_LOCATION"},
					},
				},
			},
		},
	})
}

func handleGet(w http.ResponseWriter, r *http.Request, dfReq DialogFlowRequest) {
	lat := dfReq.OriginalRequest.Data.Device.Location.Coordinates.Lat
	long := dfReq.OriginalRequest.Data.Device.Location.Coordinates.Long
	if lat == 0 || long == 0 {
		returnAPIErrorMessage(w)
		return
	}

	json.NewEncoder(w).Encode(DialogFlowResponse{
		Speech: fmt.Sprintf(userMsg),
	})
}

func returnAPIErrorMessage(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(DialogFlowResponse{
		Speech: errMsg,
	})
}
