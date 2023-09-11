package main

import "fmt"

type restaurant struct {
	name           string
	location       string
	michelin_stars int
}

func newRestaurant(name string, location string, michelin_stars int) *restaurant {
	r := restaurant{name: name, location: location, michelin_stars: michelin_stars}
	return &r
}

func main() {
	// fmt.Println(restaurant{name: "The French Laundry", location: "Yountville, CA", michelin_stars: 3})
	restaurant_name := "The French Laundry"
	location := "Yountville, CA"
	michelin_stars := 3
	fmt.Println("restaurant_name", restaurant_name, "location", location, "michelin_stars", michelin_stars)
	// newRestaurant()
}
