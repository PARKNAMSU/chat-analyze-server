package main

import (
	"fmt"
	"time"

	"chat-platform-api.com/chat-platform-api/src/tool/jwt_tool"
)

func main() {
	data := jwt_tool.GenerateToken[string]("test","aaaabbbb",time.Minute * 30)
	fmt.Println(data)
}
