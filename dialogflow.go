package app

// // DialogFlowRequestV2 struct
// type DialogFlowRequestV2 struct {
// 	Session                     string                        `json:"session"`
// 	ResponseId                  string                        `json:"responseId"`
// 	QueryResult                 QueryResultV2                 `json:"queryResult"`
// 	OriginalDetectIntentRequest OriginalDetectIntentRequestV2 `json:"originalDetectIntentRequest"`
// }

// // QueryResultV2 struct
// type QueryResultV2 struct {
// 	Session    string `json:"session"`
// 	ResponseId string `json:"responseId"`

// 	QueryText                   string              `json:"queryText"`
// 	LanguageCode                string              `json:"languageCode"`
// 	SpeechRecognitionConfidence float               `json:"speechRecognitionConfidence"`
// 	Action                      string              `json:"action"`
// 	Parameters                  Parameters          `json:"parameters"`
// 	AllRequiredParamsPresent    bool                `json:"allRequiredParamsPresent"`
// 	FulfillmentText             string              `json:"fulfillmentText"`
// 	FulfillmentMessages         FulfillmentMessages `json:"fulfillmentMessages"`
// 	WebhookSource               string              `json:"webhookSource"`
// 	WebhookPayload              WebhookPayload      `json:"webhookPayload"`
// 	OutputContexts              OutputContexts      `json:"outputContexts"`
// 	Intent                      Intent              `json:"intent"`
// 	IntentDetectionConfidence   float               `json:"intentDetectionConfidence"`
// 	DiagnosticInfo              DiagnosticInfo      `json:"diagnosticInfo"`
// }

type WebhookResponse struct {
	FulfillmentText     string `json:"fulfillmentText"`
	FulfillmentMessages struct {
		Text string `json:"Text"`
	} `json:"fulfillmentMessages"`
	Source         string  `json:"source"`
	Payload        Payload `json:"payload"`
	OutputContexts struct {
		Name          string `json:"name"`
		LifespanCount string `json:"lifespanCount"`
	} `json:"outputContexts"`
	FollowupEventInput struct {
		Name         string `json:"name"`
		LanguageCode string `json:"languageCode"`
	} `json:"followupEventInput"`
}

// Payload struct
type Payload struct {
	Google Google `json:"google"`
}

type Google struct {
	ExpectUserResponse string       `json:"expectUserResponse"`
	RichResponse       RichResponse `json:"richResponse"`
}

type RichResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	SimpleResponse SimpleResponse `json:"simpleResponse"`
}

type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech"`
}

// DialogFlowRequest struct
type DialogFlowRequest struct {
	Result struct {
		Action string `json:"action"`
	} `json:"result"`
	OriginalRequest DialogFlowOriginalRequest `json:"originalRequest"`
}

// DialogFlowOriginalRequest struct
type DialogFlowOriginalRequest struct {
	Data DialogFlowOriginalRequestData `json:"data"`
}

// DialogFlowOriginalRequestData struct
type DialogFlowOriginalRequestData struct {
	Device DialogFlowOriginalRequestDevice `json:"device"`
}

// DialogFlowOriginalRequestDevice struct
type DialogFlowOriginalRequestDevice struct {
	Location DialogFlowOriginalRequestLocation `json:"location"`
}

// DialogFlowOriginalRequestLocation struct
type DialogFlowOriginalRequestLocation struct {
	Coordinates DialogFlowOriginalRequestCoordinates `json:"coordinates"`
}

// DialogFlowOriginalRequestCoordinates struct
type DialogFlowOriginalRequestCoordinates struct {
	Lat  float32 `json:"latitude"`
	Long float32 `json:"longitude"`
}

// DialogFlowResponse struct
type DialogFlowResponse struct {
	Speech string `json:"speech"`
}

// DialogFlowLocationResponse struct
type DialogFlowLocationResponse struct {
	Speech string                 `json:"speech"`
	Data   DialogFlowResponseData `json:"data"`
}

// DialogFlowResponseData struct
type DialogFlowResponseData struct {
	Google DialogFlowResponseGoogle `json:"google"`
}

// DialogFlowResponseGoogle struct
type DialogFlowResponseGoogle struct {
	ExpectUserResponse bool                           `json:"expectUserResponse"`
	IsSsml             bool                           `json:"isSsml"`
	SystemIntent       DialogFlowResponseSystemIntent `json:"systemIntent"`
}

// DialogFlowResponseSystemIntent struct
type DialogFlowResponseSystemIntent struct {
	Intent string                             `json:"intent"`
	Data   DialogFlowResponseSystemIntentData `json:"data"`
}

// DialogFlowResponseSystemIntentData struct
type DialogFlowResponseSystemIntentData struct {
	Type        string   `json:"@type"`
	OptContext  string   `json:"optContext"`
	Permissions []string `json:"permissions"`
}
