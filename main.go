package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jcardarelli/fancy-cli/cmd"
	"github.com/mattn/go-sqlite3"
)

const (
	sqliteDbFile string = "fancy-cli.db"
)

type Restaurant struct {
	name    string
	address string
	stars   int
	db      *sql.DB
}

// initSqlDatabase: Setup database connection
func initSqlDatabase() (*Restaurant, error) {
	// Open connection to the sqlite db file
	db, err := sql.Open("sqlite3", sqliteDbFile)
	if err != nil {
		log.Fatalln("could not establish database connection", err)
	}

	// Setup sql file to initialize the restaurants table
	dbInitSqlFile := filepath.Join("sql", "create-table.sql")
	c, err := os.ReadFile(dbInitSqlFile)
	if err != nil {
		log.Fatalln("unable to find setup sql file to initialize database tables", err)
	}
	// read the contents of the sql file to a string
	sql := string(c)

	// Validate connection to sqlite database
	if _, err := db.Exec(sql); err != nil {
		log.Fatalln("could not exec database query", err)
	}
	return &Restaurant{
		db: db,
	}, nil
}

func (c *Restaurant) insertRestaurant(
	name string,
	address string,
	stars int,
) (int, error) {
	res, err := c.db.Exec(
		"insert into restaurants(name, address, stars) values(?, ?, ?);",
		name, address, stars,
	)
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

	// establish db connection
	db, err := initSqlDatabase()
	if err != nil {
		log.Fatalln("could create new Restaurant", err)
	}

	cmd.Execute()

	// insert new restaurant
	db.insertRestaurant(restaurant_name, address, michelin_stars)
	fmt.Println(sqlite3.Version())
}
