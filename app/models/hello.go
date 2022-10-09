package models

type HelloQuery struct {
	Name string `query:"name" validate:"required"`
}
