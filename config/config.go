package config

/* Database config */
var Db_name = "jx_project"
var Mysql = "mysql"
var MysqlLocation = "117.78.10.6"
var MysqlPort = "3306"
var User = "root"
var Password = "tianHAOyuQI361"

//var MysqlLocation = os.Getenv("DATABASE_ADDRESS")
//var MysqlPort = os.Getenv("DATABASE_PORT")

var Dbconnection = User + ":" + Password + "@tcp(" + MysqlLocation + ":" + MysqlPort + ")/" + Db_name
