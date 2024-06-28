package ginLib

type ResponseModel struct {
	Error interface{} `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
}
