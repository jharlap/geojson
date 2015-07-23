package main

import (
	"encoding/json"
	"fmt"

	"github.com/jharlap/geojson"
)

var jsonBlob = []byte(`[
{"type":"Point", "coordinates":[1.1,2.0]},
{"type":"LineString", "coordinates":[[1.1,2.0],[3.0,6.3]]},
{"type":"Polygon", "coordinates":[[[1.1,2.0],[3.0,6.3],[5.1,7.0],[1.1,2.0]]]},
{"type":"MultiPolygon", "coordinates":[[[[1.1,2.0],[3.0,6.3],[5.1,7.0],[1.1,2.0]]]]},

{"type":"FeatureCollection","features":[{"type":"Feature","properties":{"STATEFP":"12","COUNTYFP":"105","TRACTCE":"010300","BLKGRPCE":"2","GEOID":"121050103002","NAMELSAD":"Block Group 2","MTFCC":"G5030","FUNCSTAT":"S","ALAND":1818632,"AWATER":1112922,"INTPTLAT":"+28.0411205","INTPTLON":"-081.9336940"},"geometry":{"type":"Polygon","coordinates":[[[-81.939178,28.045386],[-81.936223,28.045389],[-81.936136,28.045118],[-81.934251,28.045141],[-81.933693,28.045299],[-81.932649,28.045823],[-81.924387,28.04591],[-81.924335,28.043554],[-81.92396,28.042763],[-81.922543,28.041462],[-81.923215,28.040831],[-81.923304,28.040356],[-81.922913,28.040008],[-81.92251,28.039602],[-81.921946,28.039514],[-81.92176,28.038963],[-81.921834,28.038254],[-81.922175,28.038138],[-81.92251,28.037527],[-81.922312,28.037087],[-81.922442,28.036592],[-81.923025,28.036608],[-81.923869,28.035805],[-81.92359,28.035382],[-81.923801,28.034595],[-81.923621,28.034083],[-81.924222,28.033494],[-81.924632,28.032108],[-81.924458,28.03131],[-81.924037,28.031007],[-81.923378,28.031821],[-81.922888,28.032227],[-81.922003,28.032243],[-81.921733,28.032427],[-81.920197,28.03241],[-81.92007,28.032094],[-81.919454,28.031655],[-81.918936,28.030834],[-81.919127,28.029614],[-81.919105,28.025963],[-81.92274,28.025932],[-81.92342,28.025704],[-81.924322,28.025723],[-81.924915,28.025929],[-81.926657,28.025951],[-81.926837,28.026111],[-81.931844,28.030708],[-81.936436,28.03488],[-81.940702,28.038785],[-81.940719,28.040547],[-81.939244,28.040561],[-81.939257,28.044087],[-81.939178,28.045386]]]}},
{"type":"Feature","properties":{"STATEFP":"12","COUNTYFP":"103","TRACTCE":"024804","BLKGRPCE":"1","GEOID":"121030248041","NAMELSAD":"Block Group 1","MTFCC":"G5030","FUNCSTAT":"S","ALAND":1105188,"AWATER":29375,"INTPTLAT":"+27.8157628","INTPTLON":"-082.7234101"},"geometry":{"type":"MultiPolygon","coordinates":[[[[-82.728555,27.816383],[-82.728484,27.820951],[-82.720426,27.820985],[-82.720427,27.813637],[-82.720392,27.808192],[-82.724059,27.808216],[-82.728597,27.808171],[-82.728555,27.816383]]]]}},
{"type":"Feature","properties":{"STATEFP":"12","COUNTYFP":"105","TRACTCE":"010800","BLKGRPCE":"1","GEOID":"121050108001","NAMELSAD":"Block Group 1","MTFCC":"G5030","FUNCSTAT":"S","ALAND":1879671,"AWATER":443468,"INTPTLAT":"+28.0365474","INTPTLON":"-081.9665586"},"geometry":{"type":"Polygon","coordinates":[[[-81.973546,28.032725],[-81.973582,28.041133],[-81.973366,28.041461],[-81.969587,28.042196],[-81.969857,28.042572],[-81.961025,28.04427],[-81.960204,28.042067],[-81.960047,28.040491],[-81.957183,28.040495],[-81.957135,28.033217],[-81.962232,28.033221],[-81.96247,28.032144],[-81.962906,28.031394],[-81.96393,28.030545],[-81.966088,28.029827],[-81.967106,28.029343],[-81.967103,28.025968],[-81.967015,28.025883],[-81.973346,28.025871],[-81.973474,28.026018],[-81.973804,28.027399],[-81.973567,28.028257],[-81.973546,28.032725]]]}}
]}
]`)

func main() {

	// unmarshal the json
	var geoObjects []geojson.Container
	err := json.Unmarshal(jsonBlob, &geoObjects)
	if err != nil {
		fmt.Println("error:", err)
	}

	// do something with the data
	for _, t := range geoObjects {
		fmt.Printf("%+v\n", t)
	}
}