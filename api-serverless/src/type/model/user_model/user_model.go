package user_model

type UserData struct {
	UserId         int     `json:"userId"`
	Status         int     `json:"status"`
	IpAddr         string  `json:"ipAddr"`
	Email          *string `json:"email"`
	OauthId        *string `json:"oauthId"`
	OauthHost      *string `json:"oauthHost"`
	Name           *string `json:"name"`
	Authentication int     `json:"authentication"`
	AuthType       *string `json:"authType"`
}

type SetUserInformation struct {
	UserId   int     `json:"userId"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Name     *string `json:"name"`
}
