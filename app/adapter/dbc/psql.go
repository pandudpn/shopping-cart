package dbc

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
	"github.com/spf13/viper"
)

func DatabaseConnection() *sql.DB {
	dsn := viper.GetString("db.dsn")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		logger.Log.Error(err)
		panic(err)
	}

	return db
}
