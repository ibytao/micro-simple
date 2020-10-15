package config

import "time"

type MysqlConfig interface {
	GetURL() string
	GetEnabled() bool
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
	GetConnMaxLifetime() time.Duration
}

type defaultMysqlConfig struct {
	URL               string        `json:"url"`
	Enable            bool          `json:"enabled"`
	MaxIdleConnection int           `json:"maxIdleConnection"`
	MaxOpenConnection int           `json:maxOpenConnection"`
	ConnMaxLifetime   time.Duration `json:"connMaxLifetime`
}

func (m defaultMysqlConfig) GetURL() string {
	return m.URL
}

func (m defaultMysqlConfig) GetEnabled() bool {
	return m.Enable
}

func (m defaultMysqlConfig) GetMaxIdleConnection() int {
	return m.MaxIdleConnection
}

func (m defaultMysqlConfig) GetMaxOpenConnection() int {
	return m.MaxOpenConnection
}

func (m defaultMysqlConfig) GetConnMaxLifetime() time.Duration {
	return m.ConnMaxLifetime
}
