package log

import (
	"fmt"
	"time"
)

func Debug(log string) {
	fmt.Printf("[DEBUG][%s] %s\n", getTime(), log)
}
func Info(log string) {
	fmt.Printf("[INFO][%s] %s\n", getTime(), log)
}

func Warn(log string) {
	fmt.Printf("[WARN][%s] %s\n", getTime(), log)
}

func Error(log string) {
	fmt.Printf("[ERROR][%s] %s\n", getTime(), log)
}

func Fatal(log string) {
	fmt.Printf("[FATAL][%s] %s\n", getTime(), log)
}

func Panic(log string) {
	fmt.Printf("[PANIC][%s] %s\n", getTime(), log)
}

func Default(log string) {
	Info(log)
}

func getTime() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return currentTime
}
