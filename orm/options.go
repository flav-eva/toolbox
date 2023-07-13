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
	EnableTLS    bool   `json:"enable_tls"`
	CaCertPEM    string `json:"ca_cert_pem"`
	CustomTLSKey string `json:"custom_tls_key"`

	Host         string `json:"host"`
	Port         string `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DbName       string `json:"db_name"`
	ConnPoolSize int    `json:"conn_pool_size"`

	// Dialect default is mysql
	Dialect string  `json:"dialect"`
	Loc     LocType `json:"loc"`

	Debug  bool             `json:"debug"`
	Logger logger.Interface // 这里可以替换成我们自己的 logger 实现
}
