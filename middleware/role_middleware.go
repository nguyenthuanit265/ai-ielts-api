package middleware

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"main/component/api"
//	"main/utils"
//	"net/http"
//)
//
//func (mid *ApiMiddleware) RequireRolesAccess(roleCodes ...string) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		currentUser, err := utils.GetCurrentUser(ctx)
//		if err.IsError {
//			api.ErrorWithMessage(ctx, http.StatusUnauthorized, "System cannot authorize user")
//			ctx.Abort()
//			return
//		}
//
//		if len(roleCodes) == 0 {
//			ctx.Next()
//			return
//		}
//
//		// Get roles by userId
//		roles, errFindRoles := mid.roleService.FindRolesByUserId(ctx, int64(currentUser.Id))
//		if errFindRoles != nil {
//			api.ErrorWithMessage(ctx, http.StatusInternalServerError, "Cannot get roles of user")
//			return
//		}
//		for _, role := range roles {
//			currentUser.Roles = append(currentUser.Roles, role.RoleCode)
//		}
//
//		if currentUser.Roles == nil || len(currentUser.Roles) == 0 {
//			api.ErrorWithMessage(ctx, http.StatusForbidden, fmt.Sprintf("User %v does not have any roles to access !!!", currentUser.Email))
//			ctx.Abort()
//			return
//		}
//
//		for _, role := range currentUser.Roles {
//			if utils.Contains(roleCodes, role) {
//				// Pass
//				ctx.Next()
//				return
//			}
//		}
//
//		api.ErrorWithMessage(ctx, http.StatusForbidden, fmt.Sprintf("User %v does not have role to access !!!", currentUser.Email))
//		ctx.Abort()
//	}
//}
