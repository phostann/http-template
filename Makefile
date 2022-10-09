.PHONY: api
api:
	swag init --outputTypes go,yaml -g  app/app.go

.PHONY: api-fmt

api-fmt:
	swag fmt
