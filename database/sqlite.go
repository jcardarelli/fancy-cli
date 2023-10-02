package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

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
func InitSqlDatabase() (*sql.DB, error) {
	var err error
	var cp *sql.DB

	// Setup sql file to initialize the restaurants table
	dbInitSqlFile := filepath.Join("sql", "create-table.sql")
	sqlFileContents, err := os.ReadFile(dbInitSqlFile)
	if err != nil {
		log.Fatalln("unable to find setup sql file to initialize database tables", err)
	}

	// Read the contents of the sql file to a string
	sql := string(sqlFileContents)

	// Validate connection to sqlite database
	if _, err := cp.Exec(sql); err != nil {
		log.Fatalln("could not exec database query", err)
	}

	return db, nil
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
func GetRestaurant(restaurantName string) types.Restaurant {
	var id int
	var name string
	var address string
	var stars int

	sqlStatement := `SELECT * FROM restaurants WHERE name = $1;`

	row := db.QueryRow(sqlStatement, restaurantName)
	switch err := row.Scan(&id, &name, &address, &stars); err {
	case sql.ErrNoRows:
		log.Fatalln("failed to get matching row", err)
	case nil:
	default:
		log.Fatalln("failed in default case", err)
	}

	return types.Restaurant{Id: id, Name: name, Address: address, Stars: stars}
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
	t.AppendHeader(table.Row{"#", "Name", "Address", "Stars"})
	t.SetStyle(table.StyleColoredBlueWhiteOnBlack)

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
