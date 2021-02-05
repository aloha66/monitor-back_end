package handler

import (
	"github.com/gin-gonic/gin"
	"monitor-back_end/pkg/errno"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // 为nil的时候舍弃
}

type PagationRequest struct {
	OrderMethod string `json:"orderMethod"`
	OrderType   string `json:"orderType"`
	PageNum     byte   `json:"pageNum"`
	PageSize    byte   `json:"pageSize"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, Response{Code: code, Message: message, Data: data})
}
