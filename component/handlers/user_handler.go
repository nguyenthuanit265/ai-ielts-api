package handlers

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/gin-gonic/gin/binding"
//	log "github.com/sirupsen/logrus"
//	"main/component/api"
//	"main/component/models"
//	"main/component/services"
//	"net/http"
//)
//
//type UserHandler interface {
//	FindUserProfile(ctx *gin.Context)
//	ChangePassword(ctx *gin.Context)
//}
//
//type userHandler struct {
//	userService services.UserService
//}
//
//func NewUserHandler(userService services.UserService) UserHandler {
//
//	return &userHandler{
//		userService: userService,
//	}
//}
//
//func (h *userHandler) FindUserProfile(ctx *gin.Context) {
//	// Get response
//	response, err := h.userService.FindUserProfile(ctx)
//	if err != nil {
//		api.ErrorWithMessage(ctx, err.Code, err.Message)
//		return
//	}
//
//	// Return
//	api.Ok(ctx, response)
//}
//
//func (h *userHandler) ChangePassword(ctx *gin.Context) {
//	req := models.ChangePasswordRequest{}
//	err := ctx.ShouldBindBodyWith(&req, binding.JSON)
//	if err != nil {
//		log.Errorf("Error binding request, err %v", err)
//		api.Error(ctx, http.StatusBadRequest)
//		return
//	}
//
//	errChange := h.userService.UpdatePassword(ctx, req)
//	if errChange != nil {
//		api.ErrorWithMessage(ctx, errChange.Code, errChange.Message)
//		return
//	}
//
//	// Return
//	api.Ok(ctx, nil)
//}
