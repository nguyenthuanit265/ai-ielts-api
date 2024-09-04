package services

//
//import (
//	"context"
//	log "github.com/sirupsen/logrus"
//	"main/component/models"
//	"main/component/repositories"
//	"net/http"
//)
//
//type AuthService interface {
//	Login(ctx context.Context, req models.AuthRequest) (models.AuthResponse, *models.AIIeltsError)
//}
//type authService struct {
//	authRepo              repositories.AuthRepo
//	userService           UserService
//	roleService           RoleService
//	rolePermissionService RolePermissionService
//}
//
//func NewAuthService(authRepo repositories.AuthRepo,
//	userService UserService,
//	roleService RoleService,
//	rolePermissionService RolePermissionService,
//) AuthService {
//
//	return &authService{
//		authRepo:              authRepo,
//		userService:           userService,
//		roleService:           roleService,
//		rolePermissionService: rolePermissionService,
//	}
//}
//
//func (s *authService) Login(ctx context.Context, req models.AuthRequest) (models.AuthResponse, *models.AIIeltsError) {
//	var response models.AuthResponse
//
//	// Find user by email
//	user, errFind := s.userService.FindByEmail(ctx, req.Email)
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
//	// Validate password
//	if !s.userService.CheckPasswordHash(req.Password, user.Password.ValueOrZero()) {
//		return response, &models.AIIeltsError{
//			IsError: true, Code: http.StatusBadRequest, Message: "Password is not match",
//		}
//	}
//
//	//if user.Password.ValueOrZero() != req.Password {
//	//	return response, &models.AIIeltsError{
//	//		IsError: true, Code: http.StatusBadRequest, Message: "Password is not match",
//	//	}
//	//}
//
//	// Generate token
//	accessToken, refreshToken, errGen := GenerateToken(models.AuthClaim{
//		Email:      user.Email.ValueOrZero(),
//		FullName:   user.FullName.ValueOrZero(),
//		Id:         int(user.UserId.ValueOrZero()),
//		Authorized: true,
//	})
//
//	if errGen != nil {
//		log.Errorf("Error generate token %v", errGen)
//		return response, &models.AIIeltsError{
//			IsError: true, Code: http.StatusInternalServerError, Message: "Error generate token",
//		}
//	}
//
//	response = models.AuthResponse{
//		AccessToken:  accessToken,
//		RefreshToken: refreshToken,
//	}
//
//	return response, nil
//}
