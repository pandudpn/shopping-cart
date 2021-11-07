package sql

import "database/sql"

type SqlDb struct {
	DB *sql.DB
}

type SqlTx struct {
	DB *sql.Tx
}

func (sd *SqlDb) Exec(query string, args ...interface{}) (sql.Result, error) {
	return sd.DB.Exec(query, args...)
}

func (sd *SqlDb) Prepare(query string) (*sql.Stmt, error) {
	return sd.DB.Prepare(query)
}

func (sd *SqlDb) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return sd.DB.Query(query, args...)
}

func (sd *SqlDb) QueryRow(query string, args ...interface{}) *sql.Row {
	return sd.DB.QueryRow(query, args...)
}

func (st *SqlTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return st.DB.Exec(query, args...)
}

func (st *SqlTx) Prepare(query string) (*sql.Stmt, error) {
	return st.DB.Prepare(query)
}

func (st *SqlTx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return st.DB.Query(query, args...)
}

func (st *SqlTx) QueryRow(query string, args ...interface{}) *sql.Row {
	return st.DB.QueryRow(query, args...)
}
