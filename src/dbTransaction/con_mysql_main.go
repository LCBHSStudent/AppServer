package dbTransaction

import (
	"database/sql"
	"log"
	
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
	log.Println("Package 'dbTransaction'")
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

func InitDatabaseMySql() {
	var err error
	
	sqlDB, err = sql.Open("mysql",
		"username:psw@/dbName?charset=utf8")
	//连接数据库，格式 用户名：密码@/数据库名？charset=编码方式
	
	if err != nil && allowSqlPanic {
		panic("In package: dbTransaction\n" +
			"open database-MySql failed.")
	}
	
}

func DoQuery(code uConfig.RequestCode) {

}