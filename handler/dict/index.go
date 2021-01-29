package dict

import (
	"github.com/gin-gonic/gin"
	"monitor-back_end/handler"
	"monitor-back_end/pkg/errno"
	service "monitor-back_end/service/dict"
)

func CreateDict(c *gin.Context) {
	var form CreateDictRequest
	if err := c.ShouldBindJSON(&form); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
	}

	service.CreateDict(form.Name, form.Value)
	//fmt.Println(service.CreateDict)

}
