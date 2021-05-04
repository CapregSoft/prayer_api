package prayerapi

import (
	"net/http"

	"github.com/CapregSoft/prayer-api/controller"
	"github.com/labstack/echo/v4"
)

type SalahApi struct {
	e *echo.Echo
}

func NewSalahApi() *SalahApi {

	return &SalahApi{
		e: echo.New(),
	}

}

func (s *SalahApi) Start(port string) {

	if port == "" {
		port = ":8080"
	} else if port[0:1] != ":" {
		port = ":" + port

	}

	s.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Prayer Server running")
	})
	//s.e.GET("/all", controller.GetPrayerTime)
	s.e.POST("/prayertime", controller.GetPrayerTimePOST)

	s.e.Logger.Fatal(s.e.Start(port))

}
