package dbconn

import (
	"database/sql"
	"sync"
)

var (
	// _mysqlConnMgmt manages mysql connection by datasource name
	_mysqlConnMgmt = make(map[string]*sql.DB)
	_mysqlConnMu   sync.RWMutex
)

// GetMysqlConn return connection to mysql
func GetMysqlConn(dsn string) *sql.DB {
	_mysqlConnMu.RLock()
	conn, ok := _mysqlConnMgmt[dsn]
	// already connected and can ping to db => return immediatedly
	if ok && conn != nil && conn.Ping() == nil {
		return conn
	}
	_mysqlConnMu.RUnlock()

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if pingErr := conn.Ping(); pingErr != nil {
		panic(pingErr)
	}

	// store conn instance to storage manager
	_mysqlConnMu.Lock()
	_mysqlConnMgmt[dsn] = conn
	_mysqlConnMu.Unlock()

	return conn
}
