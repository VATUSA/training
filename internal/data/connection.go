package data

import (
	"database/sql"
	"fmt"
	"github.com/vatusa/training/internal/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func dsn(c env.DatabaseEnvironment) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.User, c.Password, c.Host, c.Port, c.Database)
}

func Connect(c env.DatabaseEnvironment) error {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             3 * time.Second, // Slow SQL threshold
			LogLevel:                  logger.Warn,     // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,            // Don't include params in the SQL log
			Colorful:                  true,            // Disable color
		},
	)
	dsn := dsn(env.Environment.Database)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	DB, err = gorm.Open(mysql.New(mysql.Config{Conn: conn}), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		return err
	}
	mysqlDB, err := DB.DB()
	mysqlDB.SetMaxIdleConns(10)
	mysqlDB.SetMaxOpenConns(100)
	if err != nil {
		return err
	}
	return nil
}
