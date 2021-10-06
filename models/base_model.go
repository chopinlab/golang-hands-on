package models

type ErrorBadRequestResponse struct {
	Message string `json:"message" example:"Bad Request"`
	Reason  string `json:"reason"  example:"Require Parameter Omitted"`
}

type ErrorNotFoundResponse struct {
	Message string `json:"message" example:"Not Found"`
	Reason  string `json:"reason"  example:"No Data"`
}

type ErrorInternalServerErrorResponse struct {
	Message string `json:"message" example:"Internal Server Error"`
	Reason  string `json:"reason"  example:"Internal Server Error"`
}

type ErrorConflictResponse struct {
	Message string `json:"message" example:"Conflict"`
	Reason  string `json:"reason"  example:"Conflict"`
}

type SuccessResponse struct {
	Message string `json:"message" example:"Success"`
}

type SuccessHealthResponse struct {
	Message string `json:"result" example:"OK"`
}

type CloudAccountIdResponse struct {
	Message string `json:"cloudAccountId" example:"39285a58-9362-xxxx-xxxx-afda095bb612"`
}

type ProductIdResponse struct {
	Message string `json:"productId" example:"39285a58-9362-xxxx-xxxx-afda095bb612"`
}

type TaskIdResponse struct {
	Message string `json:"taskId" example:"39285a58-9362-xxxx-xxxx-afda095bb612"`
}
