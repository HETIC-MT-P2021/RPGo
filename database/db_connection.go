package database

import (
	"database/sql"
	"fmt"
	"github.com/caarlos0/env/v6"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DBCon *sql.DB

type DBConfig struct {
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"MYSQL_DATABASE"`
	DbUser     string `env:"MYSQL_USER"`
	DbPassword string `env:"MYSQL_PASSWORD"`
}

func Connect() error {
	config := DBConfig{}
	if err := env.Parse(&config); err != nil {
		return fmt.Errorf("%+v", err)
	}

	dsn := config.DbUser + ":" + config.DbPassword + "@" + config.DbHost + "/" + config.DbName + "?parseTime=true&charset=utf8"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Errorf("could not open database: %v", err)
		return err
	}

	var dbErr error
	for i := 1; i <= 3; i++ {
		dbErr = db.Ping()
		if dbErr != nil {
			if i < 3 {
				log.Printf("database connection failed, %d retry : %v", i, dbErr)
				time.Sleep(10 * time.Second)
			}
			continue
		}

		break
	}

	if dbErr != nil {
		return fmt.Errorf("cannot connect to database after 3 attempts")
	}

	DBCon = db

	return nil
}
