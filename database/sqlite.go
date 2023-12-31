package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/jcardarelli/fancy-cli/types"
	"github.com/jedib0t/go-pretty/table"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Open connection to the sqlite db file
func OpenConnectionPool(sqliteDbFile string) error {
	var err error

	db, err = sql.Open("sqlite3", sqliteDbFile)
	if err != nil {
		log.Fatalln("could not establish database connection", err)
	}

	return db.Ping()
}

// Setup database connection
func InitSqlDatabase(dbFile string, sqlInitStatements string) error {
	// Attempt to open a connection to the sqlite file
	connErr := OpenConnectionPool(dbFile)
	if connErr != nil {
		log.Fatalln("error connecting to sqlite", connErr)
	}
	log.Println("setup connection to sqlite")

	// Setup sql file to initialize the restaurants table
	sqlFileContents, sqlFileErr := os.ReadFile(sqlInitStatements)
	if sqlFileErr != nil {
		log.Fatalln("unable to find setup sql file to initialize restaurants table", sqlFileErr)
	} else {
		log.Println("found sql file at:", sqlInitStatements)
	}

	// Validate connection to sqlite database
	statement, connErr := db.Prepare(string(sqlFileContents))
	if connErr != nil {
		log.Fatalln("could not prepare CREATE TABLE query", connErr)
	}
	_, stErr := statement.Exec()
	if stErr != nil {
		log.Fatalln("unable to execute prepared sql statement", stErr)
	}

	return nil
}

// Insert a new restaurant into the restaurants table and return the id
func InsertRestaurant(name string, address string, stars int) (uint, error) {
	var id int64

	response, err := db.Exec(
		"INSERT INTO restaurants(name, address, stars) VALUES(?, ?, ?);",
		name, address, stars,
	)
	if err != nil {
		return 0, err
	}

	if id, err = response.LastInsertId(); err != nil {
		return 0, err
	}
	return uint(id), nil
}

// Get a restaurant from the restaurants table
func GetRestaurant(restaurantName string) {
	var id int
	var name string
	var address string
	var stars int

	sqlStatement := `SELECT * FROM restaurants WHERE name = $1;`

	// setup table
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredBlueWhiteOnBlack)
	t.AppendHeader(table.Row{"#", "Name", "Address", "Stars"})

	row := db.QueryRow(sqlStatement, restaurantName)
	switch err := row.Scan(&id, &name, &address, &stars); err {
	case sql.ErrNoRows:
		log.Fatalln("failed to get matching row", err)
	case nil:
	default:
		log.Fatalln("failed in default case", err)
	}

	t.AppendRow([]interface{}{id, name, address, stars})

	t.Render()
}

// Get all restaurants from the restaurants table
func GetAllRestaurants() {
	sqlStatement := `SELECT * FROM restaurants;`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalln("failed to query all rows in restaurants table", err)
	}
	defer rows.Close()

	// setup table
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredBlueWhiteOnBlack)
	t.AppendHeader(table.Row{"#", "Name", "Address", "Stars"})

	for rows.Next() {
		restaurant := types.Restaurant{}
		switch scanErr := rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Address, &restaurant.Stars); scanErr {
		case sql.ErrNoRows:
			log.Println("rows:", rows)
			log.Fatalln("failed to get any matching rows", scanErr)
		case nil:
		default:
			log.Fatalln("failed in default case", scanErr)
		}

		t.AppendRows([]table.Row{
			{restaurant.Id, restaurant.Name, restaurant.Address, restaurant.Stars},
		})
	}

	if err != nil {
		log.Fatalln(err)
	}

	t.Render()
}
