package utils

import (
	"log"
	"runtime"
	"strings"
	"time"
)

func TimeCost(fucName string, start time.Time) {
	tc := time.Since(start)
	log.Printf("%s cost = %v\n", fucName, tc)
}

/**
* GetCurrentFuncName
* runtime.Caller(1) 可以返回当前函数的调用栈信息，
* runtime.FuncForPC(pc) 可以返回对应的函数名，
* 然后通过 strings.TrimPrefix 去掉前缀 main.
 */
func GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return strings.TrimPrefix(runtime.FuncForPC(pc).Name(), "main.")
}
