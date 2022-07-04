package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/album", getAllAlbums)

	router.GET("/data", getSomeData)

	router.POST("/input", postAlbums)

	router.Run("localhost:8080")
	//fmt.Printf(albums[0].ID)
}

func getAllAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getSomeData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Salam , Mn Nima Hastam , Be Go Khosh umadi  ;-)")
}

func postAlbums(c *gin.Context) {
	var id string

	if err := c.BindJSON(&id); err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, id)

	//for i := 0; i < len(albums); i++ {
	//	if albums[i].ID == id {
	//		c.IndentedJSON(http.StatusOK, albums[i])
	//	}
	//}

}
