package dbTransaction

import (
	"database/sql"
	"log"
	
	"AppServer/src/common"
	uConfig "AppServer/src/config"
	_ "github.com/go-sql-driver/mysql"
)

var (
	sqlDB            *sql.DB
	allowSqlPanic    = true
//	sqlMutex         sync.Mutex
)

func TurnSqlPanic() {
	allowSqlPanic = !allowSqlPanic
	common.IF(allowSqlPanic, "panic of sql operation was allowed",
		"panic of sql operation was forbidden")
}

func DisconnectMySql() {
	if sqlDB == nil {
		panic("sql index is referring to nil!")
	}
	err := sqlDB.Close()
	
	if err != nil {
		panic(err)
	}
}

func CreateTable() {}

func InitDatabaseMySql() {
	var err error
	
	sqlDB, err = sql.Open("mysql",
		"root:password@/server?charset=utf8")
	//连接数据库，格式 用户名：密码@/数据库名？charset=编码方式
	
	if err != nil && allowSqlPanic {
		log.Println(err)
		panic("In package: dbTransaction\n" +
			"open database-MySql failed.")
	}
	
}

func TestSomeQuery(code uConfig.RequestCode) {

}