package global

import "gorm.io/gorm"

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

var (
	DBEngine *gorm.DB
	DBConfig *DatabaseSettingS
)
