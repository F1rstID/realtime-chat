// infrastructure/logger/logger.go
package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	requestLogger *log.Logger
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[37m"
)

// getLogPrefix returns colored prefix with timestamp
func getLogPrefix(level string) string {
	var color string
	switch level {
	case "INFO":
		color = colorGreen
	case "ERROR":
		color = colorRed
	case "REQUEST":
		color = colorBlue
	default:
		color = colorReset
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s%s [%s]%s ", colorGray, timestamp, level, color)
}

func init() {
	// 로그 파일 디렉토리 생성
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal("Failed to create log directory:", err)
	}

	// 현재 날짜로 로그 파일 생성
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")

	// 일반 로그 파일
	infoLogFile, err := os.OpenFile(
		filepath.Join(logDir, fmt.Sprintf("app_%s.log", date)),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatal("Failed to open info log file:", err)
	}

	// 에러 로그 파일
	errorLogFile, err := os.OpenFile(
		filepath.Join(logDir, fmt.Sprintf("error_%s.log", date)),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatal("Failed to open error log file:", err)
	}

	// HTTP 요청 로그 파일
	requestLogFile, err := os.OpenFile(
		filepath.Join(logDir, fmt.Sprintf("request_%s.log", date)),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatal("Failed to open request log file:", err)
	}

	// MultiWriter를 사용하여 파일과 콘솔에 동시에 출력
	infoWriter := io.MultiWriter(os.Stdout, infoLogFile)
	errorWriter := io.MultiWriter(os.Stdout, errorLogFile)
	requestWriter := io.MultiWriter(os.Stdout, requestLogFile)

	// 로거 초기화
	infoLogger = log.New(infoWriter, "", 0)
	errorLogger = log.New(errorWriter, "", 0)
	requestLogger = log.New(requestWriter, "", 0)
}

// formatLog formats the log message with file info and color
func formatLog(level, file string, line int, msg string) string {
	prefix := getLogPrefix(level)
	return fmt.Sprintf("%s%s:%d: %s%s\n",
		prefix, filepath.Base(file), line, msg, colorReset)
}

// Info logs general information
func Info(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf(format, v...)
	formattedMsg := formatLog("INFO", file, line, msg)
	infoLogger.Print(formattedMsg)
}

// Error logs error information
func Error(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf(format, v...)
	formattedMsg := formatLog("ERROR", file, line, msg)
	errorLogger.Print(formattedMsg)
}

// LogRequest logs HTTP request information
func LogRequest(method, path, ip, userAgent string, statusCode int, duration time.Duration) {
	_, file, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf(
		"Method: %s, Path: %s, IP: %s, Status: %d, Duration: %v",
		method, path, ip, statusCode, duration,
	)
	formattedMsg := formatLog("REQUEST", file, line, msg)
	requestLogger.Print(formattedMsg)
}

// LogError logs error with request context
func LogError(err error, method, path string) {
	_, file, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("Method: %s, Path: %s, Error: %v", method, path, err)
	formattedMsg := formatLog("ERROR", file, line, msg)
	errorLogger.Print(formattedMsg)
}

// Debug logs debug information (only in development)
func Debug(format string, v ...interface{}) {
	if os.Getenv("GO_ENV") != "production" {
		_, file, line, _ := runtime.Caller(1)
		msg := fmt.Sprintf(format, v...)
		formattedMsg := formatLog("DEBUG", file, line, msg)
		infoLogger.Print(formattedMsg)
	}
}
