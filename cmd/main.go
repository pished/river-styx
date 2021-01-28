package main

import (
	"net/http"
	"time"

	"github.com/pished/river-styx/api"
	"github.com/pished/river-styx/router"
	lr "github.com/pished/river-styx/util/logger"
	vr "github.com/pished/river-styx/util/validator"
	"github.com/spf13/viper"
)

func main() {
	validator := vr.New()
	logger := lr.New(viper.GetBool("server.debug"))

	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Fatal().Err(err).Msg("Config file not found")
		} else {
			logger.Fatal().Err(err).Msg("Could not properly parse config file")
		}
	}

	app := api.New(logger, validator)
	apiRouter := router.New(app)

	logger.Info().Msgf("Starting server %v", viper.GetString("server.port"))

	s := &http.Server{
		Addr:         viper.GetString("server.port"),
		Handler:      apiRouter,
		ReadTimeout:  viper.GetDuration("server.timeout_read") * time.Second,
		WriteTimeout: viper.GetDuration("server.timeout_write") * time.Second,
		IdleTimeout:  viper.GetDuration("server.timeout_idle") * time.Second,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}
