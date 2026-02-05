package config

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

func NewDBLogger() logger.Interface {
	logFile, err := os.OpenFile(
		"logs/db.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)
	if err != nil {
		log.Fatal("❌ cannot open db log file")
	}

	return logger.New(
		log.New(logFile, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // slow query warning
			LogLevel:      logger.Info, // logs all SQL
			Colorful:      false,       // file only
		},
	)
}
