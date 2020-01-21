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
	err              error
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
	err = sqlDB.Close()
	
	if err != nil {
		panic(err)
	}
}

func CreateTableIfNotExist(table string) {
//	dropSt := `
//DROP TABLE IF EXISTS ` + table + `;
//`
//	_, err := sqlDB.Exec(dropSt)
//	common.IF(err != nil, err, "drop table " + table + " succeed.")
//
	
	createSt := `
CREATE TABLE IF NOT EXISTS ` + table +`(
	msgId          int unsigned not null primary key,
	userName       varchar(32)  not null
	
);
`
	_, err  = sqlDB.Exec(createSt)
	common.IF(err != nil, err, "create table " + table + " succeed.")
}

func InitDatabaseMySql() {
	var err error
	
	sqlDB, err = sql.Open("mysql",
		"root:password@/square?charset=utf8")
	//连接数据库，格式 用户名：密码@/数据库名？charset=编码方式
	
	if err != nil && allowSqlPanic {
		log.Println(err)
		panic("In package: dbTransaction\n" +
			"open database-MySql failed.")
	}
	
	CreateTableIfNotExist("squareMsg")
}

func TestSomeQuery(code uConfig.RequestCode, queryBody string) {

}