package mysql

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql" //import mysql impolementation for sql
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

var s IDB

func InitSQL(config Config) {
	ss, er := New(config)
	if er != nil {
		panic(er)
	}
	s = ss
}

//Config Mysql的配置
type Config struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Db           string `json:"db"`
	Dbprefix     string `json:"dbprefix"`
	ConnLifeTime int    `json:"connlifetime"` //以秒为单位
	MaxIdleConn  int    `json:"maxidleconn"`
	MaxOpenConn  int    `json:"maxopenconn"`
}

//IDB 数据库接口
type IDB interface {
	GetDB() *sqlx.DB
	Prefix(str string) string
	UnPrefix(str string) string
	GetPrefix() string
}

type sqlServer struct {
	Config Config
	DB     *sqlx.DB
}

//init 初始化sql
func (sql *sqlServer) init(config Config) error {
	var dbonce sync.Once
	var db *sqlx.DB
	var err error
	dbonce.Do(func() {
		db, err = sqlx.Open(
			"mysql",
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s",
				config.User,
				config.Password,
				config.Host,
				config.Port,
				config.Db,
			),
		)
		if err != nil {
			log.Printf("get mysql database error: %s", err)
		} else {
			db.SetConnMaxLifetime(time.Duration(config.ConnLifeTime) * time.Second)
			db.SetMaxIdleConns(config.MaxIdleConn)
			db.SetMaxOpenConns(config.MaxOpenConn)
			db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
		}
	})

	sql.DB = db
	return err
}

//GetDB GetDB
func (sql *sqlServer) GetDB() *sqlx.DB {
	return sql.DB
}

//New 创建一个新的数据库
func New(config Config) (sql IDB, err error) {
	s := &sqlServer{
		Config: config,
	}
	err = s.init(config)
	sql = s
	return
}

//Prefix change the relative sql to real sql with prefix
func (sql sqlServer) Prefix(str string) string {
	return strings.Replace(str, "#__", sql.Config.Dbprefix, -1)
}

//UnPrefix change the real sql with prefix to relative one
func (sql sqlServer) UnPrefix(str string) string {
	return strings.Replace(str, sql.Config.Dbprefix, "#__", 1)
}

//GetPrefix get sql prefix
func (sql sqlServer) GetPrefix() string {
	return sql.Config.Dbprefix
}

func GetDB() *sqlx.DB {
	return s.GetDB()
}

//Prefix change the relative sql to real sql with prefix
func Prefix(sql string) string {
	return s.Prefix(sql)
}

//UnPrefix change the real sql with prefix to relative one
func UnPrefix(sql string) string {
	return s.UnPrefix(sql)
}

//GetMysql GetMysql
func GetMysql() IDB {
	return s
}
