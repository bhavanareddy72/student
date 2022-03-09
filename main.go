package main

import (
	"github.com/bhavanareddy72/student/app"
	"github.com/bhavanareddy72/student/logger"
)

func main() {
	logger.Info("starting the student application")
	app.Start()
}
