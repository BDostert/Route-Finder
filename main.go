package main

import (
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c, e := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY")))
	if e != nil {
		log.Fatalf("fatal error: %s", e)
	}

	http.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		//need to get lat, long from user
		req := maps.StaticMapRequest{Center: "Ann+Arbor,MI", Size: "600x600", Zoom: 15}

		i, e := c.StaticMap(r.Context(), &req)
		if e != nil {
			log.Fatalf("fatal error: %s", e)
		}
		png.Encode(w, i)

		//tmpl := template.Must(template.ParseFiles("index.html"))
		//tmpl.Execute(w, nil)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
