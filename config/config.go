package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
)

type Conf struct {
	TcpPort  string `env:"TCP_PORT"`
	HttpPort string `env:"HTTP_PORT"`

	PgHost     string `env:"PG_HOST"`
	PgPort     string `env:"PG_PORT"`
	PgUser     string `env:"PG_USER"`
	PgPassword string `env:"PG_PWD"`
	PgDbName   string `env:"PG_DB_NAME"`
}

var Cnf Conf

func NewConf() error {
	if err := env.Parse(&Cnf, env.Options{RequiredIfNoDef: true}); err != nil {
		zap.L().Fatal("error on parse env config:", zap.Error(err))
		return err
	}

	flag.StringVar(&Cnf.PgHost, "pgHost", Cnf.PgHost, "PGSQL host")
	flag.StringVar(&Cnf.PgPort, "pgPort", Cnf.PgPort, "PGSQL port")
	flag.StringVar(&Cnf.PgUser, "pgUser", Cnf.PgUser, "PGSQL user")
	flag.StringVar(&Cnf.PgPassword, "pgPass", Cnf.PgPassword, "PGSQL password")
	flag.StringVar(&Cnf.PgDbName, "pgDbName", Cnf.PgDbName, "PGSQL db name")

	flag.Parse()

	return nil
}
