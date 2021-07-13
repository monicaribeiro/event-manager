package main

import (
	"github.com/monicaribeiro/event-manager/app"
	"github.com/monicaribeiro/event-manager/logger"
)

func main() {
	logger.Info("Starting application...")
	app.Start()
}
