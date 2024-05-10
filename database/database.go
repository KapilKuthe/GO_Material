package database

import (
	"context"
	"fmt"
	env "nfscGofiber/environment"

	"github.com/jackc/pgx/v4"
)

// Database defines the methods for interacting with a database
type Database interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

// PostgresDB is a PostgreSQL database implementation
type PostgresDB struct {
	conn *pgx.Conn
}

// NewPostgresDB creates a new instance of PostgresDB
func NewPostgresDB() (*PostgresDB, error) {
	conn, err := pgx.Connect(context.Background(), env.Conn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	return &PostgresDB{conn: conn}, nil
}

// Query executes a query on the database
func (p *PostgresDB) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return p.conn.Query(ctx, sql, args...)
}

// Close closes the database connection
func (p *PostgresDB) Close() {
	p.conn.Close(context.Background())
}

// func InitializeDB() (*gorm.DB, error) {
// 	db, err := gorm.Open(postgres.Open(env.Conn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Auto-migrate the schema
// 	err = db.AutoMigrate(&msg.MsgRequest{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }
