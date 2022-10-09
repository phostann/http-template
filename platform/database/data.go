package database

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/phostann/template/http/pkg/configs"
	"github.com/phostann/template/http/pkg/log"
	"github.com/phostann/template/http/platform/database/ent"
)

type Data struct {
	DB     *ent.Client
	logger *log.Logger
}

func NewData(cfg *configs.Config, logger *log.Logger) (*Data, error) {
	if cfg.DataBase == nil {
		return nil, fmt.Errorf("missing database configuration")
	}
	client, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", cfg.DataBase.Username, cfg.DataBase.Password, cfg.DataBase.Host, cfg.DataBase.Port, cfg.DataBase.DB))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	if cfg.DataBase.Migrate {
		logger.Info("start database migration...")
		if err := client.Schema.Create(context.Background()); err != nil {
			return nil, fmt.Errorf("failed to migrate database: %v", err)
		}
		logger.Info("database migration completed successfully")
	}
	return &Data{DB: client, logger: logger}, nil
}

func (data *Data) Close() error {
	if data.DB != nil {
		data.logger.Info("closing database connection...")
		return data.DB.Close()
	}
	return nil
}
