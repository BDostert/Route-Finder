package main

import (
	"image/png"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//	client, e := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY")))
	//	if e != nil {
	//		log.Fatalf("fatal error: %s", e)
	//	}

	http.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
		/*
			//need to get lat, long from user either from taking location data or input address/street
			mapMaker.createReq()
			mapMaker.getLocation()
			i = mapMaker.getMap()
			png.encode(w, i)
		*/
		//		req := maps.StaticMapRequest{Center: "Ann+Arbor,MI", Size: "600x600", Zoom: 15}

		//		i, e := client.StaticMap(r.Context(), &req)
		//		if e != nil {
		//		log.Fatalf("fatal error: %s", e)
		//	}
		i := api.getMap()

		png.Encode(w, i)

		//tmpl := template.Must(template.ParseFiles("index.html"))
		//tmpl.Execute(w, nil)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
