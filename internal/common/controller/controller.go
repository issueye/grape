package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	*gin.Context
}

func New(ctx *gin.Context) *Controller {
	return &Controller{
		Context: ctx,
	}
}

func (c *Controller) Success() {
	Success(c.Context)
}

func (c *Controller) SuccessByMsg(msg string) {
	SuccessByMsg(c.Context, msg)
}

func (c *Controller) SuccessByMsgf(fmtStr string, args ...any) {
	SuccessByMsgf(c.Context, fmtStr, args...)
}

func (c *Controller) SuccessData(data interface{}) {
	SuccessData(c.Context, data)
}

func (c *Controller) SuccessAutoData(req interface{}, data interface{}) {
	SuccessAutoData(c.Context, req, data)
}

func (c *Controller) SuccessPage(data interface{}) {
	SuccessPage(c.Context, data)
}

func (c *Controller) Fail() {
	Fail(c.Context, BAD_REQUEST)
}

func (c *Controller) FailByMsg(msg string) {
	FailByMsg(c.Context, msg)
}

func (c *Controller) FailByMsgf(fmtStr string, args ...any) {
	FailByMsgf(c.Context, fmtStr, args...)
}

func (c *Controller) FailBind(err error) {
	FailBind(c.Context, err)
}
