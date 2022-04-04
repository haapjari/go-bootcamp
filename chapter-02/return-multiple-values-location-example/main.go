package main

import "fmt"

func location(name string, city string) (region string, continent string) {
	switch city {
	case "New York", "LA", "Chicago":
		continent = "North America"
	default:
		continent = "Unknown"
	}
	return
}

func main() {
	region, continent := location("Matt", "LA")
	fmt.Printf("%s lives in %s\n", region, continent)
}