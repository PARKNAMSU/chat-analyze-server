package user_response

import "chat-platform-api.com/chat-platform-api/src/type/model/user_model"

type UserTokenResponse struct {
	UserInformation user_model.UserData `json:"userInformation"`
	AccessToken     string              `json:"accessToken"`
	RefreshToken    string              `json:"refreshToken"`
}
