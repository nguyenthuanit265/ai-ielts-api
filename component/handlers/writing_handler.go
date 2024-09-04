package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
	"main/component/api"
	"main/component/models"
	"main/component/services"
	"net/http"
)

type WritingHandler interface {
	GetQuestion(ctx *gin.Context)
	GetQuestionType(ctx *gin.Context)
	SetIeltsTask(ctx *gin.Context)
	Submit(ctx *gin.Context)
}

type writingHandler struct {
	writingService services.WritingService
}

func NewWritingHandler(writingService services.WritingService) WritingHandler {

	return &writingHandler{
		writingService: writingService,
	}
}

func (h *writingHandler) GetQuestion(ctx *gin.Context) {
	res, errSubmit := h.writingService.GetQuestion(ctx)
	if errSubmit != nil {
		api.ErrorWithMessage(ctx, errSubmit.Code, errSubmit.Message)
		return
	}
	api.Ok(ctx, res)
}

func (h *writingHandler) GetQuestionType(ctx *gin.Context) {
	res, errSubmit := h.writingService.GetQuestionType(ctx)
	if errSubmit != nil {
		api.ErrorWithMessage(ctx, errSubmit.Code, errSubmit.Message)
		return
	}
	api.Ok(ctx, res)
}

func (h *writingHandler) SetIeltsTask(ctx *gin.Context) {
	var req models.WritingSetIeltsTaskRequest
	err := ctx.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		log.Errorf("Error binding request, err %v", err)
		api.Error(ctx, http.StatusBadRequest)
		return
	}

	res, errSubmit := h.writingService.SetIeltsTask(ctx, req)
	if errSubmit != nil {
		api.ErrorWithMessage(ctx, errSubmit.Code, errSubmit.Message)
		return
	}
	api.Ok(ctx, res)
}

func (h *writingHandler) Submit(ctx *gin.Context) {
	var req models.WritingSubmitRequest
	err := ctx.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		log.Errorf("Error binding request, err %v", err)
		api.Error(ctx, http.StatusBadRequest)
		return
	}

	res, errSubmit := h.writingService.Submit(ctx, req)
	if errSubmit != nil {
		api.ErrorWithMessage(ctx, errSubmit.Code, errSubmit.Message)
		return
	}
	api.Ok(ctx, res)
}
