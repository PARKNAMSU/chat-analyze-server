package user_model

type UserData struct {
	UserId         int     `json:"userId"`
	Status         int     `json:"status"`
	IpAddr         int     `json:"ipAddr"`
	Email          *string `json:"email"`
	OauthId        *string `json:"oauthId"`
	OauthHost      *string `json:"oauthHost"`
	Name           *string `json:"name"`
	Authentication int     `json:"authentication"`
	AuthType       *string `json:"authType"`
}
