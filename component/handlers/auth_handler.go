package handlers

//
//import (
//	"main/component/api"
//	"main/component/models"
//	"main/component/services"
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//	"github.com/gin-gonic/gin/binding"
//	log "github.com/sirupsen/logrus"
//)
//
//type AuthHandler interface {
//	Login(ctx *gin.Context)
//}
//
//type authHandler struct {
//	authService services.AuthService
//}
//
//func NewAuthHandler(authService services.AuthService) AuthHandler {
//
//	return &authHandler{
//		authService: authService,
//	}
//}
//
//// Login godoc
////	@Summary		Auth
////	@CategoryDescription	auth
////	@Tags			Auth
////	@Accept			json
////	@Produce		json
////	@Success		200	{object}	models.AuthResponse
////	@Security		Bearer
////	@Router			/api/v1/login 	[post]
//func (h *authHandler) Login(ctx *gin.Context) {
//	req := models.AuthRequest{}
//	err := ctx.ShouldBindBodyWith(&req, binding.JSON)
//	if err != nil {
//		log.Errorf("Error binding request, err %v", err)
//		api.Error(ctx, http.StatusBadRequest)
//		return
//	}
//
//	response, errInsert := h.authService.Login(ctx, req)
//	if errInsert != nil {
//		api.ErrorWithMessage(ctx, errInsert.Code, errInsert.Message)
//		return
//	}
//
//	// Return
//	api.Ok(ctx, response)
//}
