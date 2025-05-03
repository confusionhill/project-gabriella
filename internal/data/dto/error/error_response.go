package error

import "net/url"

type ErrorResponseDTO struct {
	Code            ErrorCode
	Reason          string
	Message         string
	Action          string
	Status          string
	StrErr          string
	StrReason       string
	StrButtonName   string
	StrButtonAction string
	StrMsg          string
}

func (e *ErrorResponseDTO) ToStringValues() string {
	values := url.Values{}
	values.Set("code", string(e.Code))
	values.Set("reason", e.Reason)
	values.Set("message", e.Message)
	values.Set("action", e.Action)
	values.Set("status", e.Status)
	values.Set("strErr", e.StrErr)
	values.Set("strReason", e.StrReason)
	values.Set("strButtonName", e.StrButtonName)
	values.Set("strButtonAction", e.StrButtonAction)
	values.Set("strMsg", e.StrMsg)

	return "&" + values.Encode()
}
