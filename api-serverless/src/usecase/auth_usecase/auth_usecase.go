package auth_usecase

import (
	"chat-platform-api.com/chat-platform-api/src/interface/repository_interface"
	"chat-platform-api.com/chat-platform-api/src/repository/user_repository"
	"chat-platform-api.com/chat-platform-api/src/tool/jwt_tool"
	"chat-platform-api.com/chat-platform-api/src/usecase/common_usecase"

	"chat-platform-api.com/chat-platform-api/src/type/dto/user_dto"
	"chat-platform-api.com/chat-platform-api/src/type/model/user_model"
	"chat-platform-api.com/chat-platform-api/src/type/response/user_response"
	"chat-platform-api.com/chat-platform-api/src/variable/auth_variable"
)

// 사용자 인증 Use Case
type AuthUseCase struct {
	*common_usecase.UseCase
	userRepository repository_interface.UserRepositoryImpl
}

var (
	usecase *AuthUseCase
)

func GetUseCase() *AuthUseCase {
	if usecase == nil {
		usecase = &AuthUseCase{
			userRepository: user_repository.GetUserRepository(),
		}
	}
	return usecase
}

// 사용자 토큰 생성
func (u *AuthUseCase) GenerateToken(user user_model.UserData, deviceId string) (*user_response.UserTokenResponse, error) {
	var err error
	defer func() {
		u.ErrorCheck(u.userRepository, err)
		if err != nil {
			u.userRepository.Rollback()
			return
		}
		u.userRepository.Commit()
	}()

	accessToken := jwt_tool.GenerateToken( // Access Token 생성
		user,
		auth_variable.JWT_SECRET_KEY,
		auth_variable.ACCESS_TOKEN_EXPIRATION,
	)

	refreshToken := jwt_tool.GenerateToken( // Refresh Token 생성
		user,
		auth_variable.JWT_SECRET_KEY,
		auth_variable.REFRESH_TOKEN_EXPIRATION,
	)

	if err = u.userRepository.SetRefreshToken(user.UserId, refreshToken, deviceId, user.IpAddr); err != nil { // Refresh Token 데이터 저장
		return nil, err
	}

	return &user_response.UserTokenResponse{
		UserInformation: user,
		AccessToken:     accessToken,
		RefreshToken:    refreshToken,
	}, nil
}

// 회원가입
func (u *AuthUseCase) SignUp(userData user_dto.SignUpDTO) (*user_response.UserTokenResponse, error) {
	// todo : 회원가입 구현
	return nil, nil
}

// 로그인
func (u *AuthUseCase) SignIn(signIn user_dto.SignInDTO) (*user_response.UserTokenResponse, error) {
	// todo : 로그인 구현
	return nil, nil
}
