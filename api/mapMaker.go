package api

import (
	"context"
	"image"
	"log"
	"os"

	"googlemaps.github.io/maps"
)

var client, e = maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY")))

const length = "600"

// gets users start lat lng from geolocate
func getStartLocation() *maps.LatLng {
	locationRes, e := client.Geolocate(context.TODO(), &maps.GeolocationRequest{})
	if e != nil {
		log.Fatalf("fatal error: %s", e)
	}
	return &locationRes.Location
}

// creates static req from
func createStaticReq() *maps.StaticMapRequest {
	return &maps.StaticMapRequest{
		Center: getStartLocation().String(),
		Size:   length + "x" + length,
		Zoom:   15,
	}

}

// returns png of map to be rendered
func getMap() *image.Image {
	req := createStaticReq()
	i, e := client.StaticMap(context.Background(), req)
	if e != nil {
		log.Fatalf("fatal error: %s", e)
	}
	return &i
}

func getRadius() {

}
