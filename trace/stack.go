package trace

import (
	"log"
	"runtime/debug"
)

// LogStack 将栈输出到日志中
func LogStack(logger *log.Logger) {
	// 捕获panic异常
	logger.Print("发生了致命错误")
	if err := recover(); err != nil {
		logger.Print("致命错误:", err)
	}

	logger.Print(string(debug.Stack()))
}
