package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/morihara-y/test-mvc/domain/model"
	"github.com/morihara-y/test-mvc/domain/service"
)

type messageHandlerImpl struct {
	service service.MessageService
}

// NewMessageHandlerImpl returns MessageHandler
func NewMessageHandlerImpl(service service.MessageService) MessageHandler {
	return &messageHandlerImpl{
		service: service,
	}
}

func (m *messageHandlerImpl) GetAll(ctx *gin.Context) {
	result := m.service.FetchAll()
	ctx.JSON(http.StatusOK, result)
}

func (m *messageHandlerImpl) Post(ctx *gin.Context) {
	var message model.MessageTrn
	ctx.BindJSON(&message)
	m.service.Insert(&message)
	ctx.JSON(http.StatusOK, message)
}

func (m *messageHandlerImpl) Delete(ctx *gin.Context) {
	m.service.DeleteAll()
	ctx.JSON(http.StatusOK, "success")
}
