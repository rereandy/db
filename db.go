package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

type Connection struct {
	*dbr.Connection
}

func (c *Connection) NewSession() *dbr.Session {
	return c.Connection.NewSession(nil)
}

func Open(option *Options) *Connection {
	if option.Port == 0 {
		option.Port = 3306
	}

	if option.Host == "" {
		option.Host = "localhost"
	}

	if option.DataSource == "" {
		option.DataSource = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
			option.UserName,
			option.Password,
			fmt.Sprintf("%s:%d", option.Host, option.Port),
			option.DBName,
		)
		if option.ReadTimeout != 0 {
			option.DataSource += fmt.Sprintf("&readTimeout=%ds", option.ReadTimeout)
		}
		if option.WriteTimeout != 0 {
			option.DataSource += fmt.Sprintf("&writeTimeout=%ds", option.WriteTimeout)
		}
	}
	if option.Driver == "" {
		option.Driver = "mysql"
	}
	conn, err := dbr.Open(option.Driver, option.DataSource, NewEventReceiver(dbName(option.DataSource), 200, 200))
	if err != nil {
		panic(err)
	}
	if option.ConnMaxLifetime == 0 {
		option.ConnMaxLifetime = 600
	}
	conn.Dialect = &mysql{}
	conn.SetMaxIdleConns(option.MaxIdleConns)
	conn.SetMaxOpenConns(option.MaxOpenConns)
	conn.SetConnMaxLifetime(time.Duration(option.ConnMaxLifetime) * time.Second)
	return &Connection{Connection: conn}
}
