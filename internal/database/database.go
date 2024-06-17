package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zcking/omega-catalog/internal/database/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB  *gorm.DB
	ddb *sql.DB
}

// NewDSN creates a new PostgreSQL DSN string for connecting to the database.
func NewDSN(host string, port int, username string, password string, properties map[string]string) string {
	databaseConnectionProperties := make([]string, len(properties))
	i := 0
	for k, v := range properties {
		databaseConnectionProperties[i] = fmt.Sprintf("%s=%s", k, v)
		i++
	}

	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s %s",
		host,
		port,
		username,
		password,
		strings.Join(databaseConnectionProperties, " "),
	)
}

// NewDatabase creates a new database interface.
func NewDatabase(dsn string) (*Database, error) {
	// https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("failed to connect to database", zap.Error(err))
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Error("failed to get database connection", zap.Error(err))
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	ddb := &Database{
		DB:  db,
		ddb: sqlDB,
	}
	return ddb, ddb.setup()
}

func (d *Database) setup() error {
	// Enable the UUID extension in PostgreSQL
	err := d.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		zap.L().Error("Failed to create UUID extension on database", zap.Error(err))
		return err
	}

	// Migrate the schemas
	d.DB.AutoMigrate(&models.Catalog{})
	return nil
}

// Close closes the database connection.
func (d *Database) Close() error {
	zap.L().Info("shutting down database connection...")
	return d.ddb.Close()
}

// Healthy checks if the database connection is healthy by pinging it.
func (d *Database) Healthy() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := d.ddb.PingContext(ctx); err != nil {
		zap.L().Error("failed ping to database", zap.Error(err))
		return false
	}
	return true
}
