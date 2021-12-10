package shopping

import "github.com/go-resty/resty/v2"

// RequestBasic is used for requests without pages
type RequestBasic struct {
	URL    string        `xml:"-"`
	Client *resty.Client `xml:"-"`
}

/*
==============================================================
*/

type RequestStandard struct {
	MessageID string `xml:"MessageID,omitempty"`
}

// WithMessageID adds message ID to request
//
// If you pass a value into the MessageID field in a request,
// the same value is returned in CorrelationID field in the response
func (r *RequestStandard) WithMessageID(message string) *RequestStandard {
	r.MessageID = message
	return r
}

/*
==============================================================
*/
