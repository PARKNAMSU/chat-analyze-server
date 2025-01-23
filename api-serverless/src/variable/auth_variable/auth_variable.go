package auth_variable

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type PermissionInformation struct {
	Type         string
	Permission   int
	MaxChatNum   int
	MaxUserNum   int
	MaxAccessNum int
}

var (
	_ = godotenv.Load()
)

var (
	JWT_SECRET_KEY     = os.Getenv("JWT_SECRET_KEY")
	ENCRYPT_SECRET_KEY = os.Getenv("ENCRYPT_SECRET_KEY")

	ACCESS_TOKEN_EXPIRATION  = time.Hour * 2
	REFRESH_TOKEN_EXPIRATION = time.Hour * 24 * 30

	Permissions = map[string]int{
		"all":          1 << 0,
		"oneToOneChat": 1 << 1,
		"groupChat":    1 << 2,
		"fileUpload":   1 << 3,
		"secretChat":   1 << 4,
	}

	PREMIUM_PERMISSION = PermissionInformation{
		Type:         "PREMIUM",
		Permission:   Permissions["all"],
		MaxChatNum:   5000,
		MaxUserNum:   20000,
		MaxAccessNum: 1000,
	}

	NORMAL_PERMISSION = PermissionInformation{
		Type: "NORMAL",
		Permission: Permissions["oneToOneChat"] +
			Permissions["groupChat"],
		MaxChatNum:   1000,
		MaxUserNum:   5000,
		MaxAccessNum: 100,
	}
)
