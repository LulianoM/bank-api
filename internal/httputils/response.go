package httputils

type HTTPError struct {
	Error   string      `json:"type,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

type Response struct {
	Data   interface{} `json:"data,omitempty"`
	Errors []HTTPError `json:"errors,omitempty"`
}

type CustomError struct {
	Msg   string `json:"message"`
	Error string `json:"error"`
}

type ReponseError struct {
	Msg   string `json:"message"`
	Error string `json:"error"`
}

type DataWithArray struct {
	Items interface{} `json:"items"`
}

func BuildArrayResponse(items interface{}) Response {
	return Response{
		Data: DataWithArray{Items: items},
	}
}

func BuildItemResponse(item interface{}) Response {
	return Response{
		Data: item,
	}
}

func BuildErrorResponse(err error) Response {
	return Response{
		Errors: []HTTPError{
			{
				Error: err.Error(),
			},
		},
	}
}

func BuildErrorItemResponse(err error, item interface{}) Response {
	return Response{
		Errors: []HTTPError{
			{
				Error: err.Error(),
				Data:  item,
			},
		},
	}
}
