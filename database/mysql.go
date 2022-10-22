package database

import "fmt"

type MySQLDB struct {
	Driver     string
	DataSource string
}

const (
	username = "root"
	password = "root"
	hostname = "127.0.0.1:3306"
	dbname   = "users"
)

func SetDrivers(driver string, dbName string) *MySQLDB {
	c := &MySQLDB{
		Driver:     driver,
		DataSource: DSN(dbName),
	}
	return c
}

func DSN(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, hostname, dbName)
}
