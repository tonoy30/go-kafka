package main

import (
	"encoding/json"
	"go-kafka/consumer"
	"go-kafka/domain"
	"go-kafka/model"
	"go-kafka/producer"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func createComment(e echo.Context) error {
	comment := new(domain.Comment)
	if err := e.Bind(comment); err != nil {
		return err
	}
	bytes, _ := json.Marshal(comment)
	msg, err := producer.PushCommitToTopic("comments", bytes)
	if err != nil {
		errResponse := model.NewResponse(http.StatusInternalServerError, "internal server error", comment)
		return e.JSON(http.StatusInternalServerError, errResponse)
	}
	response := model.NewResponse(http.StatusOK, msg, comment)
	return e.JSON(http.StatusOK, response)
}
func main() {
	e := echo.New()
	e.GET("/", helloWorld)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	kafka := e.Group("/kafka")
	kafka.POST("/comments", createComment)
	go consumer.ConsumePartition([]string{"localhost:9092"}, "comments")
	e.Logger.Fatal(e.Start(":1323"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
