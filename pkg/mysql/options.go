package mysql

// Option -.
type Option func(*Mysql)

// MaxIdleConns -.
func MaxIdleConns(size int) Option {
	return func(c *Mysql) {
		c.maxIdleConns = size
	}
}

// MaxOpenConns -.
func MaxOpenConns(size int) Option {
	return func(c *Mysql) {
		c.maxOpenConns = size
	}
}
