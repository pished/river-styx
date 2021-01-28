package api

import (
	"github.com/go-playground/validator"
	"github.com/pished/river-styx/util/logger"
)

const (
	appErrDataCreationFailure    = "data creation failure"
	appErrDataAccessFailure      = "data access failure"
	appErrDataUpdateFailure      = "data update failure"
	appErrJsonCreationFailure    = "json creation failure"
	appErrFormDecodingFailure    = "form decoding failure"
	appErrFormErrResponseFailure = "form error response failure"
)

type Api struct {
	logger    *logger.Logger
	validator *validator.Validate
}

func New(logger *logger.Logger, validator *validator.Validate) *Api {
	return &Api{
		logger:    logger,
		validator: validator,
	}
}

func (a *Api) Logger() *logger.Logger {
	return a.logger
}
