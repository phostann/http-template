package queries

import (
	"github.com/phostann/template/http/pkg/log"
	"github.com/phostann/template/http/platform/database"
)

type Query struct {
	data   *database.Data
	logger *log.Logger
}

// NewQuery returns a new Query
func NewQuery(data *database.Data, logger *log.Logger) *Query {
	return &Query{
		data:   data,
		logger: logger,
	}
}
