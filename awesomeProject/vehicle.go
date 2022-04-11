package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type vehicle struct {
	ID     string `json:"id"`
	Colour string `json:"colour"`
	Make   string `json:"make"`
	Model  string `json:"model"`
}

var vehicles = []vehicle{
	{ID: "1", Colour: "Red", Make: "BMW", Model: "Series 3"},
	{ID: "2", Colour: "White", Make: "Audi", Model: "A3"},
	{ID: "3", Colour: "Red", Make: "Mercedes", Model: "E Class"},
	{ID: "4", Colour: "Black", Make: "BMW", Model: "Series 3"},
	{ID: "5", Colour: "Red", Make: "BMW", Model: "Series 5"},
	{ID: "6", Colour: "Orange", Make: "Mercedes", Model: "C Class"},
}

func vehicleById(c *gin.Context) {
	id := c.Param("id")
	vehicle, err := getVehicleById(id)

	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, vehicle)
}

func getVehicleById(id string) (*vehicle, error) {
	for i, b := range vehicles {
		if b.ID == id {
			return &vehicles[i], nil
		}
	}
	return nil, errors.New("can't find vehicle")
}

func getVehicles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, vehicles)
}

func main() {
	router := gin.Default()
	router.GET("/vehicle", getVehicles)
	router.GET("/vehicle/:id", vehicleById)

	err := router.Run("localhost:8000")
	if err != nil {
		fmt.Println(err)
	}
}
