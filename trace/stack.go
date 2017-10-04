package trace

import (
	"log"
	"runtime/debug"
)

// LogOnPanic 发生panic时记录日志
func LogOnPanic(logger *log.Logger) {
	// 捕获panic异常
	logger.Print("发生了致命错误")
	if err := recover(); err != nil {
		logger.Print("致命错误:", err)
	}
	debug.PrintStack()
}
