package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils"
	"hello/config"
	"hello/model"
	"time"
)

var db *gorm.DB

func ConnectDB() {
	database_url := config.DB.DATABASE_URL
	var err error
	logger := dbLogger{}

	mysqlConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.DB.PREFIX, // table name prefix, table for `User` would be `t_users`
			SingularTable: true,             // use singular table name, table for `User` would be `user` with this option enabled
		},
		Logger: &logger,
	}
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       database_url, // data source name
		DefaultStringSize:         255,          // default size for string fields
		DisableDatetimePrecision:  true,         // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,         // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,         // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,        // auto configure based on currently MySQL version
	}), &mysqlConfig)
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	//设置与数据库建立连接的最大数目
	sqlDB.SetMaxOpenConns(config.DB.MaxOpenConns)
	//设置连接池中的最大闲置连接数
	sqlDB.SetMaxIdleConns(config.DB.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.DB.ConnMaxLifeTime)

	InitLogger.Info("connnect to mysql database successful")

}

func DisconnectDB() {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		panic(err)
	}
}

func DB(ctx *gin.Context) *gorm.DB {
	return db.WithContext(ctx)
}

func AutoMigrate() {

	//设置表默认属性
	table_options := "CHARSET=" + config.DB.CHARSET

	ctx := context.Background()
	gormLogger.Info(ctx, "migrate start...", "")
	db.Set("gorm:table_options", table_options).AutoMigrate(&model.User{})
	gormLogger.Info(ctx, "migrate end...", "")
	InitLogger.Info("migrate table successful")
}

type dbLogger struct {
	SlowThreshold time.Duration
	gorm_logger.Interface
}

// LogMode log mode
func (this *dbLogger) LogMode(level gorm_logger.LogLevel) gorm_logger.Interface {
	newlogger := *this
	return &newlogger
}

// Info print info
func (this dbLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	message := fmt.Sprintf(msg, append([]interface{}{utils.FileWithLineNum()}, data...))
	if ctx.Value("traceId") != nil {
		Logger.Info(ctx, "gorm info", message)
	} else {
		gormLogger.Info(ctx, "gorm info", message)
	}
}

// Warn print warn messages
func (this dbLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	message := fmt.Sprintf(msg, append([]interface{}{utils.FileWithLineNum()}, data...))
	if ctx.Value("traceId") != nil {
		Logger.Warn(ctx, "gorm warn", message)
	} else {
		gormLogger.Warn(ctx, "gorm warn", message)
	}

}

// Error print error messages
func (this dbLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	message := fmt.Sprintf(msg, append([]interface{}{utils.FileWithLineNum()}, data...))
	if ctx.Value("traceId") != nil {
		Logger.Error(ctx, "gorm error", message)
	} else {
		gormLogger.Error(ctx, "gorm error", message)
	}
	panic(message)
}

// Trace print sql message
func (this dbLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	sql, rows := fc()
	m := map[string]interface{}{
		"file": utils.FileWithLineNum(),
		"cost": float64(elapsed.Nanoseconds()) / 1e9,
		"sql":  sql,
		"rows": rows,
	}

	switch {
	case err != nil:
		m["err"] = err

		if ctx.Value("traceId") != nil {
			Logger.Error(ctx, "gorm", m)
		} else {
			gormLogger.Error(ctx, "gorm", m)
		}
		panic(m)
	case elapsed > this.SlowThreshold && this.SlowThreshold != 0:
		if ctx.Value("traceId") != nil {
			Logger.Warn(ctx, "gorm", m)
		} else {
			gormLogger.Warn(ctx, "gorm", m)
		}
	default:
		if ctx.Value("traceId") != nil {
			Logger.Debug(ctx, "gorm", m)
		} else {
			gormLogger.Debug(ctx, "gorm", m)
		}
	}
}
