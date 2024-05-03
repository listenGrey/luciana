package controller

import (
	"github.com/gin-gonic/gin"
	"luciana/errHandler"
	"luciana/errHandler/code"
	"luciana/logic"
	"luciana/model"
)

func RequestHandler(c *gin.Context) {
	var request *model.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		errHandler.ResponseMsg(c, code.InvalidParams, "Invalid JSON")
		return
	}

	prompt := request.Prompt
	res, err := logic.Generate(prompt)
	if err != nil {
		errHandler.ResponseMsg(c, code.InvalidParams, "Generate error")
		return
	}

	errHandler.ResponseSuccess(c, res)
}
