package services

//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"main/component/models"
//	"main/component/repositories"
//	"net/http"
//
//	log "github.com/sirupsen/logrus"
//)
//
//type RoleService interface {
//	FindRolesByUserId(ctx context.Context, userId int64) ([]models.RoleResponse, *models.AIIeltsError)
//}
//
//type roleService struct {
//	roleRepo repositories.RoleRepo
//}
//
//func NewRoleService(roleRepo repositories.RoleRepo) RoleService {
//
//	return &roleService{
//		roleRepo: roleRepo,
//	}
//}
//
//func (s *roleService) FindRolesByUserId(ctx context.Context, userId int64) ([]models.RoleResponse, *models.AIIeltsError) {
//	var roles []models.RoleResponse
//	if userId == 0 {
//		return nil, &models.AIIeltsError{
//			IsError: true, Code: http.StatusBadRequest, Message: "Missing userId", Error: errors.New(fmt.Sprintf("Missing userId")),
//		}
//	}
//
//	rolesRes, errFind := s.roleRepo.FindRolesByUserId(ctx, userId)
//	if errFind != nil {
//		log.Errorf(fmt.Sprintf("Error find roles by userId %v, error %v", userId, errFind))
//		return nil, &models.AIIeltsError{
//			IsError: true, Code: http.StatusInternalServerError, Error: errFind,
//		}
//	}
//
//	// Mapping
//	if rolesRes != nil && len(rolesRes) > 0 {
//		for _, item := range rolesRes {
//			role := models.RoleResponse{}
//			role.RoleId = item.RoleId.ValueOrZero()
//			role.RoleName = item.RoleName.ValueOrZero()
//			role.RoleCode = item.RoleCode.ValueOrZero()
//			role.RoleDescription = item.RoleDescription.ValueOrZero()
//
//			roles = append(roles, role)
//		}
//	}
//
//	return roles, nil
//}
