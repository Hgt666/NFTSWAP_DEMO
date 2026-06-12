package internal

import (
	"fmt"
	"time"
	"go.uber.org/zap"
	
)

// SimpleAlert 简易告警（可扩展钉钉/企业微信机器人）
func SimpleAlert(title, content string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	Logger.Error("【系统告警】"+title,
		zap.String("time", now),
		zap.String("content", content),
	)
	fmt.Printf("\n===== ALERT =====\n%s | %s | %s\n", now, title, content)
}