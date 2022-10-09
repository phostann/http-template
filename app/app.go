package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/phostann/template/http/app/controllers"
	"github.com/phostann/template/http/app/queries"
	"github.com/phostann/template/http/pkg/configs"
	"github.com/phostann/template/http/pkg/log"
	"github.com/phostann/template/http/pkg/routes"
	"github.com/phostann/template/http/platform/database"
	"github.com/spf13/cobra"
)

type App struct {
	app    *echo.Echo
	cfg    *configs.Config
	logger *log.Logger
	cmd    *cobra.Command
	data   *database.Data
	query  *queries.Query
	ctrl   *controllers.Controller
}

type Option func(*App)

func NewApp() *App {
	logger := log.NewLogger()
	cfg, err := configs.ReadConfig()
	if err != nil {
		logger.Panicf("failed to read config: %v", err)
	}

	app := &App{}

	e := echo.New()
	// 中间件使用
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	cmd := &cobra.Command{
		Use:   "http",
		Short: "An http server template based on go fiber",
		Long:  "An http server template based on go fiber",
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := fmt.Sprintf("%s:%d", cfg.Server.Address, cfg.Server.Port)
			logger.Infof("server listening at %s", addr)
			return e.Start(addr)
		},
	}

	app.app = e
	app.cfg = cfg
	app.logger = logger
	app.cmd = cmd

	return app
}

// 数据库连接
func withDatabase() Option {
	return func(a *App) {
		data, err := database.NewData(a.cfg, a.logger)
		if err != nil {
			a.logger.Panicf("failed to run with database option: %v", err)
		}
		a.data = data
	}
}

// 数据库查询
func withQueries() Option {
	return func(a *App) {
		query := queries.NewQuery(a.data, a.logger)
		a.query = query
	}
}

// 路由对应 handler
func withControllers() Option {
	return func(a *App) {
		ctrl := controllers.NewController(a.query, a.logger)
		a.ctrl = ctrl
	}
}

// 路由配置
func withRoutes() Option {
	return func(a *App) {
		route := routes.NewRoute(a.ctrl, a.logger)
		v1 := a.app.Group("/api/v1")
		route.SetUp(v1)
	}
}

func (a *App) run(opts ...Option) error {
	for _, opt := range opts {
		opt(a)
	}

	return a.cmd.Execute()

}

// Run godoc
// @title          Swagger Example API
// @version        1.0
// @description    This is a sample server celler server.
// @termsOfService http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes        http
// @host           localhost:8080
// @BasePath       /api/v1
func (a *App) Run() error {
	return a.run(
		// withDatabase(),
		withQueries(),
		withControllers(),
		withRoutes(),
	)
}
