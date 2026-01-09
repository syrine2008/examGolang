package main

import (
	// "gestAdressNetwork/internal/services"

	"gestAdressNetwork/internal/services"

	"github.com/gin-gonic/gin"
	// "log"
	// "net/http"
)

// func main() {

// 	//  mux := http.NewServeMux()
// 	//  mux.HandleFunc("/getXy", func(w http.ResponseWriter, r *http.Request) {
// 	// 	services.GetXYAdress(w,r)
// 	//  })

// 	// log.Println("HTTP server started on :8080")

// 	r := gin.Default()

// 	// r.GET("/getXy", func(c *gin.Context) {
// 	// 	c.JSON(http.StatusOK,services.GetXYAdress)
// 	// })

// 		r.GET("/getXy", func(c *gin.Context){

// 		response, err := http.Get("http://api-adresse.data.gouv.fr/search/?q=Paris&limit=1")

// 	if err != nil {
// 		fmt.Printf("The http request failed")
// 	} else {
// 		data, _ := io.ReadAll(response.Body)
// 		fmt.Println((data))
// 		// fmt.Println("rugzgurguzrugzrguzrgurz")
//        	// fmt.Printf("Type: %T", data)
// 		// fmt.Println("rugzgurguzrugzrguzrgurz")
// 		// fmt.Printf("Type: %T", response.Body)
// 		// fmt.Println(string(data))

// 		}

// 	 r.Run()

// })
// }

func main() {

	router := gin.Default()
	router.GET("/getXy", func(c *gin.Context) {
		services.GetXYAdress()
	})

	//Parse
	router.GET("/parse", func(c *gin.Context) {

		services.Parse("gestAdressNetwork/2018_01_Sites_mobiles_2G_3G_4G_France_metropolitaine_L93_ver2.csv")

	})
	router.Run(":8080")
}
