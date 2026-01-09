package services

import (
	"encoding/json"
	"fmt"


	// "gestAdressNetwork/internal/model"
	"io"
	"net/http"
)

type AdressesService struct {
}

// func NewBookService() *AdressesService {
// 	return &AdressesService{}
// }

// func GetXYAdress() (*model.XY, error) {

// 	var xy *model.XY

// 	response, err := http.Get("http://api-adresse.data.gouv.fr/search/?q=Paris&limit=1")
// 	if err != nil {
// 		fmt.Printf("The http request failed")
// 	} else {
// 		data, _ := io.ReadAll(response.Body)

// 		fmt.Println(string(data))
// 	}

	

// 	return xy, nil

// }


func NewBookService() *AdressesService {
	return &AdressesService{}
}


type Response struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`

}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}


func GetXYAdress()  (float64, float64 , error){
   
		response, err := http.Get("http://api-adresse.data.gouv.fr/search/?q=Paris&limit=1")
	

		data, _ := io.ReadAll(response.Body)

		var result Response
		err = json.Unmarshal(data, &result)

		fmt.Println(" je suis laaa ", result.Features[0].Type)

		var x float64   
		var y float64

		if len(result.Features) > 0 {
        x = result.Features[0].Geometry.Coordinates[0]
        y = result.Features[0].Geometry.Coordinates[1]

        fmt.Println("Latitude :", x)
        fmt.Println("Longitude:", y)
    }



		return  x, y ,  err

}