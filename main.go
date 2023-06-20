package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type dummy struct {
	ID         string `json:"id"`
	Restaurant string `json:"restaurant"`
	Good       bool   `json:"good"`
}

var dummys = []dummy{
	{ID: "1", Restaurant: "Taco Bell", Good: false},
	{ID: "2", Restaurant: "Wendy's", Good: false},
	{ID: "3", Restaurant: "Mcdonald's", Good: false},
	{ID: "4", Restaurant: "Arby's", Good: false},
}

func getDummys(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, dummys)
}

func addDummy(context *gin.Context) {
	var newDummy dummy

	if err := context.BindJSON(&newDummy); err != nil {
		return
	}

	dummys = append(dummys, newDummy)
	context.IndentedJSON(http.StatusCreated, newDummy)
}

func main() {
	router := gin.Default()
	router.GET("/dummys", getDummys)
	router.POST("/dummys", addDummy)
	router.Run("localhost:9090")
}
