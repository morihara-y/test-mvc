package handler

import (
	"github.com/gin-gonic/gin"
)

// MessageHandler is interface of handler to do messages
type MessageHandler interface {
	GetAll(ctx *gin.Context)
	Post(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
