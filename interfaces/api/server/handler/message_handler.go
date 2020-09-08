package handler

import (
	"github.com/gin-gonic/gin"
)

// MessageHandler is interface of handler to do messages
type MessageHandler interface {
	getAll(ctx *gin.Context)
	post(ctx *gin.Context)
	delete(ctx *gin.Context)
}
