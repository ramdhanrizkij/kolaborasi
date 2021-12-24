package helper

type Response struct {
	Meta Meta        `json:"meta"'`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"'`
	Status  string `json:"status"'`
	Message string `json:"message"'`
}

func APIResponse(status string, message string, code int, data interface{}) Response {
	meta := Meta{
		Code:    code,
		Status:  status,
		Message: message,
	}
	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}
