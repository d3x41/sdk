// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (o *Error) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *Error) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}
