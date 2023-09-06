package models

type Error struct {
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
}
type Paging struct {
	CursorQuery string `json:"cursorQuery,omitempty"`
	NextPageURI string `json:"nextPageUri,omitempty"`
}
type Request struct {
	Method string `json:"method,omitempty"`
	URI    string `json:"uri,omitempty"`
}
type Meta struct {
	Count        int     `json:"count,omitempty"`
	Error        Error   `json:"error,omitempty"`
	Paging       Paging  `json:"paging,omitempty"`
	Request      Request `json:"request,omitempty"`
	ResponseCode int     `json:"responseCode,omitempty"`
}

type EventMsgSpoolUsageThreshold struct {
	ClearPercent int `json:"clearPercent,omitempty"`
	ClearValue   int `json:"clearValue,omitempty"`
	SetPercent   int `json:"setPercent,omitempty"`
	SetValue     int `json:"setValue,omitempty"`
}

type ErrorResponse struct {
	Meta struct {
		Error struct {
			Description string `json:"description"`
		} `json:"error"`
		ResponseCode int `json:"responseCode"`
	} `json:"meta"`
}
