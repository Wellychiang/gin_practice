package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 返回結構體
type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty`
	Count int         `json:"count,omitempty`
}

func (res *Response) Json(c *gin.Context) {
	c.JSON(http.StatusOK, res)
	return
}
