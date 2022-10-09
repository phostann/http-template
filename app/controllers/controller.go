package controllers

import (
	"github.com/phostann/template/http/app/queries"
	"github.com/phostann/template/http/pkg/log"
)

type Controller struct {
	query  *queries.Query
	logger *log.Logger
}

func NewController(query *queries.Query, logger *log.Logger) *Controller {
	return &Controller{
		query:  query,
		logger: logger,
	}
}
