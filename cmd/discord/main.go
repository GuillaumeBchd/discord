package main

import (
	"github.com/guillaumebchd/discord"
	"github.com/guillaumebchd/discord/config"
	"go.uber.org/zap"
)

func main() {

	// Logger creation
	logger, _ := zap.NewProduction()

	// Load configuration
	config, err := config.Load()
	if err != nil {
		logger.Fatal("Couldn't read configuration", zap.Error(err))
	}

	// Run the app
	d := discord.New(config.Token, discord.WithGuild(config.Guild), discord.WithLogger(logger))
	if err := d.Run(); err != nil {
		logger.Fatal("Error while running", zap.Error(err))
	}
}
