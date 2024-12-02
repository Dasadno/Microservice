package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Text string `json:text`
	ID   int    `json:"id`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var messages []Message
var nextID = 1

func GetHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, &messages)
}
func PostHandler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}

	message.ID = nextID
	nextID++

	messages = append(messages, message)
	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message was successfully added",
	})
}

func PatchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Bad ID",
		})
	}

	var updatedMessage Message
	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not update the message",
		})
	}

	updated := false

	for i, message := range messages {
		if message.ID == id {
			updatedMessage.ID = id
			messages[i] = updatedMessage
			updated = true
			break
		}
	}

	if !updated {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not find the message",
		})
	}

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message was updated",
	})
}

func main() {
	e := echo.New()

	e.GET("/messages", GetHandler)
	e.POST("/messages", PostHandler)
	e.PATCH("/messages/:id", PatchHandler)
	e.Start(":8080")
}
