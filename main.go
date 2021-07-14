package main

import (
	"github.com/monicaribeiro/event-manager/app"
	"github.com/monicaribeiro/event-manager/logger"
)

func main() {
	logger.Info("V2 Starting application...")
	app.Start()
}
