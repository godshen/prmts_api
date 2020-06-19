package config

/* Database config */
var Db_name = "jx_test"
var Mysql = "mysql"
var MysqlLocation = "127.0.0.1"
var MysqlPort = "3306"
var User = "root"
var Password = "root"

//var MysqlLocation = os.Getenv("DATABASE_ADDRESS")
//var MysqlPort = os.Getenv("DATABASE_PORT")

var Dbconnection = User + ":" + Password + "@tcp(" + MysqlLocation + ":" + MysqlPort + ")/" + Db_name
