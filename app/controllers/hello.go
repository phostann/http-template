package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"

	"github.com/phostann/template/http/app/models"
	"github.com/phostann/template/http/pkg/response"
	"github.com/phostann/template/http/pkg/utils"
)

// Hello 		godoc
// @Summary     Hello
// @Description Hello World
// @Accept      json
// @Produce     json
// @Param       name query    string false "your name"
// @Success     200  {object} response.Response
// @Router      /hello  [get]
func (c *Controller) Hello(ctx echo.Context) error {
	req := &models.HelloQuery{}
	err := ctx.Bind(req)
	if err != nil {
		return response.ErrMsg(ctx, err.Error())
	}
	errs := utils.ValidateStruct(req)
	if len(errs) > 0 {
		return response.BadRequest(ctx, errs)
	}
	return response.Msg(ctx, fmt.Sprintf("hello %s", req.Name))
}
