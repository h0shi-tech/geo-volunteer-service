package main

import (
 "fmt"
 "net/http"
 "html/template"
)

type GeoData struct {
 Latitude  float64 `json:"latitude"`
 Longitude float64 `json:"longitude"`
}

func MapHandler(w http.ResponseWriter, r *http.Request) {
 tmpl := template.Must(template.New("map").Parse(`<html>
<head>
    <script src="https://api.mapbox.com/mapbox-gl-js/v2.6.1/mapbox-gl.js"></script>
    <link href="https://api.mapbox.com/mapbox-gl-js/v2.6.1/mapbox-gl.css" rel="stylesheet">
</head>
<body>
    <div id="map" style="width: 100%; height: 400px;"></div>
    <script>
        mapboxgl.accessToken = 'pk.eyJ1Ijoic2NvdGhpcyIsImEiOiJjaWp1Y2ltYmUwMDBicmJrdDQ4ZDBkaGN4In0.sbihZCZJ56-fsFNKHXF8YQ'; // Replace with your Mapbox access token
        var map = new mapboxgl.Map({
            container: 'map',
            style: 'mapbox://styles/mapbox/streets-v11',
            center: [{{.Longitude}}, {{.Latitude}}],
            zoom: 12
        });

        var marker = new mapboxgl.Marker()
            .setLngLat([{{.Longitude}}, {{.Latitude}}])
            .addTo(map);
    </script>
</body>
</html>`))

 
 data := GeoData{Latitude: 40.7128, Longitude: -74.0060}

 if err := tmpl.Execute(w, data); err != nil {
  http.Error(w, err.Error(), http.StatusInternalServerError)
 }
}

func main() {
 http.HandleFunc("/map", MapHandler)
 fmt.Println("Server is on, listening on port 8080...")
 http.ListenAndServe(":8080", nil)
}
/*
Код еще допиливается, функции main1 и main2 пока неиспользуемые :)
*/