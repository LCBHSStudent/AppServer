package dbTransaction

import (
	"database/sql"
	"log"
	
	"AppServer/src/common"
	uConfig "AppServer/src/config"
	_ "github.com/go-sql-driver/mysql"
)

//alter table tableName add columnName varchar(30) 增加列
//EXEC  sp_rename   'tableName.column1' , 'column2'
//      (把表名为tableName的column1列名修改为column2)
//alter table tableName drop column columnName     删除列

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

func CreateTableIfNotExist(table string, tableFormat string) {
//	dropSt := `
//DROP TABLE IF EXISTS ` + table + `;
//`
//	_, err := sqlDB.Exec(dropSt)
//	common.IF(err != nil, err, "drop table " + table + " succeed.")
//
	
	createSt := `
CREATE TABLE IF NOT EXISTS ` + table + tableFormat
	_, err  = sqlDB.Exec(createSt)
	common.IF(err != nil, err, "create table " + table + " succeed.")
}

func checkTableExist(table string) bool {
	query := "SELECT nsp FROM " + table
	_, err = sqlDB.Query(query)
	if err != nil && err.Error() == "Error 1054: Unknown column 'nsp' in 'field list'" {
		return true
	} else {
		return false
	}
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
	
	if !checkTableExist("user") {
		CreateTableIfNotExist("user",
			`
(
	user_id               varchar(32) not null primary key,
	user_name             varchar(32) not null,
	password              varchar(32) not null,
	oauth_name            varchar(32),
	oauth_access_token    varchar(255),
	oauth_expires         int,
	location              varchar(128),
	host                  varchar(32),
	api_key               varchar(32),
	api_secret            varchar(255),
	last_login_date       date
);
`)
	}
	
	if !checkTableExist("squareMsg") {
		CreateTableIfNotExist("squareMsg",
			`
(
	msgId          int unsigned not null primary key,
	userName       varchar(32)  not null
);
`)
	}
}

func TestSomeQuery(code uConfig.RequestCode, queryBody string) {

}