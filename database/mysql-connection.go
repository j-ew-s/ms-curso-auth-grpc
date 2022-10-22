package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLCommand struct {
	SqlConn *MySQLDB
}

func (sqlCommand SQLCommand) Ping() error {

	db, err := gorm.Open(mysql.Open(sqlCommand.SqlConn.DataSource), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println("FAIL TO DEFINE DB!")
		return err
	}

	err = sqlDb.Ping()
	if err != nil {
		fmt.Println("FAIL TO CONNECT TO SQL SERVER!")
		return err
	}

	fmt.Println("Connected!")
	return nil
}

func (sqlCommand SQLCommand) OpenConnection() (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(sqlCommand.SqlConn.DataSource), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil, err
	}

	return db, nil

}
