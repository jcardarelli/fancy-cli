package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/mattn/go-sqlite3"
)

const (
	file   string = "goproj.db"
	create string = `
create table if not exists restaurants (
	id INTEGER NOT NULL PRIMARY KEY,
	name TEXT,
	address TEXT,
	stars INTEGER
);`
)

type Restaurant struct {
	name    string
	address string
	stars   int
	db      *sql.DB
}

// initSqlDatabase: Setup database connection
func initSqlDatabase() (*Restaurant, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatalln("could not establish database connection", err)
	}
	if _, err := db.Exec(create); err != nil {
		log.Fatalln("could not exec database query", err)
	}
	return &Restaurant{
		db: db,
	}, nil
}

func (c *Restaurant) insertRestaurant(name string, address string, stars int) (int, error) {
	res, err := c.db.Exec("insert into restaurants(name, address, stars) values(?, ?, ?);", name, address, stars)
	if err != nil {
		return 0, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}
	return int(id), nil
}

func main() {
	restaurant_name := "The French Laundry"
	address := "Yountville, CA"
	michelin_stars := 3
	// restaurant_name := "Birdsong"
	// address := "San Francisco, CA"
	// michelin_stars := 2
	// fmt.Println(newRestaurant(restaurant_name, address, michelin_stars))

	dbConnection, err := initSqlDatabase()
	if err != nil {
		log.Fatalln("could create new Restaurant", err)
	}
	dbConnection.insertRestaurant(restaurant_name, address, michelin_stars)
	fmt.Println(sqlite3.Version())
}
