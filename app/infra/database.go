package infra

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MAX_DB_CONN_LIFETIME = 60
	DB_PING_ATTEMPTS     = 18
	DB_PING_TIMEOUT_SECS = 10
)

// OpenConnectionMySQL open connection bd mysql
func OpenConnectionMySQL(url string) *gorm.DB {
	if len(url) > 0 {
		conn, err := gorm.Open(mysql.Open(url), &gorm.Config{})
		if err != nil {
			fmt.Println("[ERROR] Connection MySQL - ", err)
			time.Sleep(time.Second)
			os.Exit(101)
		}

		sqlDB, err := conn.DB()

		ctx, cancel := context.WithTimeout(context.Background(), DB_PING_TIMEOUT_SECS*time.Second)
		defer cancel()

		err = sqlDB.PingContext(ctx)
		if err != nil {
			fmt.Println("[ERROR] Connection MySQL ping fail - ", err)
		}

		sqlDB.SetConnMaxLifetime(time.Duration(MAX_DB_CONN_LIFETIME) * time.Minute)
		sqlDB.SetMaxIdleConns(20)
		sqlDB.SetMaxOpenConns(300)

		return conn
	}

	return nil
}

// OpenConnectionNoSQL open connection bd noSql
func OpenConnectionNoSQL(url string) (*mongo.Client, context.Context) {
	if len(url) > 0 {
		conn, err := mongo.NewClient(options.Client().ApplyURI(url))
		if err != nil {
			fmt.Println("[ERROR] Connection MySQL - ", err)
			time.Sleep(time.Second)
			os.Exit(101)
		}

		fmt.Println("[INFO] Check connection NOSQL")
		ctx, cancel := context.WithTimeout(context.Background(), DB_PING_TIMEOUT_SECS*time.Second)
		defer cancel()

		err = conn.Connect(ctx)
		if err != nil {
			fmt.Println("[ERROR] Connection MySQL ping fail - ", err)
		}

		return conn, ctx
	}

	return nil, nil
}
