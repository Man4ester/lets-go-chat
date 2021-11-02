package main

import (
	"github.com/labstack/echo/v4"
	handlers "lets-go-chat/internal/handlers"
	spec "lets-go-chat/pkg/openapi3"
)

func main(){
	e := echo.New()
	spec.RegisterHandlers(e, handlers.Server{})
	e.Logger.Fatal(e.Start(":8080"))


}