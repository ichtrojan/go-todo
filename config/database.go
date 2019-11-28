package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Database() *sql.DB {
	database, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/?parseTime=true")

	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("Database Connection Successful")
	}

	_, err = database.Exec(`CREATE DATABASE gotodo`)

	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec(`USE gotodo`)

	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec(`
		CREATE TABLE todos (
		    id INT AUTO_INCREMENT,
		    item TEXT NOT NULL,
		    completed BOOLEAN DEFAULT FALSE,
		    PRIMARY KEY (id)
		);
	`)

	if err != nil {
		fmt.Println(err)
	}

	return database
}
