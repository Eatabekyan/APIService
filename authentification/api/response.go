package api

type Response struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func MakeResponse(message string, token string) Response {
	return Response{
		Message: message,
		Token:   token,
	}
}
