package shopping

type responseStandard struct {
	Ack   string `xml:"Ack"`
	Build string `xml:"Build"`
	// CorrelationID. If you pass a value in MessageID in a request, we will return the same value
	// in CorrelationID in the response. You can use this for tracking that a response is returned
	// for every request and to match particular responses to particular requests.
	// Only returned if MessageID was used.
	CorrelationID string  `xml:"CorrelationID"`
	Errors        []Error `xml:"Errors"`
	Timestamp     string  `xml:"Timestamp"`
	Version       string  `xml:"Version"`
}

// Error is request errors (as opposed to system errors) that occur due to problems with
// business-level data (e.g., an invalid combination of arguments) that the application passed in.
type Error struct {
	// ErrorClassification. Look here
	// https://developer.ebay.com/Devzone/shopping/docs/CallRef/types/ErrorClassificationCodeType.html
	ErrorClassification string           `xml:"ErrorClassification"`
	ErrorCode           string           `xml:"ErrorCode"`
	ErrorParameters     []ErrorParameter `xml:"ErrorParameters"`
	LongMessage         string           `xml:"LongMessage"`
	// SeverityCode. Look here https://developer.ebay.com/Devzone/shopping/docs/CallRef/types/SeverityCodeType.html
	SeverityCode string `xml:"SeverityCode"`
	ShortMessage string `xml:"ShortMessage"`
}

// ErrorParameter is used by the ErrorParameters container if one or more errors or warnings occur with the call,
// and if a specific request parameter has been pinpointed as the reason why the error or warning was triggered.
type ErrorParameter struct {
	ParamID string `xml:"ParamID,attr"`
	Value   string `xml:"Value"`
}

/*
=========================================================================
*/
