package dict

import "monitor-back_end/handler"

type CreateDictRequest struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type GetDictParamsRequest struct {
	Name  string `json:"name" `
	Value string `json:"value"`
}

type GetDictRequest struct {
	handler.PagationRequest
	Parameters GetDictParamsRequest `json:"parameters"`
}
