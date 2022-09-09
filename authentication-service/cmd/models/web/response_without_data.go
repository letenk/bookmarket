package web

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ApiResponseWithoutData(code int, status string, message string) ResponseWithoutData {
	jsonResponse := ResponseWithoutData{
		Code:    code,
		Status:  status,
		Message: message,
	}

	return jsonResponse
}
