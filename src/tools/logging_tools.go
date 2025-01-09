package tools

import (
	"fmt"
	"log"
)

var (
	errTag  = "[Error]"
	infoTag = "[Info]"
)

func PanicError(funcName string, message string) {
	log.Panicf("%s[%s]:[%s]", errTag, funcName, message)
}

func PrintErrorLog(funcName string, message string) {
	fmt.Printf("%s[%s]:[%s]", errTag, funcName, message)
}

func PrintInfoLog(funcName string, message string) {
	fmt.Printf("%s[%s]:[%s]", infoTag, funcName, message)
}
