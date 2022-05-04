package dbDriver

import (
	"context"
	"database/sql"
	"nft_object/library/helper"
	"nft_object/library/logge"
	"os"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// MyDriver is a custom database driver, which is used for testing only.
// For simplifying the unit testing case purpose, MyDriver struct inherits the mysql driver
// gdb.DriverMysql and overwrites its functions DoQuery and DoExec.
// So if there's any sql execution, it goes through MyDriver.DoQuery/MyDriver.DoExec firstly
// and then gdb.DriverMysql.DoQuery/gdb.DriverMysql.DoExec.
// You can call it sql "HOOK" or "HiJack" as your will.
type MysqlDriver struct {
	*gdb.DriverMysql
}

var (
	// customDriverName is my driver name, which is used for registering.
	customDriverName = "MysqlDriver"
)

func init() {
	// It here registers my custom driver in package initialization function "init".
	// You can later use this type in the database configuration.
	if err := gdb.Register(customDriverName, &MysqlDriver{}); err != nil {
		panic(err)
	}
}

// New creates and returns a database object for mysql.
// It implements the interface of gdb.Driver for extra database driver installation.
func (d *MysqlDriver) New(core *gdb.Core, node *gdb.ConfigNode) (gdb.DB, error) {
	return &MysqlDriver{
		&gdb.DriverMysql{
			Core: core,
		},
	}, nil
}

// DoQuery commits the sql string and its arguments to underlying driver
// through given link object and returns the execution result.
func (d *MysqlDriver) DoQuery(ctx context.Context, link gdb.Link, sql string, args ...interface{}) (rows *sql.Rows, err error) {

	tsMilli := gtime.TimestampMilli()
	rows, err = d.DriverMysql.DoQuery(ctx, link, sql, args...)
	//  查询sql日志   上线应该关闭
	recordLog(ctx, sql, float64(gtime.TimestampMilli()-tsMilli), err, "doQuery", args)
	return
}

// DoExec commits the query string and its arguments to underlying driver
// through given link object and returns the execution result.
func (d *MysqlDriver) DoExec(ctx context.Context, link gdb.Link, sql string, args ...interface{}) (result sql.Result, err error) {

	tsMilli := gtime.TimestampMilli()
	result, err = d.DriverMysql.DoExec(ctx, link, sql, args...)
	recordLog(ctx, sql, float64(gtime.TimestampMilli()-tsMilli), err, "doExec", args)
	return

}

// 记录sql 日志
func recordLog(ctx context.Context, sql string, cost float64, err error, topic string, args ...interface{}) {
	if err != nil {
		// sql 出错
		topic = "error"
	}
	if cost > 500 {
		// sql 运行时长太长
		topic = "warn"
	}
	logge.WriteSql("sql_"+topic, topic, map[string]interface{}{
		"sql":        gdb.FormatSqlWithArgs(sql, args),
		"cost":       cost,
		"time":       gtime.Now(),
		"error":      err,
		"pid":        os.Getpid(),
		"login_name": helper.GetRtx(ctx),
	})
}
