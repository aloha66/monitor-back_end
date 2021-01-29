package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"monitor-back_end/handler"
	"monitor-back_end/pkg/errno"
)

func Create(c *gin.Context) {
	var r CreateRequest
	//ShouldBind 交由开发者处理请求结果 Bind会直接400
	if err := c.ShouldBind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	admin2 := c.Param("username")
	desc := c.Query("desc")
	contentType := c.GetHeader("Content-Type")
	log.Println("username", admin2, desc, contentType)

	log.Println("r.Username, r.Password", r.Username, r.Password)
	if r.Username == "" {
		handler.SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		return
	}

	if r.Password == "" {
		handler.SendResponse(c, fmt.Errorf("password is empty"), nil)
		return
	}

	res := CreateResponse{Username: r.Username}

	handler.SendResponse(c, nil, res)

}
