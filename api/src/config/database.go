package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() (*sql.DB, error) {
	datasourceNameString := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBName,
	)

	db, err := sql.Open("mysql", datasourceNameString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
