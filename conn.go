package hivedriver

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"github.com/mattbaird/hive"
)

// Common error types
var (
	ErrSSLNotSupported     = errors.New("hive: SSL is not enabled on the server")
	ErrNotSupported        = errors.New("hive: Unsupported command")
	ErrInFailedTransaction = errors.New("hive: Could not complete operation in a failed transaction")
)

type drv struct{}

type conn struct {
	c *hive.HiveConnection
}

type stmt struct {
	cn      *conn
	name    string
	query   string
	cols    []string
	closed  bool
	lasterr error
}

type result struct {
}
type rows struct {
}

func (d *drv) Open(connectString string) (driver.Conn, error) {
	return Open(connectString)
}

func Open(connectString string) (_ driver.Conn, err error) {
	hive.MakePool("192.168.99.9:10000")
	hiveConnection, err := hive.GetHiveConn()
	cn := &conn{c: hiveConnection}
	return cn, err
}

func init() {
	sql.Register("hive", &drv{})
}

// connection implementation
func (cn *conn) Prepare(query string) (driver.Stmt, error) {
	return stmt{}, nil
}

func (cn *conn) Close() error {
	return nil
}

func (cn *conn) Begin() (driver.Tx, error) {
	return *cn, nil
}

func (cn conn) Commit() error {
	return nil
}
func (cn *conn) Rollback() error {
	return nil
}

// stmt impl

func (s *stmt) Close() error {
	return nil
}

func (s *stmt) NumInput() int {
	return 0
}

func (s *stmt) Exec(args []driver.Value) (driver.Result, error) {
	return result{}, nil
}

func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	return rows{}, nil
}

// result impl
func (r *result) LastInsertId() (int64, error) {
	return 0, nil
}

func (r *result) RowsAffected() (int64, error) {
	return 0, nil
}

// rows impl
func (rows *rows) Columns() []string {
	return []string{}
}
func (rows *rows) Close() error {
	return nil
}
func (rows *rows) Next(dest []driver.Value) error {
	return nil
}
