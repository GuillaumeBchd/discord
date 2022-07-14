package main

import (
	"github.com/guillaumebchd/discord"
	"github.com/guillaumebchd/discord/config"
	"go.uber.org/zap"
)

func main() {

	// Logger creation
	l, _ := zap.NewProduction()
	logger := l.Sugar()

	// Load configuration
	config, err := config.Load()
	if err != nil {
		logger.Fatal("Couldn't read configuration", err)
	}

	// Run the app
	d := discord.New(logger, config.Token, config.Guild)
	if err := d.Run(); err != nil {
		logger.Fatal(err)
	}
}
