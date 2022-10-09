package response

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Page struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}

type Response struct {
	Msg  string      `json:"msg" validate:"required"`
	Data interface{} `json:"data,omitempty" validate:"optional" swaggertype:"object"`
	Page *Page       `json:"page,omitempty" validate:"optional"`
}

func Msg(ctx echo.Context, msg string) error {
	return ctx.JSON(http.StatusOK, Response{
		Msg: msg,
	})
}

func Data(data interface{}) Response {
	return Response{
		Msg:  "ok",
		Data: data,
	}
}

func PageData(data interface{}, page, pageSize, total int) Response {
	return Response{
		Msg:  "ok",
		Data: data,
		Page: &Page{
			Page:     page,
			PageSize: pageSize,
			Total:    total,
		},
	}
}

func ErrMsg(ctx echo.Context, errMsg string) error {
	return ctx.JSON(http.StatusInternalServerError, Response{
		Msg: errMsg,
	})
}

func BadRequest(ctx echo.Context, errs validator.ValidationErrorsTranslations) error {
	return ctx.JSON(http.StatusBadRequest, Response{Msg: "非法的参数", Data: errs})
}
