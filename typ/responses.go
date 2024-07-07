package typ

type SuccessResponse struct{
	Message string `json:"message"`
	Status string `json:"status"`
	Data interface{} `json:"data"`
}

type UnSuccessResponse struct{
	Message string `json:"message"`
	Status string `json:"status"`
	StatusCode int `json:"statusCode"`
}