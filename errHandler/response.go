package errHandler

import (
	"github.com/gin-gonic/gin"
	"luciana/errHandler/code"
	"luciana/model"
	"net/http"
)

func ResponseError(c *gin.Context, code code.Code) {
	re := &model.ResponseContent{
		Code:    code.Code(),
		Msg:     code.Msg(),
		Content: nil,
	}
	c.JSON(http.StatusOK, re)
}

func ResponseMsg(c *gin.Context, code code.Code, msg interface{}) {
	re := &model.ResponseContent{
		Code:    code.Code(),
		Msg:     code.Msg(),
		Content: msg,
	}
	c.JSON(http.StatusOK, re)
}

func ResponseSuccess(c *gin.Context, content interface{}) {
	re := &model.ResponseContent{
		Code:    code.Success.Code(),
		Msg:     code.Success.Msg(),
		Content: content,
	}
	c.JSON(http.StatusOK, re)
}
