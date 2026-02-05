package config

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
)

func InitFileLogger() {
	
	// open log file
	file, err := os.OpenFile(
		"logs/app.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}

	// redirect ALL logs to file
	log.SetOutput(file)

	// optional: add timestamp + file line
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// redirect fmt.Println also
	os.Stdout = file
	os.Stderr = file
	// 🔥 REDIRECT GIN LOGS
	gin.DefaultWriter = file
	gin.DefaultErrorWriter = file

}
