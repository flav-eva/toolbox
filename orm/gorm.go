package orm

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"

	"github.com/flav-eva/toolbox/orm/mysqldialector"
)

var once sync.Once

/*
// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

db, err := gorm.Open(mysql.New(mysql.Config{
  DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
  DefaultStringSize: 256, // string 类型字段的默认长度
  DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
  DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
  DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
  SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
}), &gorm.Config{})

// postgres
dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// https://github.com/go-gorm/postgres
db, err := gorm.Open(postgres.New(postgres.Config{
  DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
  PreferSimpleProtocol: true, // disables implicit prepared statement usage
}), &gorm.Config{})
*/

func OpenDatabaseForMySQL(cfg *DBConfig) (*gorm.DB, error) {
	loc, err := time.LoadLocation(string(cfg.Loc))
	if err != nil {
		return nil, err
	}
	// dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBName, cfg.Loc)
	dsnCfg := &mysql.Config{
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPassword,
		Net:                  "tcp",
		Addr:                 cfg.DBHost,
		DBName:               cfg.DBName,
		Loc:                  loc,
		Params:               map[string]string{"charset": "utf8mb4"},
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	if cfg.DBEnableTLS {
		var onceErr error
		once.Do(func() {
			var pem []byte
			rootCertPool := x509.NewCertPool()
			pem, onceErr = os.ReadFile(cfg.DBCaCertPEM)
			if err != nil {
				return
			}
			if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
				onceErr = fmt.Errorf("failed append PEM")
				return
			}
			if onceErr = mysql.RegisterTLSConfig(cfg.CustomTLSKey, &tls.Config{
				RootCAs: rootCertPool,
			}); onceErr != nil {
				return
			}
			dsnCfg.TLSConfig = cfg.CustomTLSKey
		})
		if onceErr != nil {
			return nil, onceErr
		}
	}

	if cfg.Debug && cfg.Logger != nil {
		cfg.Logger.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysqldialector.Open(dsnCfg.FormatDSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: cfg.Logger,
	})
	if err != nil {
		return nil, fmt.Errorf("gorm open database: %w", err)
	}

	if err := db.Use(dbresolver.Register(dbresolver.Config{}).
		SetConnMaxLifetime(60 * time.Second).
		SetMaxIdleConns(cfg.DBConnPoolSize).
		SetMaxOpenConns(cfg.DBConnPoolSize)); err != nil {
		return nil, fmt.Errorf("gorm initialize dbresolver: %w", err)
	}

	// TODO add opentracing:span opentracing integration with GORM
	// 参考：https://github.com/go-gorm/opentracing/tree/main

	return db, nil
}

//
//  ******************* For Postgres **********************
//

// OpenDatabaseForPostgres postgres
func OpenDatabaseForPostgres(cfg *DBConfig) (*gorm.DB, error) {

	return nil, nil
}
