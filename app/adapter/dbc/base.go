// Package dbc is (Database Connection) is used for create to represents low level database interfaces
// in order to have an unified way to access database handler
package dbc

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-redis/redis/v8"
)

// SqlDbc (SQL Database Connection) is a wrapper for SQL Database handler (can be *sql.DB or *sql.Tx)
type SqlDbc interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	// If you want support transactional
	Transactioner
}

type RDbc interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Del(ctx context.Context, key ...string) *redis.IntCmd
	GetDel(ctx context.Context, key string) *redis.StringCmd
}

type Transactioner interface {
	// Rollback a transaction
	Rollback() error
	// Commit a transaction
	Commit() error
	// TxEnd commits a transaction if no errors, otherwise callback
	// txFunc is the operations wrapped in a transaction
	TxEnd(txFunc func() error) error
}
