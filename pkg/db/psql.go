package db

import (
	"context"
	"fmt"

	"os"

	"github.com/jackc/pgx/v5"
	"github.com/temur-shamshidinov/task_app/config"
)

func ConnectToDb(pgCfg config.PgConfig) (*pgx.Conn, error) {
	var ctx = context.Background()

	url := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s",
		pgCfg.Username,
		pgCfg.Password,
		pgCfg.Host,
		pgCfg.Port,
		pgCfg.DatabaseName,
	)
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v \n", err)
		return nil, err
	}
	return conn, nil
}
