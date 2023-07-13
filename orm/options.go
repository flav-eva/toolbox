package orm

import (
	"gorm.io/gorm/logger"
)

type LocType string

const (
	Local LocType = "Local"
	UTC   LocType = "UTC"
)

type DBConfig struct {
	DBEnableTLS  bool   `json:"db_enable_tls" mapstructure:"db_enable_tls"`
	DBCaCertPEM  string `json:"db_ca_cert_pem" mapstructure:"db_ca_cert_pem"`
	CustomTLSKey string `json:"db_custom_tls_key" mapstructure:"db_custom_tls_key"`

	DBHost         string `json:"db_host" mapstructure:"db_host"`
	DBUser         string `json:"db_user" mapstructure:"db_user"`
	DBPassword     string `json:"db_password" mapstructure:"db_password"`
	DBName         string `json:"db_name" mapstructure:"db_name"`
	DBConnPoolSize int    `json:"db_conn_pool_size" mapstructure:"db_conn_pool_size"`

	// Dialect default is mysql
	Dialect string  `json:"dialect" mapstructure:"dialect"`
	Loc     LocType `json:"loc" mapstructure:"loc"`

	Debug bool `json:"debug" mapstructure:"debug"`
	// Logger 这里可以替换成我们自己的 logger 实现
	Logger logger.Interface
}
