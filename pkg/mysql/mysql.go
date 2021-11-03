package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

const (
	_defaultMaxIdleConns = 10
	_defaultMaxOpenConns = 20
)

type Mysql struct {
	maxIdleConns int
	maxOpenConns int
	Engine       *xorm.Engine
}

// user:password@(addr)/dbname?charset=utf8&parseTime=True&loc=Local
func New(url string, opts ...Option) (*Mysql, error) {
	mysql := &Mysql{
		maxIdleConns: _defaultMaxIdleConns,
		maxOpenConns: _defaultMaxOpenConns,
	}

	// Custom options
	for _, opt := range opts {
		opt(mysql)
	}

	var err error
	mysql.Engine, err = xorm.NewEngine("mysql", url)
	if err != nil {
		return nil, fmt.Errorf("mysql - NewMySQL - NewEngine: %w", err)
	}
	mysql.Engine.SetMaxIdleConns(mysql.maxIdleConns)
	mysql.Engine.SetMaxOpenConns(mysql.maxOpenConns)

	if err = mysql.Engine.DB().Ping(); err != nil {
		return nil, fmt.Errorf("mysql - NewMySQL - Ping == 0: %w", err)
	}

	return mysql, nil
}

// Close -.
func (m *Mysql) Close() {
	m.Engine.Close()
}
