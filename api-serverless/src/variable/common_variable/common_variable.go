package common_variable

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	_ = godotenv.Load()
)

var (
	ENVIRONMENT = os.Getenv("ENVIRONMENT")
)
