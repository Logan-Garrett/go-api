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

func main() {
	router := gin.Default()
	router.GET("/dummys", getDummys)
	router.Run("localhost:9090")
}
