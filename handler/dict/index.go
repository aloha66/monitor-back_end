package dict

import (
	"github.com/gin-gonic/gin"
	"log"
	"monitor-back_end/handler"
	"monitor-back_end/pkg/errno"
	service "monitor-back_end/service/dict"
)

func CreateDict(c *gin.Context) {
	var form CreateDictRequest
	if err := c.ShouldBindJSON(&form); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	if err := service.CreateDict(form.Name, form.Value); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	//fmt.Println(service.CreateDict)
	handler.SendResponse(c, nil, nil)
}

func GetDictList(c *gin.Context) {
	var form GetDictRequest
	if err := c.ShouldBindJSON(&form); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	log.Println("form", form)
}
