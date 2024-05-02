package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:      "Sydney",
		Destination: "Perth",
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	pretty.Println(route)
}
