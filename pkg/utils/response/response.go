package response

type Response struct {
	StatusCode int         `json:"statuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

func ClientResponse(sts int, msg string, data interface{}, err interface{}) Response {
	return Response{
		StatusCode: sts,
		Message:    msg,
		Data:       data,
		Error:      err,
	}
}
