package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func doLog(level string, args ...interface{}) {
	info := buildInfoStr(args...)

	info = "- " + level + " - " + info
	log.Println(getTime(), info)

	logsDir := getLogsDir()
	logFilePath := filepath.Join(logsDir, "backend.log")

	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("FATAL: cannot open log file:", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println(info)
}

func Debug(args ...interface{}) {
	doLog("DEBUG", args...)
}

func Info(args ...interface{}) {
	doLog("INFO", args...)
}

func Warn(args ...interface{}) {
	doLog("WARN", args...)
}

func Error(args ...interface{}) {
	doLog("ERROR", args...)
}

func getLogsDir() string {
	logsDir := "./logs"
	if _, err := os.Stat(logsDir); os.IsNotExist(err) {
		err = os.Mkdir(logsDir, 0755)
		if err != nil {
			log.Fatalln("FATAL: cannot create logs directory:", err)
		}
	}
	return logsDir
}

func buildInfoStr(args ...interface{}) string {
	var info string
	for i, arg := range args {
		switch v := arg.(type) {
		case string:
			info += v
		default:
			info += fmt.Sprintf("%+v", v)
		}
		if i < len(args)-1 {
			info += " "
		}
	}
	return info
}

func getTime() string {
	now := time.Now()
	return now.Format("2006-01-02 15:04:05")
}
