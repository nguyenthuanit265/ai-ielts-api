package services

//
//import (
//	"context"
//	"errors"
//	"main/component/models"
//	"main/component/repositories"
//	"main/utils"
//	"net/http"
//
//	log "github.com/sirupsen/logrus"
//)
//
//type RolePermissionService interface {
//	FindPermissionsByUserIdAndRoleCodes(ctx context.Context, userId int64, roleCode []string) ([]models.RolePermissionQueryResponse, *models.AIIeltsError)
//}
//
//type rolePermissionService struct {
//	rolePermissionRepo repositories.RolePermissionRepo
//}
//
//func NewRolePermissionService(rolePermissionRepo repositories.RolePermissionRepo) RolePermissionService {
//
//	return &rolePermissionService{
//		rolePermissionRepo: rolePermissionRepo,
//	}
//}
//
//func (s *rolePermissionService) FindPermissionsByUserIdAndRoleCodes(ctx context.Context, userId int64, roleCodes []string) ([]models.RolePermissionQueryResponse, *models.AIIeltsError) {
//	// Build filter
//	var filter models.RolePermissionFilter
//
//	if userId <= 0 {
//		return nil, &models.AIIeltsError{
//			IsError: true, Error: errors.New("UserId is empty or negative"), Code: http.StatusBadRequest, Message: "UserId is empty or negative",
//		}
//	}
//
//	filter.UserId = userId
//	filter.RoleCodes = roleCodes
//	response, errFind := s.rolePermissionRepo.FindPermissionsBy(ctx, filter)
//	if errFind != nil {
//		log.Errorf("Error get permissions with filter %v, error %v", utils.LogFull(filter), errFind)
//		return nil, &models.AIIeltsError{
//			IsError: true, Error: errFind, Code: http.StatusBadRequest, Message: "Error get permissions",
//		}
//	}
//
//	return response, nil
//}
