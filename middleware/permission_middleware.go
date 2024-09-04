package middleware

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/gin-gonic/gin/binding"
//	log "github.com/sirupsen/logrus"
//	"main/component/api"
//	"main/component/models"
//	"main/utils"
//	"net/http"
//	"strconv"
//)
//
//func getOrgIdFromContext(ctx *gin.Context) (int, *models.AIIeltsError) {
//	req := models.IMSRequest{}
//	var orgId int
//
//	// Binding param
//	queryValues := ctx.Request.URL.Query()
//	orgId, _ = strconv.Atoi(queryValues.Get("orgId"))
//	if orgId != 0 {
//		return orgId, nil
//	}
//
//	errBindBody := ctx.ShouldBindBodyWith(&req, binding.JSON)
//	if errBindBody != nil {
//		log.Errorf("Error binding request, err %v", errBindBody)
//		return 0, &models.AIIeltsError{
//			IsError: true, Code: http.StatusBadRequest, Error: errBindBody,
//		}
//	}
//
//	orgId = req.OrgId
//	return orgId, nil
//}
//func (mid *ApiMiddleware) RequirePermissionsAccess(permissionCodes ...string) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		if len(permissionCodes) == 0 {
//			ctx.Next()
//			return
//		}
//
//		currentUser, _ := utils.GetCurrentUser(ctx)
//		// Get orgId from context
//		orgId, errGetOrgId := getOrgIdFromContext(ctx)
//		if errGetOrgId != nil {
//			log.Errorf("Error get orgId from context, error %v", errGetOrgId)
//			api.Error(ctx, http.StatusBadRequest)
//			ctx.Abort()
//			return
//		}
//
//		// If orgId existed -> check permission by geo
//		if orgId != 0 {
//			var orgIdsByUser []int
//			geosResponse, errGetOrgIdByUser := mid.geoService.FindGeosByCurrentUser(ctx, utils.ModeGeo)
//			if errGetOrgIdByUser != nil {
//				log.Errorf("Error get orgId by user %v, error %v", utils.LogFull(currentUser), errGetOrgId)
//				api.Error(ctx, http.StatusBadRequest)
//				ctx.Abort()
//				return
//			}
//
//			mapExist := make(map[int]bool)
//			for _, geo := range geosResponse {
//				if _, existed := mapExist[geo.OrgId]; !existed {
//					orgIdsByUser = append(orgIdsByUser, geo.OrgId)
//					mapExist[geo.OrgId] = true
//				}
//			}
//
//			if !utils.Contains(orgIdsByUser, orgId) {
//				api.ErrorWithMessage(ctx, http.StatusForbidden, fmt.Sprintf("User %v need to be configured to access this geo", currentUser.Email))
//				ctx.Abort()
//				return
//			}
//		}
//
//		// Get permission codes by role code and check match with parameter permissionCodes
//		permissionResponse, errFindPermission := mid.rolePermissionService.FindPermissionsByUserIdAndRoleCodes(ctx, int64(currentUser.Id), nil)
//		if errFindPermission != nil {
//			log.Errorf("Error get permission of user %v, error %v", utils.LogFull(currentUser), errFindPermission)
//			api.Error(ctx, http.StatusInternalServerError)
//			ctx.Abort()
//			return
//		}
//
//		mapRoleWithPermission := make(map[string][]string)
//		mapExist := make(map[string]bool)
//		var permissions []string
//		for _, item := range permissionResponse {
//			if !item.RoleCode.IsZero() && item.RoleCode.ValueOrZero() != "" {
//				mapRoleWithPermission[item.RoleCode.ValueOrZero()] = append(mapRoleWithPermission[item.RoleCode.ValueOrZero()], item.PermissionCode.ValueOrZero())
//			}
//
//			if _, existed := mapExist[item.PermissionCode.ValueOrZero()]; !existed {
//				permissions = append(permissions, item.PermissionCode.ValueOrZero())
//				mapExist[item.PermissionCode.ValueOrZero()] = true
//			}
//		}
//
//		for _, code := range permissionCodes {
//			if utils.Contains(permissions, code) {
//				ctx.Next()
//				return
//			}
//		}
//
//		api.ErrorWithMessage(ctx, http.StatusForbidden, fmt.Sprintf("User %v does not have permission to access !!!", currentUser.Email))
//		ctx.Abort()
//	}
//}
