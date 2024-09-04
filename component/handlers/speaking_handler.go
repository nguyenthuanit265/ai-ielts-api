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

type SpeakingHandler interface {
	GetQuestion(ctx *gin.Context)
	SetIeltsPart(ctx *gin.Context)
	Format(ctx *gin.Context)
	Submit(ctx *gin.Context)
}

type speakingHandler struct {
	speakingService services.SpeakingService
}

func NewSpeakingHandler(speakingService services.SpeakingService) SpeakingHandler {

	return &speakingHandler{
		speakingService: speakingService,
	}
}

func (h *speakingHandler) GetQuestion(ctx *gin.Context) {
	res, errSubmit := h.speakingService.GetQuestion(ctx)
	if errSubmit != nil {
		api.ErrorWithMessage(ctx, errSubmit.Code, errSubmit.Message)
		return
	}
	api.Ok(ctx, res)
}

func (h *speakingHandler) SetIeltsPart(ctx *gin.Context) {
	var req models.SpeakingSetIeltsPartRequest
	err := ctx.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		log.Errorf("Error binding request, err %v", err)
		api.Error(ctx, http.StatusBadRequest)
		return
	}

	res, errSubmit := h.speakingService.SetIeltsPart(ctx, req)
	if errSubmit != nil {
		api.ErrorWithMessage(ctx, errSubmit.Code, errSubmit.Message)
		return
	}
	api.Ok(ctx, res)
}

func (h *speakingHandler) Format(ctx *gin.Context) {
	var req models.SpeakingFormatAnswerRequest
	err := ctx.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		log.Errorf("Error binding request, err %v", err)
		api.Error(ctx, http.StatusBadRequest)
		return
	}

	res, errSubmit := h.speakingService.FormatAnswer(ctx, req)
	if errSubmit != nil {
		api.ErrorWithMessage(ctx, errSubmit.Code, errSubmit.Message)
		return
	}
	api.Ok(ctx, res)
}

func (h *speakingHandler) Submit(ctx *gin.Context) {
	var req models.SpeakingSubmitAnswerRequest
	err := ctx.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		log.Errorf("Error binding request, err %v", err)
		api.Error(ctx, http.StatusBadRequest)
		return
	}

	res, errSubmit := h.speakingService.Submit(ctx, req)
	if errSubmit != nil {
		api.ErrorWithMessage(ctx, errSubmit.Code, errSubmit.Message)
		return
	}
	api.Ok(ctx, res)
}
