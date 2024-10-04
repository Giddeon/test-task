package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"test/config"
)

const (
	defaultMaxPoolSize  = 1
	defaultConnAttempts = 10
	defaultConnTimeout  = time.Second
)

type Connection struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	Builder squirrel.StatementBuilderType
	DB      *sqlx.DB
}

func New(cnf *config.Conf) (*Connection, error) {
	pg := &Connection{
		maxPoolSize:  defaultMaxPoolSize,
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	dsn := fmt.Sprintf("host=%s port=%s database=%s user=%s password=%s sslmode=%s",
		cnf.PgHost,
		cnf.PgPort,
		cnf.PgDbName,
		cnf.PgUser,
		cnf.PgPassword,
		"disable",
	)

	var err error
	for pg.connAttempts > 0 {
		pg.DB, err = sqlx.Connect("postgres", dsn)
		if err == nil {
			break
		}

		zap.L().Info("Connection is trying to connect:", zap.Int("attempts left", pg.connAttempts))

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
	}

	return pg, nil
}

func (p *Connection) Close() {
	if p.DB != nil {
		_ = p.DB.Close()
	}
}
