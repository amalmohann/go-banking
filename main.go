package main

import (
	"github.com/amalmohann/banking/app"
	"github.com/amalmohann/banking/logger"
)

func main() {
	logger.Info("Starting the application!")
	app.Start()
}
