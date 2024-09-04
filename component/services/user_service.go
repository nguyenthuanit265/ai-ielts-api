package services

//
//import (
//	"context"
//	"golang.org/x/crypto/bcrypt"
//	"main/component/models"
//	"main/component/repositories"
//	"main/utils"
//	"net/http"
//	"strings"
//
//	log "github.com/sirupsen/logrus"
//)
//
//type UserService interface {
//	FindByEmail(ctx context.Context, email string) (models.UserResponse, *models.AIIeltsError)
//	FindUserProfile(ctx context.Context) (models.UserProfile, *models.AIIeltsError)
//	FindUserById(ctx context.Context, userId int64) (models.UserResponse, *models.AIIeltsError)
//	CheckPasswordHash(password string, hash string) bool
//	HashPassword(password string) (string, error)
//	UpdatePassword(ctx context.Context, req models.ChangePasswordRequest) *models.AIIeltsError
//}
//type userService struct {
//	userRepo              repositories.UserRepo
//	roleService           RoleService
//	rolePermissionService RolePermissionService
//}
//
//func NewUserService(
//	userRepo repositories.UserRepo,
//	roleService RoleService,
//	rolePermissionService RolePermissionService,
//) UserService {
//
//	return &userService{
//		userRepo:              userRepo,
//		roleService:           roleService,
//		rolePermissionService: rolePermissionService,
//	}
//}
//
//func (s *userService) FindByEmail(ctx context.Context, email string) (models.UserResponse, *models.AIIeltsError) {
//	user, errFindUser := s.userRepo.FindByEmail(ctx, email)
//	if errFindUser != nil {
//		log.Errorf("Error findByEmail email %v, error %v", email, utils.LogFull(errFindUser))
//		return models.UserResponse{}, &models.AIIeltsError{
//			IsError: true, Error: errFindUser, Code: http.StatusInternalServerError,
//		}
//	}
//
//	return user, nil
//}
//
//func (s *userService) FindUserById(ctx context.Context, userId int64) (models.UserResponse, *models.AIIeltsError) {
//	user, errFindUser := s.userRepo.FindUserById(ctx, userId)
//	if errFindUser != nil {
//		log.Errorf("Error FindUserById userId %v, error %v", userId, utils.LogFull(errFindUser))
//		return models.UserResponse{}, &models.AIIeltsError{
//			IsError: true, Error: errFindUser, Code: http.StatusInternalServerError,
//		}
//	}
//
//	return user, nil
//}
//
//func (s *userService) FindUserProfile(ctx context.Context) (models.UserProfile, *models.AIIeltsError) {
//	var response models.UserProfile
//	currentUser, _ := utils.GetCurrentUser(ctx)
//	// Find user by id
//	user, errFind := s.FindUserById(ctx, int64(currentUser.Id))
//	if errFind != nil {
//		log.Errorf("Error findByEmail %v", errFind)
//		return response, &models.AIIeltsError{
//			IsError: true, Code: http.StatusInternalServerError, Message: "Error login",
//		}
//	}
//
//	if user.UserId.IsZero() {
//		return response, &models.AIIeltsError{
//			IsError: true, Code: http.StatusNotFound, Message: "User is not found",
//		}
//	}
//
//	// Get roles
//	var roleCodes []string
//	roles, errFindRoles := s.roleService.FindRolesByUserId(ctx, user.UserId.ValueOrZero())
//	if errFindRoles != nil {
//		return response, &models.AIIeltsError{
//			IsError: true, Code: http.StatusInternalServerError, Message: "Cannot get roles of user",
//		}
//	}
//	for _, role := range roles {
//		roleCodes = append(roleCodes, role.RoleCode)
//	}
//
//	// Get permissions
//	var permissionCodes []string
//	permissions, errFindPermissions := s.rolePermissionService.FindPermissionsByUserIdAndRoleCodes(ctx, user.UserId.ValueOrZero(), roleCodes)
//	if errFindPermissions != nil {
//		return response, &models.AIIeltsError{
//			IsError: true, Code: http.StatusInternalServerError, Message: "Cannot get permissions of user",
//		}
//	}
//	mapExistPermission := make(map[string]bool)
//	for _, permission := range permissions {
//		if permission.PermissionCode.IsZero() || strings.TrimSpace(permission.PermissionCode.ValueOrZero()) == "" {
//			continue
//		}
//
//		if _, existed := mapExistPermission[permission.PermissionCode.ValueOrZero()]; !existed {
//			permissionCodes = append(permissionCodes, permission.PermissionCode.ValueOrZero())
//			mapExistPermission[permission.PermissionCode.ValueOrZero()] = true
//		}
//	}
//
//	response = models.UserProfile{
//		UserId:      int(user.UserId.ValueOrZero()),
//		FullName:    user.FullName.ValueOrZero(),
//		Email:       user.Email.ValueOrZero(),
//		Roles:       roleCodes,
//		Permissions: permissionCodes,
//	}
//
//	return response, nil
//}
//
//func (s *userService) HashPassword(password string) (string, error) {
//	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//	return string(bytes), err
//}
//
//func (s *userService) CheckPasswordHash(password string, hash string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//	return err == nil
//}
//
//func (s *userService) UpdatePassword(ctx context.Context, req models.ChangePasswordRequest) *models.AIIeltsError {
//	// Get currentUser
//	currentUser, errGetCurrentUser := utils.GetCurrentUser(ctx)
//	if errGetCurrentUser.IsError {
//		return &models.AIIeltsError{
//			IsError: true, Code: http.StatusInternalServerError, Message: "Error get current user",
//		}
//	}
//
//	// Validate new password and confirm password
//	if req.NewPassword != req.ConfirmPassword {
//		return &models.AIIeltsError{
//			IsError: true, Code: http.StatusBadRequest, Message: "New password and confirm password is not match",
//		}
//	}
//
//	// Validate old password
//	userEntity, errFind := s.FindUserById(ctx, int64(currentUser.Id))
//	if errFind != nil {
//		log.Errorf("Error find user by id %v", errFind)
//		return &models.AIIeltsError{
//			IsError: true, Code: http.StatusInternalServerError, Message: "Error change password",
//		}
//	}
//
//	if !s.CheckPasswordHash(req.OldPassword, userEntity.Password.ValueOrZero()) {
//		return &models.AIIeltsError{
//			IsError: true, Code: http.StatusBadRequest, Message: "Old password is not match",
//		}
//	}
//
//	// Hash new password
//	hashedPassword, errHash := s.HashPassword(req.NewPassword)
//	if errHash != nil {
//		log.Errorf("Error hash password %v", errHash)
//		return &models.AIIeltsError{
//			IsError: true, Code: http.StatusInternalServerError, Message: "Error change password",
//		}
//	}
//
//	// Update password
//	_, errUpdate := s.userRepo.UpdatePassword(ctx, hashedPassword, int64(currentUser.Id))
//	if errUpdate != nil {
//		log.Errorf("Error update password %v", errUpdate)
//		return &models.AIIeltsError{
//			IsError: true, Code: http.StatusInternalServerError, Message: "Error change password",
//		}
//	}
//
//	return nil
//}
